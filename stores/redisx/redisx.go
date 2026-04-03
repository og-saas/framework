package redisx

import (
	"context"
	"sync"
	"time"

	"github.com/og-saas/framework/utils/tenant"

	"github.com/redis/go-redis/v9"
)

var (
	Engine *RDBEngine
	once   sync.Once
)

type RDBEngine struct {
	pool sync.Map
}

func Must(c Config) {
	must(tenant.Default, c)
}

func MustTenant(providers ...TenantConfigProvider) {
	for _, p := range providers {
		configMap, err := p.Load()
		if err != nil {
			panic(err)
		}

		for key, val := range configMap {
			must(key, val)
		}
	}
}

func must(tenant int64, cfg Config) {
	once.Do(func() {
		Engine = &RDBEngine{}
	})
	rdb := cfg.NewRdb()
	if rdb == nil {
		panic("rdb init failed")
	}
	Engine.pool.Store(tenant, rdb)
}

func (c Config) NewRdb() (rdb redis.UniversalClient) {
	opt := &redis.UniversalOptions{
		Addrs:      c.Addrs,
		Username:   c.Username,
		Password:   c.Password,
		MasterName: c.MasterName,
		DB:         c.DB,
	}

	if c.PoolSize > 0 {
		opt.PoolSize = c.PoolSize
	} else {
		opt.PoolSize = 100 // default pool size
	}

	if c.MinIdleConns > 0 {
		opt.MinIdleConns = c.MinIdleConns
	} else {
		opt.MinIdleConns = 20 // default min idle conns
	}

	// Default to reasonable timeouts if not configured
	if c.DialTimeout > 0 {
		opt.DialTimeout = time.Duration(c.DialTimeout) * time.Millisecond
	} else {
		opt.DialTimeout = 5 * time.Second
	}

	if c.ReadTimeout > 0 {
		opt.ReadTimeout = time.Duration(c.ReadTimeout) * time.Millisecond
	} else {
		opt.ReadTimeout = 3 * time.Second
	}

	if c.WriteTimeout > 0 {
		opt.WriteTimeout = time.Duration(c.WriteTimeout) * time.Millisecond
	} else {
		opt.WriteTimeout = 3 * time.Second
	}

	rdb = redis.NewUniversalClient(opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	if c.Debug {
		rdb.AddHook(DebugHook{})
	}

	return rdb
}

func (e *RDBEngine) RDB(ctx context.Context) redis.UniversalClient {

	if rdb, ok := e.getClientForTenant(tenant.GetTenantId(ctx)); ok {
		return rdb
	}

	if rdb, ok := e.getClientForTenant(tenant.Default); ok {
		return rdb
	}

	return nil

}

// getClientForTenant 从 pool 安全获取 redis client
func (e *RDBEngine) getClientForTenant(tenantId int64) (redis.UniversalClient, bool) {
	v, ok := e.pool.Load(tenantId)
	if !ok || v == nil {
		return nil, false
	}
	rdb, ok := v.(redis.UniversalClient)
	return rdb, ok
}

func (e *RDBEngine) Map() map[int64]redis.UniversalClient {
	m := make(map[int64]redis.UniversalClient)
	e.pool.Range(func(key, value interface{}) bool {
		m[key.(int64)] = value.(redis.UniversalClient)
		return true
	})
	return m
}
