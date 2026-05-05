package xerr

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/stringx"
)

type Error struct {
	Code ErrCode `json:"code"`
	Data any     `json:"data"`
	Msg  string  `json:"message"`
}

var conf atomic.Value

func Must(c *Config) {
	conf.Store(c)
}

// NewError 业务错误
func NewError(code ErrCode, data any) Error {
	return Error{
		Code: code,
		Data: data,
		Msg:  code.String(),
	}
}

// NewParamError 参数错误
func NewParamError(formatMsg string, data any, formatMsgArgs ...any) Error {
	return Error{
		Code: ErrCodeParamError,
		Data: data,
		Msg:  fmt.Sprintf(formatMsg, formatMsgArgs...),
	}
}

// NewUnauthorizedError 未授权错误
func NewUnauthorizedError() Error {
	return Error{
		Code: ErrCodeUnauthorized,
		Data: nil,
		Msg:  "",
	}
}

// NewForbiddenError 禁止访问错误
func NewForbiddenError(formatMsg string, data any, formatMsgArgs ...any) Error {
	return Error{
		Code: ErrCodeForbidden,
		Data: data,
		Msg:  fmt.Sprintf(formatMsg, formatMsgArgs...),
	}
}

// NewServerInternalError 服务内部错误
func NewServerInternalError(err error) Error {
	if err == nil {
		err = fmt.Errorf("unknown error")
	}
	return Error{
		Code: ErrCodeServerInternalError,
		Data: nil,
		Msg:  err.Error(),
	}
}

// NewServiceUnreachableError 服务不可用错误
func NewServiceUnreachableError(formatMsg string, data any, formatMsgArgs ...any) Error {
	return Error{
		Code: ErrCodeServiceUnavailable,
		Data: data,
		Msg:  fmt.Sprintf(formatMsg, formatMsgArgs...),
	}
}

func (err Error) Error() string {
	return err.Msg
}

// IsXerr 判断 error 是否为 xerr.Error 类型
func IsXerr(err error) bool {
	if err == nil {
		return false
	}
	var e Error
	return errors.As(err, &e)
}

// FromError 从 error 中提取或构造 xerr.Error
// 如果 err 本身是 xerr.Error，直接返回
// 如果 err 是 gRPC 错误（格式：rpc error: code = Unknown desc = ErrCode(xxxxx)），解析错误码并返回
// 否则返回 ServerInternalError
func FromError(err error) Error {
	if err == nil {
		return NewServerInternalError(err)
	}

	// 先尝试用 errors.As 提取
	var e Error
	if errors.As(err, &e) {
		return e
	}

	// 尝试从 gRPC 错误消息中解析错误码
	// 格式：rpc error: code = Unknown desc = ErrCode(30003)
	errMsg := err.Error()
	re := regexp.MustCompile(`ErrCode\((\d+)\)`)
	if matches := re.FindStringSubmatch(errMsg); len(matches) > 1 {
		if code, parseErr := strconv.Atoi(matches[1]); parseErr == nil {
			return Error{
				Code: ErrCode(code),
				Msg:  ErrCode(code).String(),
			}
		}
	}

	// 兜底返回 ServerInternalError
	return NewServerInternalError(err)
}

// GetMessage 获取多语言错误信息
func (err Error) GetMessage(language string) string {
	return TransErrMsg(err.Code.Int(), err.Msg, language)
}

// TransErrMsg 根据错误码和语言获取多语言错误信息
func TransErrMsg(code int, defaultMsg, language string) string {
	// 配置为空，返回默认错误信息
	confVal, ok := conf.Load().(*Config) // 原子读取
	if !ok || confVal == nil {
		return defaultMsg
	}

	if ErrCode(code) == ErrCodeServerInternalError && confVal.ShowServerInternalError {
		return defaultMsg
	}

	// 未找到code，返回默认错误信息
	languageMsgMap, ok := confVal.ErrorMessages[strconv.Itoa(code)]
	if !ok {
		return defaultMsg
	}

	// 语言参数为空，则使用默认语言
	if stringx.HasEmpty(language) {
		language = confVal.LanguageDefault
	}
	// 语言参数还是为空，则使用默认语言
	if stringx.HasEmpty(language) {
		return defaultMsg
	}

	// 返回多语言错误信息
	msg, ok := languageMsgMap[language]
	if ok {
		return msg
	}

	// 该语言未配置错误信息，尝试获取默认语言的错误信息
	if stringx.NotEmpty(confVal.LanguageDefault) && confVal.LanguageDefault != language {
		if msg, ok = languageMsgMap[confVal.LanguageDefault]; ok {
			return msg
		}
	}

	return defaultMsg
}
