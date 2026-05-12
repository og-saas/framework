package message_center

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

// WebhookHeaderKey 回调请求头 Key
type WebhookHeaderKey string

const (
	WebhookHeaderSignature WebhookHeaderKey = "X-Webhook-Signature" // 回调签名
	WebhookHeaderAppKey    WebhookHeaderKey = "X-Event-AppKey"      // 租户AppKey
	WebhookHeaderEventType WebhookHeaderKey = "X-Event-Type"        // 事件类型
	WebhookHeaderEventTime WebhookHeaderKey = "X-Event-Time"        // 事件时间戳
)

type Webhook struct {
	Header WebhookHeader
	Body   WebhookBody
}

// WebhookHeader 回调请求头
type WebhookHeader struct {
	Signature   string           // 回调签名
	AppKey      string           // 租户AppKey
	EventType   WebhookEventType // 事件类型
	EventTime   string           // 事件时间戳
	ContentType string           // 内容类型
}

// WebhookBody 回调内容
type WebhookBody struct {
	EventType WebhookEventType `json:"eventType"` // 事件类型
	EventTime int64            `json:"eventTime"` // 事件时间戳
	EventData WebhookEventData `json:"eventData"` // 事件数据
}

// WebhookEventData 事件数据
type WebhookEventData struct {
	ClientId string `json:"clientId"` // 客户端ID
}

func (k WebhookHeaderKey) String() string {
	return string(k)
}

// WebhookEventType 事件类型
type WebhookEventType string

const (
	WebhookEventTypeConnect    WebhookEventType = "connect"    // 连接
	WebhookEventTypeDisconnect WebhookEventType = "disconnect" // 断开连接
)

func (e WebhookEventType) String() string {
	return string(e)
}

// VerifyWebhookSignature 验证 Webhook 签名
func (c *Client) VerifyWebhookSignature(eventType WebhookEventType, eventTime int64, signature string) bool {
	// 构建排序的参数字符串，按照Java代码逻辑
	params := map[string]string{
		"tenantAppKey": c.config.AppKey,
		"eventType":    eventType.String(),
		"eventTime":    cast.ToString(eventTime),
	}

	// 按照键名排序构建查询字符串
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var signString strings.Builder
	for _, k := range keys {
		signString.WriteString(k)
		signString.WriteString("=")
		signString.WriteString(params[k])
		signString.WriteString("&")
	}
	// 添加secret参数
	signString.WriteString("secret=")
	signString.WriteString(c.config.AppSecret)

	// 计算SHA256哈希值
	hash := sha256.Sum256([]byte(signString.String()))
	expectedSignHex := hex.EncodeToString(hash[:])

	//fmt.Println("signString: ", signString.String())
	//fmt.Println("go sign: ", expectedSignHex)
	//fmt.Println("java sign: ", signature)

	return subtle.ConstantTimeCompare([]byte(expectedSignHex), []byte(signature)) == 1
}
