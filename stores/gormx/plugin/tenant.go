package plugin

import (
	"fmt"
	"github.com/og-saas/framework/utils/tenant"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

const (
	MultiTenantPluginName       = "multi_tenant_plugin"
	MultiTenantPluginQueryName  = "multi_tenant_plugin:query"
	MultiTenantPluginCreateName = "multi_tenant_plugin:create"
	MultiTenantPluginUpdateName = "multi_tenant_plugin:update"
	MultiTenantPluginDeleteName = "multi_tenant_plugin:delete"
)

// MultiTenantPlugin
//
// 插入时，不处理租户字段 仅支持单表查询、更新、删除
type MultiTenantPlugin struct {
	TenantDBName string
}

// NewMultiTenantPlugin 创建插件实例
func NewMultiTenantPlugin(tenantDbName string) *MultiTenantPlugin {
	return &MultiTenantPlugin{
		TenantDBName: tenantDbName,
	}
}

func (p *MultiTenantPlugin) Initialize(db *gorm.DB) error {
	var err error
	if err = db.Callback().Query().Before("gorm:query").Register(MultiTenantPluginQueryName, p.queryHook); err != nil {
		return err
	}
	//if err = db.Callback().Create().Before("gorm:create").Register(MultiTenantPluginCreateName, p.createHook); err != nil {
	//	return err
	//}
	if err = db.Callback().Update().Before("gorm:update").Register(MultiTenantPluginUpdateName, p.updateHook); err != nil {
		return err
	}
	if err = db.Callback().Delete().Before("gorm:delete").Register(MultiTenantPluginDeleteName, p.deleteHook); err != nil {
		return err
	}
	return nil
}

func (p *MultiTenantPlugin) Name() string {
	return MultiTenantPluginName
}

// ==================== Hook 核心 ====================

// Create Hook
//func (p *MultiTenantPlugin) createHook(db *gorm.DB) {
//
//	var field *schema.Field
//	if p.isSkipTenant(db) {
//		return
//	}
//
//	if field = p.getMultiTenantField(db); field == nil {
//		return
//	}
//
//	tenantId := tenant.GetTenantId(db.Statement.Context)
//	reflectValue := db.Statement.ReflectValue
//	switch reflectValue.Kind() {
//	case reflect.Slice, reflect.Array:
//		for i := 0; i < reflectValue.Len(); i++ {
//			_ = field.Set(db.Statement.Context, reflectValue.Index(i), tenantId)
//		}
//	default:
//		_ = field.Set(db.Statement.Context, reflectValue, tenantId)
//	}
//}

// queryHook
func (p *MultiTenantPlugin) queryHook(db *gorm.DB) {
	p.tenantScope(db)
}

// updateHook
func (p *MultiTenantPlugin) updateHook(db *gorm.DB) {
	p.tenantScope(db)
}

// deleteHook
func (p *MultiTenantPlugin) deleteHook(db *gorm.DB) {
	p.tenantScope(db)
}

// 检查是否跳过
func (p *MultiTenantPlugin) isSkipTenant(db *gorm.DB) bool {
	if db.Statement.Schema == nil {
		return true
	}

	if tenant.IsSkipTenant(db.Statement.Context) {
		return true
	}
	return false
}

// 获取租户字段
func (p *MultiTenantPlugin) getMultiTenantField(db *gorm.DB) *schema.Field {
	return db.Statement.Schema.FieldsByDBName[p.TenantDBName]
}

func (p *MultiTenantPlugin) tenantScope(db *gorm.DB) {
	var field *schema.Field

	if p.isSkipTenant(db) {
		return
	}
	if field = p.getMultiTenantField(db); field == nil {
		return
	}

	db.Statement.AddClause(clause.Where{
		Exprs: []clause.Expression{
			clause.Eq{Column: fmt.Sprintf("%s.%s", db.Statement.Table, field.DBName), Value: tenant.GetTenantId(db.Statement.Context)},
		},
	})
}
