package message_center

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/og-saas/framework/utils"
)

var testConfig = Config{
	AppKey:    "testKey",
	AppSecret: "testSecret",
	HttpURL:   "http://message.ng200.bingthy.xyz",
	Timeout:   5,
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
			client.BuildTopic(TopicGlobalUser),                     // 全局用户消息
			client.BuildTopic(TopicSiteUser, siteId),               // 站点用户消息
			client.BuildTopic(TopicSiteUserSingle, siteId, userId), // 点对点消息
		},
	})
	if err != nil {
		log.Fatalln("Otp error: ", err)
		return
	}
	utils.PrettyJSON(resp)
}

func TestClient_NormalSend(t *testing.T) {
	client, err := NewClient(testConfig)
	if err != nil {
		t.Fatalf("New error: %v", err)
	}
	resp, err := client.Send(context.Background(), SendMessageReq{
		Topic:   "testKey/site/1546549056912754661/device/2a166308-ac41-409e-808a-955295fa420d",
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
	resp, err := client.Send(context.Background(), SendMessageReq{
		Topic:    fmt.Sprintf(TopicSiteUser.String(), siteId),
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
		Topic:     client.BuildTopic(TopicSiteUserSingle, 1001, 123),
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
