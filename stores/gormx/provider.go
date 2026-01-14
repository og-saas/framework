package gormx

import "github.com/spf13/cast"

type TenantConfigProvider interface {
	Load() (map[int64]Config, error)
}

type ConfigTenantProvider struct {
	configs map[string]Config
}

func NewConfigTenantProvider(config map[string]Config) *ConfigTenantProvider {
	return &ConfigTenantProvider{
		configs: config,
	}
}

func (p *ConfigTenantProvider) Load() (map[int64]Config, error) {
	configs := make(map[int64]Config)
	for k, v := range p.configs {
		tenantId, err := cast.ToInt64E(k)
		if err != nil {
			return nil, err
		}
		configs[tenantId] = v
	}
	return configs, nil
}
