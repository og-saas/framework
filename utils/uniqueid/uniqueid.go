package uniqueid

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sony/sonyflake"
)

var (
	flake     *sonyflake.Sonyflake
	startTime = time.Date(1997, 1, 14, 0, 0, 0, 0, time.UTC)
)

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: startTime,
	})
	if flake == nil {
		panic("sony flake init failed")
	}
}

// GenId 生成一个唯一的雪花ID
func GenId() (id uint64, err error) {
	id, err = flake.NextID()
	return
}

func GenOrderNO(prefix string) string {
	id, _ := GenId()
	return fmt.Sprintf("%s-%d", prefix, id)
}

// ExtractTime 从 sonyflake ID 中提取时间戳
func ExtractTime(id uint64) time.Time {
	// sonyflake ID 结构：
	// 39 bits: 时间戳（10ms 精度）
	// 8 bits: sequence number
	// 16 bits: machine id

	// 提取时间戳部分（右移 24 位）
	elapsedTime := id >> 24

	// 转换为实际时间（10ms 单位）
	return startTime.Add(time.Duration(elapsedTime) * 10 * time.Millisecond)
}

// ExtractOrderTime 从订单号中提取时间戳
// orderNo 格式为 "prefix-id"，例如 "ORDER-123456789"
func ExtractOrderTime(orderNo string) (time.Time, error) {
	// 分割订单号，获取 ID 部分
	parts := strings.Split(orderNo, "-")
	if len(parts) < 2 {
		return time.Time{}, fmt.Errorf("invalid order number format: %s", orderNo)
	}

	// 解析 ID
	idStr := parts[len(parts)-1]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse order ID: %w", err)
	}

	// 提取时间戳
	return ExtractTime(id), nil
}
