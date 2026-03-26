package gormx

import (
	"errors"
	"fmt"

	"github.com/og-saas/framework/utils/tenant"
	"gorm.io/gorm"
)

var ShardingErr = errors.New("site_id is empty")

var tables = make(map[string]int)

type ShardingPlugin struct {
	tables map[string]int
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
	tables = p.tables
	return p
}

func (p *ShardingPlugin) Name() string {
	return "gorm:sharding_plugin"
}

func (p *ShardingPlugin) Initialize(db *gorm.DB) error {
	// 注册所有需要的回调
	for _, op := range []string{"create", "query", "update", "delete", "row", "raw"} {
		switch op {
		case "create":
			db.Callback().Create().Before("*").Register("sharding_plugin:before_create", p.beforeCallback)
		case "query":
			db.Callback().Query().Before("*").Register("sharding_plugin:before_query", p.beforeCallback)
		case "update":
			db.Callback().Update().Before("*").Register("sharding_plugin:before_update", p.beforeCallback)
		case "delete":
			db.Callback().Delete().Before("*").Register("sharding_plugin:before_delete", p.beforeCallback)
		case "row":
			db.Callback().Row().Before("*").Register("sharding_plugin:before_row", p.beforeCallback)
		case "raw":
			db.Callback().Raw().Before("*").Register("sharding_plugin:before_raw", p.beforeCallback)
		}
	}
	return nil
}

// beforeCallback 在 SQL 执行前修改表名
func (p *ShardingPlugin) beforeCallback(db *gorm.DB) {

	table := db.Statement.Table
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
func ShardingSuffix(siteId int, table string) string {
	if c, ok := tables[table]; ok {
		return fmt.Sprintf("%s_%d", table, siteId%c)
	}
	return table
}
