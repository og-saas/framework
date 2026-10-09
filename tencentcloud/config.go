package tencentcloud

type Config struct {
	SecretID           string `json:"secret_id,optional"`                           // 腾讯云 Secret ID
	SecretKey          string `json:"secret_key,optional"`                          // 腾讯云 Secret Key
	Token              string `json:"token,optional"`                               // 腾讯云 Token
	ReqTimeout         int    `json:"req_timeout,optional,default=10"`              // 请求超时时间，单位为秒
	EdgeOnePurgeEnable bool   `json:"edge_one_purge_enable,optional,default=false"` // 是否启用EdgeOnePurge
}
