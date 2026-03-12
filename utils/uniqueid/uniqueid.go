package uniqueid

import (
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
