package utils

import (
	"fmt"
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
