package mq

// UserWalletTransferNotify 用户钱包交易通知
type UserWalletTransferNotify struct {
	UserId          int64  `json:"user_id,omitempty"`          // 用户ID
	SiteId          int64  `json:"site_id,omitempty"`          // 站点ID
	ChannelId       int64  `json:"channel_id,omitempty"`       // 渠道ID
	TradeTime       int64  `json:"trade_time,omitempty"`       // 交易时间（毫秒级时间戳）
	Username        string `json:"username,omitempty"`         // 用户账号
	OrderNo         string `json:"order_no,omitempty"`         // 订单号
	CurrencyCode    string `json:"currency_code,omitempty"`    // 币种编码
	WalletType      int32  `json:"wallet_type,omitempty"`      // 账变钱包
	Category        int32  `json:"category,omitempty"`         // 账变大类
	SubCategory     int32  `json:"sub_category,omitempty"`     // 账变小类
	Amount          string `json:"amount,omitempty"`           // 变动金额
	BalancePrevious string `json:"balance_previous,omitempty"` // 变动前余额, 单位:微
	BalanceAfter    string `json:"balance_after,omitempty"`    // 变动后余额, 单位:微
	Remark          string `json:"remark,omitempty"`           // 备注
	Operator        string `json:"operator,omitempty"`         // 操作人
	RelatedId       int64  `json:"related_id,omitempty"`       // 关联ID，如游戏ID，商品ID，活动ID
	TransactionId   int64  `json:"transaction_id,omitempty"`   // 主键 id
}

// GameBetRecordNotify 游戏下注记录通知
type GameBetRecordNotify struct {
	RecordNo            string `json:"record_no,omitempty"`              // 记录号
	ThirdSummaryId      int64  `json:"third_summary_id,omitempty"`       // 三方汇总 id
	UserId              int64  `json:"user_id,omitempty"`                // 用户 id
	Username            string `json:"username,omitempty"`               // 用户名
	GameId              int64  `json:"game_id,omitempty"`                // 游戏 id
	GameName            string `json:"game_name,omitempty"`              // 游戏名称
	ThirdGameCategoryId int64  `json:"third_game_category_id,omitempty"` // 三方游戏分类 id
	GameCategoryId      int64  `json:"game_category_id,omitempty"`       // 游戏分类 id
	GameCategoryName    string `json:"game_category_name,omitempty"`     // 游戏分类名称
	ThirdPlatformId     int64  `json:"third_platform_id,omitempty"`      // 三方厂商 id
	GamePlatformId      int64  `json:"game_platform_id,omitempty"`       // 厂商 id
	GamePlatformName    string `json:"game_platform_name,omitempty"`     // 厂商名称
	CurrencyCode        string `json:"currency_code,omitempty"`          // 币种
	RoundId             string `json:"round_id,omitempty"`               // 牌局 id
	BetCount            int64  `json:"bet_count,omitempty"`              // 投注笔数
	BetAmount           string `json:"bet_amount,omitempty"`             // 投注金额
	ValidBetAmount      string `json:"valid_bet_amount,omitempty"`       // 有效投注金额
	CancelBetCount      int64  `json:"cancel_bet_count,omitempty"`       // 取消投注次数
	CancelBetAmount     string `json:"cancel_bet_amount,omitempty"`      // 取消投注金额
	SettleCount         int64  `json:"settle_count,omitempty"`           // 结算次数
	SettleAmount        string `json:"settle_amount,omitempty"`          // 结算金额
	CancelSettleCount   int64  `json:"cancel_settle_count,omitempty"`    // 取消结算次数
	CancelSettleAmount  string `json:"cancel_settle_amount,omitempty"`   // 取消结算金额
	BetStatus           int32  `json:"bet_status,omitempty"`             // 下注状态 1=未结算，2=已结算，3=已撤单
	BetAt               int64  `json:"bet_at,omitempty"`                 // 下注时间
	SettledAt           int64  `json:"settled_at,omitempty"`             // 结算时间
	ThirdUpdateAt       int64  `json:"third_update_at,omitempty"`        // 三方更新时间
	SiteId              int64  `json:"site_id,omitempty"`                // 站点 id
}

// RechargeOrderNotify 充值订单通知
type RechargeOrderNotify struct {
	UserId                int64  `json:"user_id,omitempty"`                  // 用户 ID
	OrderNo               string `json:"order_no,omitempty"`                 // 订单编号
	ThirdOrderNo          string `json:"third_order_no,omitempty"`           // 三方订单号
	Username              string `json:"username,omitempty"`                 // 用户名
	CurrencyCode          string `json:"currency_code,omitempty"`            // 币种
	RechargeAmount        string `json:"recharge_amount,omitempty"`          // 充值金额
	FeeAmount             string `json:"fee_amount,omitempty"`               // 手续费金额
	GiftAmount            string `json:"gift_amount,omitempty"`              // 赠送金额
	ActualAmount          string `json:"actual_amount,omitempty"`            // 实际到账金额
	Phone                 string `json:"phone,omitempty"`                    // 手机号
	RealName              string `json:"real_name,omitempty"`                // 真实姓名
	OrderStatus           int32  `json:"order_status,omitempty"`             // 订单状态
	FailReason            string `json:"fail_reason,omitempty"`              // 失败原因
	SitePaymentPlatformId int64  `json:"site_payment_platform_id,omitempty"` // 站点三方支付平台 ID
	PaymentPlatformCode   string `json:"payment_platform_code,omitempty"`    // 支付平台编码
	SitePaymentChannelId  int64  `json:"site_payment_channel_id,omitempty"`  // 站点三方支付通道 ID
	PaymentChannelCode    string `json:"payment_channel_code,omitempty"`     // 通道编码
	PaymentTypeCode       string `json:"payment_type_code,omitempty"`        // 支付类型编码
	SuccessTime           int64  `json:"success_time,omitempty"`             // 订单成功时间
	FailTime              int64  `json:"fail_time,omitempty"`                // 订单失败时间
	Remark                string `json:"remark,omitempty"`                   // 备注
	SiteId                int64  `json:"site_id,omitempty"`                  // 站点 ID
}
