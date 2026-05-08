package message_center

type MessageCode int

const (
	MessageCodeUserLogin           MessageCode = 1001 // 用户登录
	MessageCodeUserRechargeSuccess MessageCode = 1002 // 用户充值成功
)

type Message struct {
	Code    MessageCode `json:"code"`    // 消息码
	Content string      `json:"content"` // 消息内容
}
