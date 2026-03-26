package gormx

import (
	"errors"
	"fmt"
	"sync"

	"github.com/og-saas/framework/utils/tenant"
	"gorm.io/gorm"
)

var ShardingErr = errors.New("site_id is empty")

var (
	shardingTables sync.Map // map[string]int
)

type ShardingPlugin struct {
	tables      map[string]int
	initialized bool
}

func NewShardingPlugin(conf []ShardingInfo) *ShardingPlugin {
	p := &ShardingPlugin{
		tables: make(map[string]int),
	}
	for _, item := range conf {
		for _, t := range item.Tables {
			p.tables[t] = item.Count
		}
	}
	for k, v := range p.tables {
		shardingTables.Store(k, v)
	}
	return p
}

func (p *ShardingPlugin) Name() string {
	return "gorm:sharding_plugin"
}

func (p *ShardingPlugin) Initialize(db *gorm.DB) error {
	if p.initialized {
		return errors.New("sharding plugin already initialized")
	}
	p.initialized = true

	db.Callback().Create().Before("*").Register("sharding_plugin:before_create", p.beforeCallback)
	db.Callback().Query().Before("*").Register("sharding_plugin:before_query", p.beforeCallback)
	db.Callback().Update().Before("*").Register("sharding_plugin:before_update", p.beforeCallback)
	db.Callback().Delete().Before("*").Register("sharding_plugin:before_delete", p.beforeCallback)
	db.Callback().Row().Before("*").Register("sharding_plugin:before_row", p.beforeCallback)
	db.Callback().Raw().Before("*").Register("sharding_plugin:before_raw", p.beforeCallback)

	return nil
}

// beforeCallback 在 SQL 执行前修改表名
func (p *ShardingPlugin) beforeCallback(db *gorm.DB) {
	table := db.Statement.Table

	// 如果 Table 为空，尝试从 Model 解析表名
	if table == "" && db.Statement.Model != nil {
		if err := db.Statement.Parse(db.Statement.Model); err == nil && db.Statement.Schema != nil {
			table = db.Statement.Schema.Table
		}
	}

	if table == "" {
		_ = db.AddError(ShardingErr)
		return
	}

	count := p.tables[table]
	if count <= 0 {
		return
	}

	ctx := db.Statement.Context
	siteId := tenant.GetTenantId(ctx)
	if siteId <= 0 { // 如果配置分表 但是没有 siteId 返回错误
		_ = db.AddError(ShardingErr)
		return
	}

	// 修改 Statement 中的表名
	db.Statement.Table = fmt.Sprintf("%s_%d", table, siteId%int64(count))
}

// ShardingSuffix 获取分表
func ShardingSuffix(siteId int64, table string) string {
	if c, ok := shardingTables.Load(table); ok {
		return fmt.Sprintf("%s_%d", table, siteId%int64(c.(int)))
	}
	return table
}
