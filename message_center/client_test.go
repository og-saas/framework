package message_center

import (
	"context"
	"testing"
	"time"

	"github.com/og-saas/framework/utils"
)

var testConfig = Config{
	AppKey:    "testKey",
	AppSecret: "testSecret",
	BaseURL:   "http://message.ng200.bingthy.xyz",
}

func TestNew(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
}

func TestClient_BuildTopic(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	tests := []struct {
		name     string
		topic    string
		expected string
	}{
		{"全局用户", client.BuildGlobalUserTopic(), "testKey/global/user"},
		{"全局管理员", client.BuildGlobalAdminTopic(), "testKey/global/admin"},
		{"站点用户", client.BuildSiteUserTopic(1001), "testKey/site/1001/user"},
		{"站点指定用户", client.BuildSiteUserSingleTopic(1001, 123), "testKey/site/1001/user/123"},
		{"站点用户标签", client.BuildSiteUserTagTopic(1001, "vip"), "testKey/site/1001/tag/vip"},
		{"站点用户渠道", client.BuildSiteUserChannelTopic(1001, 100), "testKey/site/1001/channel/100"},
		{"站点管理员", client.BuildSiteAdminTopic(1001), "testKey/site/1001/admin"},
		{"站点指定管理员", client.BuildSiteAdminSingleTopic(1001, "admin456"), "testKey/site/1001/admin/admin456"},
		{"自定义消息", client.BuildCustomTopic("chat/room001"), "testKey/custom/chat/room001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.topic != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, tt.topic)
			} else {
				t.Logf("✓ %s: %s", tt.name, tt.topic)
			}
		})
	}
}

func TestClient_Otp(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	siteId := int64(1001)
	userId := int64(123)
	resp, err := client.Otp(context.Background(), OtpReq{
		ClientId: "user123",
		Topics: []string{
			client.BuildGlobalUserTopic(),                   // 全局用户消息
			client.BuildSiteUserTopic(siteId),               // 站点用户消息
			client.BuildSiteUserSingleTopic(siteId, userId), // 点对点消息
		},
	})
	if err != nil {
		t.Logf("Otp error: %v", err)
		return
	}
	utils.PrettyJSON(resp)
}

func TestClient_NormalSend(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	siteId := int64(1001)
	userId := int64(123)
	resp, err := client.NormalSend(context.Background(), SendMessageReq{
		Topic:   client.BuildSiteUserSingleTopic(siteId, userId),
		Content: `{"text":"Hello"}`,
		Qos:     QosAtMostOnce,
	})
	if err != nil {
		t.Logf("NormalSend error: %v", err)
		return
	}
	utils.PrettyJSON(resp)
}

func TestClient_TimerSend(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	siteId := int64(1001)
	resp, err := client.TimerSend(context.Background(), SendTimerMessageReq{
		Topic:    client.BuildSiteUserTopic(siteId),
		Content:  `{"text":"Meeting will start"}`,
		SendTime: time.Now().Add(10 * time.Second).UnixMilli(), // 10秒后发送
	})
	if err != nil {
		t.Logf("TimerSend error: %v", err)
		return
	}
	utils.PrettyJSON(resp)
}

func TestClient_CancelTimer(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	// 需要先发送定时消息获取 messageId
	resp, err := client.CancelTimer(context.Background(), CancelTimerMessageReq{
		MessageId: "T1234567890123456789",
	})
	if err != nil {
		t.Logf("CancelTimer error: %v", err)
		return
	}
	utils.PrettyJSON(resp)
}

func TestClient_QueryHistory(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	resp, err := client.QueryHistory(context.Background(), QueryHistoryReq{
		Topic:     client.BuildSiteUserSingleTopic(1001, 123),
		StartTime: 1773964800974,
		EndTime:   0,
		Limit:     10,
	})
	if err != nil {
		t.Logf("QueryHistory error: %v", err)
		return
	}
	utils.PrettyJSON(resp)
}
