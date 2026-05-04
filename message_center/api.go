package message_center

import "time"

// API 路径
const (
	OtpURL          = "/tenant/auth/otp"              // 获取连接凭证
	NormalSendURL   = "/tenant/message/normal/send"   // 发送即时消息
	TimerSendURL    = "/tenant/message/timer/send"    // 发送定时消息
	TimerCancelURL  = "/tenant/message/timer/cancel"  // 取消定时消息
	HistoryQueryURL = "/tenant/message/history/query" // 查询历史消息
)

// RequestTimeout 超时时间
const RequestTimeout = 3 * time.Second

// 请求头 Key
const (
	HeaderAppKey    = "appKey"    // 应用Key
	HeaderTimestamp = "timestamp" // 时间戳
	HeaderNonce     = "nonce"     // 随机字符串
	HeaderSign      = "sign"      // 签名
)

// QoS 等级
type QoS int

const (
	QosAtMostOnce  QoS = iota // 最多一次
	QosAtLeastOnce            // 至少一次
	QosExactlyOnce            // 仅一次
)

func (q QoS) Int() int {
	return int(q)
}

// Retain 状态
type Retain int

const (
	RetainDisabled Retain = iota // 普通消息
	RetainEnabled                // 保留消息
)

func (r Retain) Int() int {
	return int(r)
}

type ClientPrefix string

const (
	ClientPrefixUser  = "user_"
	ClientPrefixAdmin = "admin_"
)
