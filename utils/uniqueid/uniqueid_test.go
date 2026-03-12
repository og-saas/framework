package uniqueid

import (
	"testing"
	"time"
)

func TestExtractTime(t *testing.T) {
	// 生成一个 ID
	id, err := GenId()
	if err != nil {
		t.Fatalf("GenId failed: %v", err)
	}

	// 提取时间
	extractedTime := ExtractTime(id)

	// 验证提取的时间应该接近当前时间（允许误差 1 秒）
	now := time.Now()
	diff := now.Sub(extractedTime)
	if diff < 0 {
		diff = -diff
	}

	if diff > time.Second {
		t.Errorf("ExtractTime returned time too far from now: got %v, now %v, diff %v", extractedTime, now, diff)
	}
}

func Test_GenOrderNO(t *testing.T) {
	id := GenOrderNO("ORDER")
	t.Log(id)
	tt, er := ExtractOrderTime(id)
	t.Log(tt, er)
}

func TestExtractTime_KnownID(t *testing.T) {
	// 测试已知的 ID 值
	// sonyflake ID 结构：39 bits 时间戳 + 8 bits sequence + 16 bits machine id
	// 假设 elapsedTime = 1000（从 startTime 开始经过 10000ms = 10s）
	// ID = 1000 << 24 = 16777216000
	testID := uint64(1000) << 24

	extractedTime := ExtractTime(testID)
	expectedTime := startTime.Add(10 * time.Second)

	if !extractedTime.Equal(expectedTime) {
		t.Errorf("ExtractTime(%d) = %v, want %v", testID, extractedTime, expectedTime)
	}
}

func TestExtractTime_Zero(t *testing.T) {
	// 测试 ID 为 0 的情况
	extractedTime := ExtractTime(0)

	if !extractedTime.Equal(startTime) {
		t.Errorf("ExtractTime(0) = %v, want %v", extractedTime, startTime)
	}
}

func TestExtractTime_RoundTrip(t *testing.T) {
	// 测试多次生成和提取的一致性
	for i := 0; i < 10; i++ {
		beforeGen := time.Now()

		id, err := GenId()
		if err != nil {
			t.Fatalf("GenId failed on iteration %d: %v", i, err)
		}

		afterGen := time.Now()
		extractedTime := ExtractTime(id)

		// 提取的时间应该在生成前后之间
		if extractedTime.Before(beforeGen.Add(-time.Second)) || extractedTime.After(afterGen.Add(time.Second)) {
			t.Errorf("Iteration %d: extracted time %v not between %v and %v",
				i, extractedTime, beforeGen, afterGen)
		}

		// 避免在同一毫秒内生成多个 ID
		time.Sleep(time.Millisecond)
	}
}
