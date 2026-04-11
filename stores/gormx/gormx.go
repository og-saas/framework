package gormx

import (
	"context"
	"fmt"
	"sync"

	"github.com/og-saas/framework/utils/tenant"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	Engine *DBEngine
	once   sync.Once
)

type DBEngine struct {
	Mysql      *DBManager
	Postgres   *DBManager
	Clickhouse *DBManager
}

type DBManager struct {
	pool sync.Map
}

func (e *DBEngine) DB(drivers ...string) *DBManager {
	var driver string
	if len(drivers) == 0 || drivers[0] == "" {
		driver = DriverMysql // 默认 MySQL
	} else {
		driver = drivers[0]
	}
	switch driver {
	case DriverMysql:
		return e.Mysql
	case DriverPostgres:
		return e.Postgres
	case DriverClickHouse:
		return e.Clickhouse
	}
	panic("gorm: unknown driver")
}

func (s *DBManager) WithContext(ctx context.Context) *gorm.DB {
	if db, ok := s.getClientForTenant(tenant.GetTenantId(ctx)); ok {
		return db.WithContext(ctx)
	}

	if db, ok := s.getClientForTenant(tenant.Default); ok {
		return db.WithContext(ctx)
	}
	panic("gorm: database not initialized")
}

// OriginalDb 返回原始的 gorm.DB 对象
func (s *DBManager) OriginalDb(ctx context.Context) *gorm.DB {
	if db, ok := s.getClientForTenant(tenant.GetTenantId(ctx)); ok {
		return db
	}

	if db, ok := s.getClientForTenant(tenant.Default); ok {
		return db
	}
	panic("gorm: database not initialized")
}

// WithWriteContext 创建一个只写的 gorm.DB 对象
func (s *DBManager) WithWriteContext(ctx context.Context) *gorm.DB {
	if db, ok := s.getClientForTenant(tenant.GetTenantId(ctx)); ok {
		return db.Clauses(dbresolver.Write).WithContext(ctx)
	}

	if db, ok := s.getClientForTenant(tenant.Default); ok {
		return db.Clauses(dbresolver.Write).WithContext(ctx)
	}
	panic("gorm: database not initialized")
}

func (s *DBManager) getClientForTenant(tenantId int64) (*gorm.DB, bool) {
	v, ok := s.pool.Load(tenantId)
	if !ok || v == nil {
		return nil, false
	}
	db, ok := v.(*gorm.DB)
	return db, ok
}

func Must(configs ...Config) {
	must(tenant.Default, configs...)
}

// UpdateTenant 更新租户配置，configMap 中的全部更新，pool 中存在但 configMap 中没有的删除
func UpdateTenant(providers ...TenantConfigProvider) {
	for _, p := range providers {
		configMap, err := p.Load()
		if err != nil {
			logx.Errorf("gorm: update tenant load config error: %v", err)
			return
		}

		// 按 driver 分组记录 configMap 中的 tenantId
		driverTenants := make(map[string]map[int64]struct{})
		for tenantId, cfg := range configMap {
			if driverTenants[cfg.Driver] == nil {
				driverTenants[cfg.Driver] = make(map[int64]struct{})
			}
			driverTenants[cfg.Driver][tenantId] = struct{}{}
			must(tenantId, cfg)
		}

		// 删除 pool 中存在但 configMap 中没有的租户
		for driver, tenantIds := range driverTenants {
			mgr := getManager(driver)
			if mgr == nil {
				continue
			}
			mgr.pool.Range(func(key, _ interface{}) bool {
				tid, ok := key.(int64)
				if !ok {
					return true
				}
				if tid == tenant.Default {
					return true
				}
				if _, exists := tenantIds[tid]; !exists {
					mgr.pool.Delete(tid)
				}
				return true
			})
		}
	}
}

// AppendTenant 追加租户配置，不存在则添加
func AppendTenant(providers ...TenantConfigProvider) {
	for _, p := range providers {
		configMap, err := p.Load()
		if err != nil {
			logx.Errorf("gorm: append tenant load config error: %v", err)
			return
		}

		for tenantId, cfg := range configMap {
			mgr := getManager(cfg.Driver)
			if mgr == nil {
				logx.Errorf("gorm: append tenant manager not initialized, driver: %s", cfg.Driver)
				return
			}
			if _, ok := mgr.getClientForTenant(tenantId); !ok {
				must(tenantId, cfg)
			}
		}
	}
}

func getManager(driver string) *DBManager {
	if Engine == nil {
		return nil
	}
	switch driver {
	case DriverMysql:
		return Engine.Mysql
	case DriverPostgres:
		return Engine.Postgres
	case DriverClickHouse:
		return Engine.Clickhouse
	default:
		logx.Errorf("gorm: unknown driver: %s", driver)
		return nil
	}
}

func MustTenant(providers ...TenantConfigProvider) {
	for _, p := range providers {
		configMap, err := p.Load()
		if err != nil {
			panic(err)
		}

		for key, val := range configMap {
			tenantId := cast.ToInt64(key)
			must(tenantId, val)
		}
	}
}

func must(tenantId int64, configs ...Config) {
	once.Do(func() {
		Engine = &DBEngine{}
	})
	if len(configs) == 0 {
		panic("gorm: empty config")
	}

	for _, cfg := range configs {
		db := cfg.NewDB()
		if db == nil {
			panic("gorm: db init failed")
		}

		var mgr *DBManager

		switch cfg.Driver {
		case DriverMysql:
			if Engine.Mysql == nil {
				Engine.Mysql = &DBManager{}
			}
			mgr = Engine.Mysql

		case DriverPostgres:
			if Engine.Postgres == nil {
				Engine.Postgres = &DBManager{}
			}
			mgr = Engine.Postgres

		case DriverClickHouse:
			if Engine.Clickhouse == nil {
				Engine.Clickhouse = &DBManager{}
			}
			mgr = Engine.Clickhouse

		default:
			panic("gorm: unknown driver")
		}

		if tenantId == tenant.Default {
			if err := db.Use(NewShardingPlugin(cfg.ShardingConf)); err != nil {
				panic(fmt.Errorf("gorm dbresolver error: %w", err))
			}
		}

		mgr.pool.Store(tenantId, db)
	}
}
