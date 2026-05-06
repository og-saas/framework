package xerr

//go:generate stringer -type=ErrCode -output=code_string.go --linecomment
type ErrCode uint32

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
	// ErrCodeThirdPartyAuthError 第三方账号校验失败，请重试
	ErrCodeThirdPartyAuthError ErrCode = 10012 // ThirdPartyAuthError
	// ErrCodeNicknameTimeError 昵称修改时间不满足
	ErrCodeNicknameTimeError ErrCode = 10013 // NicknameTimeError
	// ErrCodeMaxNumberError 最大数量上限
	ErrCodeMaxNumberError ErrCode = 10014 // MaxNumberError
	// ErrCodeBalanceRetrievingError 余额找回中，请稍后
	ErrCodeBalanceRetrievingError ErrCode = 10015 // BalanceRetrievingError
	// ErrCodeNicknameExistError 昵称已存在
	ErrCodeNicknameExistError ErrCode = 10016 // NicknameExistError
	// ErrCodeOldPasswordError 旧密码错误
	ErrCodeOldPasswordError ErrCode = 10017 // OldPasswordError
	// ErrCodeLoginLocked 登录失败次数过多，账号已锁定
	ErrCodeLoginLocked ErrCode = 10018 // LoginLocked
	// ErrCodeSetPinLockedFirst 设置PIN码失败，账号已锁定（首次锁定）
	ErrCodeSetPinLockedFirst ErrCode = 10019 // SetPinLockedFirst
	// ErrCodeOldPinError 旧PIN码错误 带剩余次数
	ErrCodeOldPinError ErrCode = 10020 // OldPinError
	// ErrCodeSetPinLockedMax 设置PIN码失败，账号已锁定（最大锁定）
	ErrCodeSetPinLockedMax ErrCode = 10022 // SetPinLockedMax
	// ErrCodeWithdrawAccountNotExists 提现账号不存在
	ErrCodeWithdrawAccountNotExists ErrCode = 10023 // WithdrawAccountNotExists
	// ErrCodeWithdrawAccountNotSetPin 提现账号未设置PIN码
	ErrCodeWithdrawAccountNotSetPin ErrCode = 10024 // WithdrawAccountNotSetPin
	// ErrCodeSetPinCaptchaNumError 设置PIN码验证码次数达到上限错误
	ErrCodeSetPinCaptchaNumError ErrCode = 10025 // SetPinCaptchaNumError
	// ErrCodeUserStatusForbidden 用户状态异常禁止操作错误码
	ErrCodeUserStatusForbidden ErrCode = 10026 // UserStatusForbidden

)

// 游戏相关
const (
	// ErrCodeGameEnterLogExists 用户进入游戏日志已存在
	ErrCodeGameEnterLogExists ErrCode = 20001 // GameEnterLogExists
	// ErrCodeGameEnterBalanceNotEnough 用户进入游戏余额不足
	ErrCodeGameEnterBalanceNotEnough ErrCode = 20002 // ErrCodeGameEnterLogExists
	// ErrCodeGamePlatformUnreachable 游戏中台请求未到达（被拦截/网络不可达）
	ErrCodeGamePlatformUnreachable ErrCode = 20003 // GamePlatformUnreachable
)

