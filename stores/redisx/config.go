package redisx

// Config redis配置
type Config struct {
	Addrs        []string `json:"addrs"`
	Debug        bool     `json:"debug,default=false"`
	Trace        bool     `json:"trace,default=false"`
	MasterName   string   `json:"master_name,optional"`
	Username     string   `json:"username,optional"`
	Password     string   `json:"password,optional"`
	DB           int      `json:"db,default=0"`
	IsCluster    bool     `json:"is_cluster,optional"`
	PoolSize     int      `json:"pool_size,optional"`
	MinIdleConns int      `json:"min_idle_conns,optional"`
	DialTimeout  int      `json:"dial_timeout,optional"`  // milliseconds
	ReadTimeout  int      `json:"read_timeout,optional"`  // milliseconds
	WriteTimeout int      `json:"write_timeout,optional"` // milliseconds
}

type TenantConfig struct {
	Default Config            `json:"default"`
	Tenants map[string]Config `json:"tenants,optional"`
}
