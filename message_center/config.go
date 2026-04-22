package message_center

import "time"

// Config 消息中心配置
type Config struct {
	AppKey    string        `json:"app_key"`
	AppSecret string        `json:"app_secret"`
	BaseURL   string        `json:"base_url"` // HTTP API 地址
	Timeout   time.Duration `json:"timeout"`  // 请求超时时间，默认 3 秒
}

// validate 验证配置
func (c *Config) validate() error {
	if c.AppKey == "" {
		return ErrAppKeyRequired
	}
	if c.AppSecret == "" {
		return ErrAppSecretRequired
	}
	if c.BaseURL == "" {
		return ErrBaseURLRequired
	}
	return nil
}

// getTimeout 获取超时时间
func (c *Config) getTimeout() time.Duration {
	if c.Timeout > 0 {
		return c.Timeout
	}
	return RequestTimeout // 默认 3 秒
}
