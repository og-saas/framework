package message_center

type CommonResp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// OtpReq 获取连接凭证请求
type OtpReq struct {
	ClientId       string         `json:"clientId"`       // 客户端唯一标识
	Topics         []string       `json:"topics"`         // 可订阅的Topic列表
	Expire         int64          `json:"expire"`         // OTP有效时间（分钟）
	ConnectionType ConnectionType `json:"connectionType"` // 连接类型 1:mqtt 2:ws 默认1
}

// otpReqInternal 内部请求（包含AppKey）
type otpReqInternal struct {
	AppKey         string         `json:"appKey"`         // 租户应用Key
	ClientId       string         `json:"clientId"`       // 客户端唯一标识
	Topics         []string       `json:"topics"`         // 可订阅的Topic列表
	Expire         int64          `json:"expire"`         // OTP有效时间（分钟）
	ConnectionType ConnectionType `json:"connectionType"` // 连接类型 1:mqtt 2:ws 默认1
}

// OtpResp 获取连接凭证响应
type OtpResp struct {
	ClientId      string `json:"clientId"`      // 实际连接用的客户端ID
	UserName      string `json:"userName"`      // 连接用户名（即appKey）
	BrokerAddress string `json:"brokerAddress"` // 连接地址
}

// SendMessageReq 发送消息请求
type SendMessageReq struct {
	Topic              string `json:"topic"`                        // 消息接收Topic
	Content            string `json:"content"`                      // 消息内容
	Qos                QoS    `json:"qos,omitempty"`                // 消息等级，默认0
	Retain             Retain `json:"retain,omitempty"`             // 是否保留，默认0
	RetainTimeDuration int64  `json:"retainTimeDuration,omitempty"` // 保留时长（秒），retain=1时必填
	SendTime           int64  `json:"sendTime,omitempty"`           // 发送时间（毫秒时间戳，0时区）, 定时消息比传此参数
}

// sendMessageReqInternal 内部请求（包含AppKey）
type sendMessageReqInternal struct {
	AppKey             string `json:"appKey"`                       // 租户应用Key
	Topic              string `json:"topic"`                        // 消息接收Topic
	Content            string `json:"content"`                      // 消息内容
	Qos                QoS    `json:"qos,omitempty"`                // 消息等级，默认0
	Retain             Retain `json:"retain,omitempty"`             // 是否保留，默认0
	RetainTimeDuration int64  `json:"retainTimeDuration,omitempty"` // 保留时长（秒），retain=1时必填
}

// SendMessageResp 发送消息响应
type SendMessageResp struct {
	MessageId string `json:"messageId"` // 消息唯一ID
}

// sendTimerMessageReqInternal 内部请求（包含AppKey）
type sendTimerMessageReqInternal struct {
	AppKey             string `json:"appKey"`                       // 租户应用Key
	Topic              string `json:"topic"`                        // 消息接收Topic
	Content            string `json:"content"`                      // 消息内容
	Qos                QoS    `json:"qos,omitempty"`                // 消息等级，默认0
	Retain             Retain `json:"retain,omitempty"`             // 是否保留，默认0
	RetainTimeDuration int64  `json:"retainTimeDuration,omitempty"` // 保留时长（秒）
	SendTime           int64  `json:"sendTime"`                     // 发送时间（毫秒时间戳，0时区）
}

// CancelTimerMessageReq 取消定时消息请求
type CancelTimerMessageReq struct {
	MessageId string `json:"messageId"` // 定时消息ID
}

// cancelTimerMessageReqInternal 内部请求（包含AppKey）
type cancelTimerMessageReqInternal struct {
	AppKey    string `json:"appKey"`    // 租户应用Key
	MessageId string `json:"messageId"` // 定时消息ID
}

// CancelTimerMessageResp 取消定时消息响应
type CancelTimerMessageResp struct{}

// QueryHistoryReq 历史消息查询请求
type QueryHistoryReq struct {
	Topic     string `json:"topic,omitempty"`     // 查询的topic
	StartTime int64  `json:"startTime,omitempty"` // 查询时间范围起始时间（毫秒时间戳），默认为系统部署时间
	EndTime   int64  `json:"endTime,omitempty"`   // 查询时间范围结束时间（毫秒时间戳），默认为当前时间
	Limit     int    `json:"limit,omitempty"`     // 查询条数，范围1-100，默认100
	Cursor    string `json:"cursor,omitempty"`    // 翻页游标，格式：sendTime:normalMessageId
}

// QueryHistoryResp 历史消息查询响应
type QueryHistoryResp struct {
	Items      []HistoryMessage `json:"items"`      // 历史消息列表（按发送时间倒序排列）
	HasData    bool             `json:"hasData"`    // 是否还有下一页数据
	NextCursor string           `json:"nextCursor"` // 下一页游标
}

// HistoryMessage 历史消息
type HistoryMessage struct {
	Topic    string `json:"topic"`    // 消息topic
	SendTime int64  `json:"sendTime"` // 消息发送时间（毫秒时间戳）
	Content  string `json:"content"`  // 消息内容
}
