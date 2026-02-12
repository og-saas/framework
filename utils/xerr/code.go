package xerr

//go:generate stringer -type=ErrCode -output=code_string.go --linecomment
type ErrCode int

// 通用错误码
const (
	// ErrCodeSuccess 成功
	ErrCodeSuccess ErrCode = 0 // Success
	// ErrCodeFail 失败
	ErrCodeFail ErrCode = 1 // Fail
	// ErrCodeParamError 参数错误
	ErrCodeParamError ErrCode = 400 // ParamError
	// ErrCodeUnauthorized 未授权
	ErrCodeUnauthorized ErrCode = 401 // Unauthorized
	// ErrCodeForbidden 禁止
	ErrCodeForbidden ErrCode = 403 // Forbidden
	// ErrCodeNotFound 未找到
	ErrCodeNotFound ErrCode = 404 // NotFound
	// ErrCodeServerInternalError 服务器内部错误
	ErrCodeServerInternalError ErrCode = 500 // ServerInternalError
	// ErrCodeServiceUnavailable 服务不可用
	ErrCodeServiceUnavailable ErrCode = 503 // ServiceUnavailable
)

// 用户相关
const (
	// ErrCodeUserNotExists 用户不存在
	ErrCodeUserNotExists ErrCode = 10001 // UserNotExists
	// ErrCodeUserExists 用户已存在
	ErrCodeUserExists ErrCode = 10002 // UserExists
	// ErrCodeUserPwdError 用户密码错误
	ErrCodeUserPwdError ErrCode = 10003 // UserPwdError
	// ErrCodeHadBindError 账号已被绑定
	ErrCodeHadBindError ErrCode = 10004 // HadBindError
	// ErrCodeNotBindError 账号未被绑定
	ErrCodeNotBindError ErrCode = 10005 // NotBindError
	// ErrCodeRepeatUpdateError 账号重复更新
	ErrCodeRepeatUpdateError ErrCode = 10006 // RepeatUpdateError
	// ErrCodeTemporaryTokenEmptyError 账号更新令牌为空
	ErrCodeTemporaryTokenEmptyError ErrCode = 10007 // TemporaryTokenEmptyError
	// ErrCodeTemporaryTokenInvalidError 账号更新令牌无效
	ErrCodeTemporaryTokenInvalidError ErrCode = 10008 // TemporaryTokenInvalidError
	// ErrCodeSelfHadBindError 当前你已绑定该账号
	ErrCodeSelfHadBindError ErrCode = 10009 // SelfHadBindError
	// ErrCodeBindSameError 换绑账号跟老账号相同
	ErrCodeBindSameError ErrCode = 10010 // BindSameError
	// ErrCodeUserAbnormalError 用户账号状态异常
	ErrCodeUserAbnormalError ErrCode = 10011 // UserAbnormalError
	// ErrCodeGoogleAuthError 谷歌校验失败，请重试
	ErrCodeGoogleAuthError ErrCode = 10012 // GoogleAuthError
)

// 游戏相关
const (
	// ErrCodeGameEnterLogExists 用户进入游戏日志已存在
	ErrCodeGameEnterLogExists ErrCode = 20001 // GameEnterLogExists
)

// 财务相关
const (
	// ErrCodeWithdrawBalanceNotEnough 提现余额不足
	ErrCodeWithdrawBalanceNotEnough ErrCode = 30001 // WithdrawBalanceNotEnough
	// ErrCodeWithdrawAmountOutOfRange 提现金额超出范围
	ErrCodeWithdrawAmountOutOfRange ErrCode = 30002 // WithdrawAmountOutOfRange
)

func (code ErrCode) Int() int {
	return int(code)
}

func (code ErrCode) Uint32() uint32 {
	return uint32(code)
}

func (code ErrCode) Error() error {
	return NewError(code, nil)
}
