package consts

// ScheduleJobKey 定时任务key 统一定义，防止重复
type ScheduleJobKey string

func (key ScheduleJobKey) String() string {
	return string(key)
}

const (
	// ScheduleJobKeyDemo 测试任务
	ScheduleJobKeyDemo ScheduleJobKey = "demo"
	// ScheduleJobKeySyncGameRecord 同步游戏记录
	ScheduleJobKeySyncGameRecord ScheduleJobKey = "sync_game_record"
)

type ScheduleHandlerName string

func (name ScheduleHandlerName) String() string {
	return string(name)
}

const (
	// ScheduleHandlerNameDemo 测试任务
	ScheduleHandlerNameDemo ScheduleHandlerName = "demo"
	// ScheduleHandlerNameSyncGameRecord 同步游戏记录
	ScheduleHandlerNameSyncGameRecord ScheduleHandlerName = "sync_game_record_handler"
)
