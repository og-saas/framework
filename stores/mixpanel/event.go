package mixpanel

// EventType Mixpanel 事件类型。
type EventType string

const (
	EventOTPVerifyResult      EventType = "otp_verify_result"      // 验证码校验结果
	EventRegisterResult       EventType = "register_result"        // 注册结果
	EventLoginResult          EventType = "login_result"           // 登录结果
	EventDepositOrderResult   EventType = "deposit_order_result"   // 充值订单创建结果
	EventWithdrawOrderResult  EventType = "withdraw_order_result"  // 提现申请结果
	EventPromotionJoinResult  EventType = "promotion_join_result"  // 活动参与结果
	EventRewardClaimResult    EventType = "reward_claim_result"    // 奖励领取结果
	EventTournamentJoinResult EventType = "tournament_join_result" // 锦标赛报名结果
	EventBetResult            EventType = "bet_result"             // 体育投注结果
)
