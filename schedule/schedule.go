package schedule

// JobKey 定时任务key 统一定义，防止重复
type JobKey string

func (key JobKey) String() string {
	return string(key)
}

const (
	// JobKeyDemo 测试任务
	JobKeyDemo JobKey = "demo"
	// JobKeySyncGameRecord 同步游戏记录
	JobKeySyncGameRecord JobKey = "sync_game_record"
)

type HandlerName string

func (name HandlerName) String() string {
	return string(name)
}

const (
	// HandlerNameDemo 测试任务
	HandlerNameDemo HandlerName = "demo"
	// HandlerNameSyncGameRecord 同步游戏记录
	HandlerNameSyncGameRecord HandlerName = "sync_game_record_handler"
	// HandlerNameCheckGameTransfer 检查游戏转账
	HandlerNameCheckGameTransfer HandlerName = "check_game_transfer_handler"
)
