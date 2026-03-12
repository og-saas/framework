package utils

import "github.com/dromara/carbon/v2"

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

// GetTimeRange 获取时间范围
func (s TimeRangeType) GetTimeRange() (int64, int64) {
	now := carbon.SetWeekStartsAt(carbon.Monday)
	switch s {
	case TimeRangeTypeToday: // 当日
		return now.StartOfDay().Timestamp(), now.EndOfDay().Timestamp()
	case TimeRangeTypeYesterday: // 昨天
		return now.AddDays(-1).StartOfDay().Timestamp(), now.AddDays(-1).EndOfDay().Timestamp()
	case TimeRangeTypeLastThreeDay: // 最近3天
		return now.AddDays(-3).Timestamp(), now.Timestamp()
	case TimeRangeTypeLastSevenDay: // 最近7天
		return now.AddDays(-7).Timestamp(), now.Timestamp()
	case TimeRangeTypeLastThirtyDay: // 最近30天
		return now.AddDays(-30).Timestamp(), now.Timestamp()
	case TimeRangeTypeThisWeek: // 本周
		return now.StartOfWeek().Timestamp(), now.EndOfWeek().Timestamp()
	case TimeRangeTypeThisMonth: // 本月
		return now.StartOfMonth().Timestamp(), now.EndOfMonth().Timestamp()
	default:
		// 默认返回当日开始，结束时间戳
		return now.StartOfDay().Timestamp(), now.EndOfDay().Timestamp()
	}
}
