package message_center

import "errors"

// 配置验证错误
var (
	ErrAppKeyRequired    = errors.New("AppKey is required")
	ErrAppSecretRequired = errors.New("AppSecret is required")
	ErrHttpURLRequired   = errors.New("HttURL is required")
)

// 错误码
const (
	ErrCodeSignParamsMissing       = 10030 // 签名参数缺失
	ErrCodeSignInvalid             = 10031 // 签名无效
	ErrCodeSignExpired             = 10032 // 签名已过期
	ErrCodeSignDuplicate           = 10033 // 签名重复
	ErrCodeSignError               = 10034 // 签名错误
	ErrCodeSignParamsError         = 10035 // 签名参数错误
	ErrCodeInvalidRequest          = 10036 // 无效请求
	ErrCodeAppKeyNotFound          = 10011 // 租户appKey不存在
	ErrCodeAppKeyDisabled          = 10013 // 租户已禁用
	ErrCodeTopicInvalid            = 10004 // topic必须以appKey开头
	ErrCodeMessageTooLong          = 10020 // 消息内容超过最大长度
	ErrCodeRetainDurationTooLong   = 10027 // 超过最大消息保留时长
	ErrCodeMessageNotFound         = 10022 // 消息不存在
	ErrCodeMessageAlreadySent      = 10024 // 消息已发送
	ErrCodeMessageSendTimeTooSmall = 10026 // 延迟消息发送时间不能小于当前时间
)
