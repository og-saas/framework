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
	// TopicAgentBetRebatNotify 代理返佣投注更新通知
	TopicAgentBetRebatNotify = "agent_bet_rebat_notify"
	// TopicUserRegisterNotify 用户注册通知
	TopicUserRegisterNotify = "user_register_notify"
	// TopicUserWithdrawCreateNotify 用户发起提现通知
	TopicUserWithdrawCreateNotify = "user_withdraw_create_notify"
	// TopicWithdrawOrderNotify 提现订单回调通知
	TopicWithdrawOrderNotify = "withdraw_order_notify"
	// TopicUserRiskMonitorNotify 用户行为风控监控通知
	TopicUserRiskMonitorNotify = "user_risk_monitor_notify"
	// TopicUserVipLevelChangeNotify VIP等级变化通知
	TopicUserVipLevelChangeNotify = "user_vip_level_change_notify"
	// TopicJackpotMatchRankNotify 锦标赛达成指定名次通知
	TopicJackpotMatchRankNotify = "jackpot_match_rank_notify"
	// TopicUserPageLoginNotify 登录指定页面通知
	TopicUserPageLoginNotify = "user_page_login_notify"
	// TopicUserActivityCompleteNotify 指定活动全部完成通知
	TopicUserActivityCompleteNotify = "user_activity_complete_notify"
	// TopicWebsocketOnlineNotify websocket上线通知
	TopicWebsocketOnlineNotify = "websocket_online_notify"

	// TopicSiteMsgVipNotifyOrdered VIP消息通知（含类型枚举区分升级/奖励创建/奖励发放） 有序消息
	TopicSiteMsgVipNotifyOrdered = "site_msg_vip_notify_ordered"
	// TopicSiteMsgActivityReward 活动奖励消息（奖励可领取/奖励发放成功）
	TopicSiteMsgActivityReward = "site_msg_activity_reward"
	// TopicSiteMsgActivityJackpot 活动Jackpot触发
	TopicSiteMsgActivityJackpot = "site_msg_activity_jackpot"
	// TopicSiteMsgActivitySchedule 活动定时调度（开始通知/结束前提醒）
	TopicSiteMsgActivitySchedule = "site_msg_activity_schedule"
	// TopicSiteMsgWithdrawAudit 提现审核结果通知
	TopicSiteMsgWithdrawAudit = "site_msg_withdraw_audit"
	// TopicSiteMsgRechargeFail 充值失败通知
	TopicSiteMsgRechargeFail = "site_msg_recharge_fail"

	// TopicUserJourneyActionNotify 用户旅程动作通知
	TopicUserJourneyActionNotify = "user_journey_action_notify"
	// TopicRewardUnclaimedNotify 奖励24小时未领取通知
	TopicRewardUnclaimedNotify = "reward_unclaimed_notify"

	// TopicUserBankruptNotify 用户破产事件
	TopicUserBankruptNotify = "user_bankrupt_notify"
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
	TopicAgentBetRebatNotify = prefix + TopicAgentBetRebatNotify
	TopicUserRegisterNotify = prefix + TopicUserRegisterNotify
	TopicUserWithdrawCreateNotify = prefix + TopicUserWithdrawCreateNotify
	TopicWithdrawOrderNotify = prefix + TopicWithdrawOrderNotify
	TopicUserRiskMonitorNotify = prefix + TopicUserRiskMonitorNotify
	TopicUserVipLevelChangeNotify = prefix + TopicUserVipLevelChangeNotify
	TopicJackpotMatchRankNotify = prefix + TopicJackpotMatchRankNotify
	TopicUserPageLoginNotify = prefix + TopicUserPageLoginNotify
	TopicUserActivityCompleteNotify = prefix + TopicUserActivityCompleteNotify
	TopicWebsocketOnlineNotify = prefix + TopicWebsocketOnlineNotify
	TopicSiteMsgVipNotifyOrdered = prefix + TopicSiteMsgVipNotifyOrdered
	TopicSiteMsgActivityReward = prefix + TopicSiteMsgActivityReward
	TopicSiteMsgActivityJackpot = prefix + TopicSiteMsgActivityJackpot
	TopicSiteMsgActivitySchedule = prefix + TopicSiteMsgActivitySchedule
	TopicSiteMsgWithdrawAudit = prefix + TopicSiteMsgWithdrawAudit
	TopicSiteMsgRechargeFail = prefix + TopicSiteMsgRechargeFail
	TopicUserJourneyActionNotify = prefix + TopicUserJourneyActionNotify
	TopicRewardUnclaimedNotify = prefix + TopicRewardUnclaimedNotify

	return
}
