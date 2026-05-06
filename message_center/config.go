package message_center

import "time"

// Config 消息中心配置
type Config struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	HttpURL   string `json:"http_url"`   // HTTP 地址
	WsURL     string `json:"ws_url"`     // Websocket 地址
	MqttURL   string `json:"mqtt_url"`   // MQTT 地址
	Timeout   int64  `json:"timeout"`    // 请求超时时间，单位：秒，默认 3 秒
	OtpExpire int64  `json:"otp_expire"` // OTP 有效时间，单位：分钟
}

// validate 验证配置
func (c *Config) validate() error {
	if c.AppKey == "" {
		return ErrAppKeyRequired
	}
	if c.AppSecret == "" {
		return ErrAppSecretRequired
	}
	if c.HttpURL == "" {
		return ErrHttpURLRequired
	}
	return nil
}

// getTimeout 获取超时时间
func (c *Config) getTimeout() time.Duration {
	if c.Timeout > 0 {
		return time.Duration(c.Timeout) * time.Second
	}
	return RequestTimeout // 默认 3 秒
}
