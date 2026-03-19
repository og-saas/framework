package xerr

type Config struct {
	LanguageDefault         string                       `json:"language_default,optional"`           // 默认语言
	ErrorMessages           map[string]map[string]string `json:"error_messages,optional"`             // 多语言错误信息
	ShowServerInternalError bool                         `json:"show_server_internal_error,optional"` // 是否显示服务内部错误
}
