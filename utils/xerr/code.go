package xerr

import (
	"fmt"
)

type ErrCode int

const (
	// 通用错误码
	ErrCodeSuccess             ErrCode = 0   // 成功
	ErrCodeFail                ErrCode = 1   // 失败
	ErrCodeParamError          ErrCode = 400 // 参数错误
	ErrCodeUnauthorized        ErrCode = 401 // 未授权
	ErrCodeForbidden           ErrCode = 403 // 禁止
	ErrCodeNotFound            ErrCode = 404 // 未找到
	ErrCodeServerInternalError ErrCode = 500 // 服务器内部错误
	ErrCodeServiceUnavailable  ErrCode = 503 // 服务不可用

	// 用户相关
	ErrCodeUserNotExists              ErrCode = 10001 // 用户不存在
	ErrCodeUserExists                 ErrCode = 10002 // 用户已存在
	ErrCodeUserPwdError               ErrCode = 10003 // 用户密码错误
	ErrCodeHadBindError               ErrCode = 10004 // 账号已被绑定
	ErrCodeNotBindError               ErrCode = 10005 // 账号未被绑定
	ErrCodeRepeatUpdateError          ErrCode = 10006 // 账号重复更新
	ErrCodeTemporaryTokenEmptyError   ErrCode = 10007 // 账号更新令牌为空
	ErrCodeTemporaryTokenInvalidError ErrCode = 10008 // 账号更新令牌无效
	ErrCodeSelfHadBindError           ErrCode = 10009 // 当前你已绑定该账号
	ErrCodeBindSameError              ErrCode = 10010 // 换绑账号跟老账号相同
	ErrCodeUserAbnormalError          ErrCode = 10011 // 用户账号状态异常

	// 游戏相关
	ErrCodeGameEnterLogExists ErrCode = 20001 // 用户进入游戏日志已存在
)

func (s ErrCode) Int() int {
	return int(s)
}

func (s ErrCode) String() string {
	return fmt.Sprintf("%d", int(s))
}
