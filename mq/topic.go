package mq

import (
	"os"
)

var (
	// TopicUserWalletTransferNotify 用户钱包交易通知
	TopicUserWalletTransferNotify = "user_wallet_transfer_notify"
	// TopicUserLoginLogNotify 用户登录日志通知
	TopicUserLoginLogNotify = "user_login_log_notify"
	// TopicGameBetRecordNotify 游戏下注记录通知
	TopicGameBetRecordNotify = "game_bet_record_notify"
	// TopicRechargeOrderNotify 充值订单通知
	TopicRechargeOrderNotify = "recharge_order_notify"
	// TopicAgentGradeGrowthNotify 代理等级成长消息通知
	TopicAgentGradeGrowthNotify = "agent_grade_growth_notify"
)

func UpdateTopicPrefix(prefixes ...string) (prefix string) {
	if len(prefixes) > 0 && prefixes[0] != "" {
		prefix = prefixes[0]
	}
	if prefix == "" {
		prefix = os.Getenv("ROCKETMQ_TOPIC_PREFIX")
	}

	TopicUserWalletTransferNotify = prefix + TopicUserWalletTransferNotify
	TopicUserLoginLogNotify = prefix + TopicUserLoginLogNotify
	TopicGameBetRecordNotify = prefix + TopicGameBetRecordNotify
	TopicRechargeOrderNotify = prefix + TopicRechargeOrderNotify
	TopicAgentGradeGrowthNotify = prefix + TopicAgentGradeGrowthNotify

	return
}
