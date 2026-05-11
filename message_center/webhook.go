package message_center

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
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
	TenantId       int64  `json:"tenantId"`       // 租户ID
	ActualClientId string `json:"actualClientId"` // 客户端ID
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
	signString := fmt.Sprintf("tenantAppKey=%s&eventType=%s&eventTime=%d&secret=%s",
		c.config.AppKey, eventType, eventTime, c.config.AppSecret)
	expectedSign := sha256.Sum256([]byte(signString))
	expectedSignHex := hex.EncodeToString(expectedSign[:])

	return subtle.ConstantTimeCompare([]byte(expectedSignHex), []byte(signature)) == 1
}