// 财务相关
const (
	// ErrCodeWithdrawBalanceNotEnough 提现余额不足
	ErrCodeWithdrawBalanceNotEnough ErrCode = 30001 // WithdrawBalanceNotEnough
	// ErrCodeWithdrawAmountOutOfRange 提现金额超出范围
	ErrCodeWithdrawAmountOutOfRange ErrCode = 30002 // WithdrawAmountOutOfRange
	// ErrCodeRechargeAmountOutOfChannelDayMax 充值金额超出渠道每日最大限制
	ErrCodeRechargeAmountOutOfChannelDayMax ErrCode = 30003 // RechargeAmountOutOfChannelDayMax
	// ErrCodeWithdrawLocked 提现功能已被锁定 (设置PIN码超次数，提现输入超次数)
	ErrCodeWithdrawLocked ErrCode = 30004 // WithdrawLocked
	// ErrCodeWithdrawPinCheckError PIN码错误(发起提现) 带剩余次数
	ErrCodeWithdrawPinCheckError ErrCode = 30005 // WithdrawPinCheckError
	// ErrCodeUserWithdrawForbidden 用户禁止提现
	ErrCodeUserWithdrawForbidden ErrCode = 30006 // UserWithdrawForbidden
	// ErrCodeWithdrawLockedFirst 提现PIN码错误，账号已锁定（首次锁定）
	ErrCodeWithdrawLockedFirst ErrCode = 30007 // WithdrawLockedFirst
)

// 代理相关
const (
	// ErrCodeInvalidPromotionCode 无效的推广码
	ErrCodeInvalidPromotionCode ErrCode = 40000 // ErrCodeInvalidPromotionCode
	// ErrCodeBindParentNotAllowed 不允许绑定上级
	ErrCodeBindParentNotAllowed ErrCode = 40001 // ErrCodeBindParentNotAllowed
)

// share相关
const (
	// ErrCodeCaptchaErr 验证码错误
	ErrCodeCaptchaErr ErrCode = 50001 // ErrCodeCaptchaErr
)

// 优惠相关
const (
	// ErrCodeActivityNotStart 活动还未开始
	ErrCodeActivityNotStart ErrCode = 60001 // ErrCodeActivityNotStart
	// ErrCodeActivityEnded 活动已结束
	ErrCodeActivityEnded ErrCode = 60002 // ErrCodeActivityEnded
	// ErrCodeActivityClosed 活动已关闭
	ErrCodeActivityClosed ErrCode = 60003 // ErrCodeActivityClosed
	// ErrCodeRewardNotAvailableYet 奖励未到领取时间
	ErrCodeRewardNotAvailableYet ErrCode = 60004 // ErrCodeRewardNotAvailableYet
	// ErrCodeRewardNotBelongToYou 奖励不属于你
	ErrCodeRewardNotBelongToYou ErrCode = 60005 // ErrCodeRewardNotBelongToYou
	// ErrCodeRewardConditionNotMet 未达到领取奖励的条件
	ErrCodeRewardConditionNotMet ErrCode = 60006 // ErrCodeRewardConditionNotMet
	// ErrCodeRewardNotClaimable 奖励不可领取
	ErrCodeRewardNotClaimable ErrCode = 60007 // ErrCodeRewardNotClaimable
	// ErrCodeRewardExpired 奖励已过期
	ErrCodeRewardExpired ErrCode = 60008 // ErrCodeRewardExpired
	// ErrCodeClaimRewardSmsLimit 奖励领取短信未验证
	ErrCodeClaimRewardSmsLimit ErrCode = 60009 // ErrCodeClaimRewardSmsLimit
	// ErrCodeClaimRewardEmailLimit 奖励领取邮件未验证
	ErrCodeClaimRewardEmailLimit ErrCode = 60010 // ErrCodeClaimRewardEmailLimit
	// ErrCodeClaimRewardSameActivityTypeLimit 已领取同类活动奖励
	ErrCodeClaimRewardSameActivityTypeLimit ErrCode = 60011 // ErrCodeClaimRewardSameActivityTypeLimit
	// ErrCodeClaimRewardIPLimit 活动IP风控限制
	ErrCodeClaimRewardIPLimit ErrCode = 60012 // ErrCodeClaimRewardIPLimit
	// ErrCodeClaimRewardDeviceLimit 活动设备风控限制
	ErrCodeClaimRewardDeviceLimit ErrCode = 60013 // ErrCodeClaimRewardDeviceLimit
	// ErrCodeClaimRewardEndpointLimit 领取终端限制
	ErrCodeClaimRewardEndpointLimit ErrCode = 60014 // ErrCodeClaimRewardEndpointLimit
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
