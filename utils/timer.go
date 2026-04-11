package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"time"

	"github.com/dromara/carbon/v2"
)

// TimeRangeType 1-今天，2-昨天，3-最近3天，4-最近7天，5-最近30天，6-本周，7-本月
type TimeRangeType int

const (
	TimeRangeTypeToday         TimeRangeType = 1 // 今天
	TimeRangeTypeYesterday     TimeRangeType = 2 // 昨天
	TimeRangeTypeLastThreeDay  TimeRangeType = 3 // 最近3天
	TimeRangeTypeLastSevenDay  TimeRangeType = 4 // 最近7天
	TimeRangeTypeLastThirtyDay TimeRangeType = 5 // 最近30天
	TimeRangeTypeThisWeek      TimeRangeType = 6 // 本周
	TimeRangeTypeThisMonth     TimeRangeType = 7 // 本月
)

// GetTimestampRange 获取时间范围
func (s TimeRangeType) GetTimestampRange() (int64, int64) {
	start, end := s.GetTimeRange()
	return start.Timestamp(), end.Timestamp()
}

func (s TimeRangeType) GetTimeRange() (*carbon.Carbon, *carbon.Carbon) {
	now := carbon.Now().SetWeekStartsAt(carbon.Monday)
	switch s {
	case TimeRangeTypeToday: // 当日
		return now.StartOfDay(), now.EndOfDay()
	case TimeRangeTypeYesterday: // 昨天
		return now.AddDays(-1).StartOfDay(), now.AddDays(-1).EndOfDay()
	case TimeRangeTypeLastThreeDay: // 最近 3 天
		return now.AddDays(-3), now
	case TimeRangeTypeLastSevenDay: // 最近 7 天
		return now.AddDays(-7), now
	case TimeRangeTypeLastThirtyDay: // 最近 30 天
		return now.AddDays(-30), now
	case TimeRangeTypeThisWeek: // 本周
		return now.StartOfWeek(), now.EndOfWeek()
	case TimeRangeTypeThisMonth: // 本月
		return now.StartOfMonth(), now.EndOfMonth()
	default:
		// 默认返回当日开始，结束时间
		return now.StartOfDay(), now.EndOfDay()
	}
}

// GetStdTimeRange 获取时间范围，返回标准库 time.Time 类型
func (s TimeRangeType) GetStdTimeRange() (time.Time, time.Time) {
	start, end := s.GetTimeRange()
	return start.StdTime(), end.StdTime()
}

func FormatSeconds(sec int64) string {
	h := sec / 3600
	m := (sec % 3600) / 60
	s := sec % 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// RandomDecimal 在 [min, max] 内随机，结果固定保留2位小数
func RandomDecimal(min, max decimal.Decimal) (decimal.Decimal, error) {
	if min.GreaterThan(max) {
		return decimal.Zero, fmt.Errorf("min > max")
	}

	hundred := decimal.NewFromInt(100)

	// 只允许2位小数范围：min向上取整到分，max向下取整到分
	minCents := min.Mul(hundred).Ceil().IntPart()
	maxCents := max.Mul(hundred).Floor().IntPart()

	if minCents > maxCents {
		return decimal.Zero, fmt.Errorf("no valid 2-decimal number in range")
	}

	diff := maxCents - minCents + 1
	n, err := rand.Int(rand.Reader, big.NewInt(diff))
	if err != nil {
		return decimal.Zero, err
	}

	valCents := minCents + n.Int64()
	return decimal.NewFromInt(valCents).Div(hundred), nil
}
