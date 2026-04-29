package mq

import (
	"github.com/og-saas/framework/utils/consts"
)

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
	UserId              int64  `json:"user_id,omitempty"`                // 用户 id
	Username            string `json:"username,omitempty"`               // 用户名
	GameId              int64  `json:"game_id,omitempty"`                // 游戏 id
	ThirdGameCategoryId int64  `json:"third_game_category_id,omitempty"` // 三方游戏分类 id
	GameCategoryId      int64  `json:"game_category_id,omitempty"`       // 游戏分类 id
	ThirdPlatformId     int64  `json:"third_platform_id,omitempty"`      // 三方厂商 id
	GamePlatformId      int64  `json:"game_platform_id,omitempty"`       // 厂商 id
	CurrencyCode        string `json:"currency_code,omitempty"`          // 币种
	RoundId             string `json:"round_id,omitempty"`               // 牌局 id
	BetAmount           string `json:"bet_amount,omitempty"`             // 投注金额
	CancelBetAmount     string `json:"cancel_bet_amount,omitempty"`      // 取消投注金额
	SettleAmount        string `json:"settle_amount,omitempty"`          // 结算金额
	CancelSettleAmount  string `json:"cancel_settle_amount,omitempty"`   // 取消结算金额
	BetStatus           int32  `json:"bet_status,omitempty"`             // 下注状态 1=未结算，2=已结算，3=已撤单
	BetAt               int64  `json:"bet_at,omitempty"`                 // 下注时间
	SettledAt           int64  `json:"settled_at,omitempty"`             // 结算时间
	SiteId              int64  `json:"site_id,omitempty"`                // 站点 id

	ValidBetAmount    string         `json:"valid_bet_amount,omitempty"`    // 有效投注金额
	ValidPtbAmount    consts.PtbCoin `json:"valid_ptb_amount,omitempty"`    // 有效投注平台币(有效投注金额换算出来的)
	UserWinAmount     consts.PtbCoin `json:"user_win_amount,omitempty"`     // 玩家输赢金额 赢为正 输为负
	BetPtbAmount      consts.PtbCoin `json:"bet_ptb_amount,omitempty"`      // 实际投注金额（平台币） BetAmount - CancelBetAmount
	ValidSettleAmount consts.PtbCoin `json:"valid_settle_amount,omitempty"` // 实际结算金额（平台币） SettleAmount - CancelSettleAmount

	RecordId     int64  `json:"record_id,omitempty"`     // 投注记录表主键id
	ConvertRatio string `json:"convert_ratio,omitempty"` // 转换比例
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
	FirstSign             bool   `json:"first_sign,omitempty"`               // 是否首充

	RechargePtbAmount consts.PtbCoin `json:"valid_ptb_amount,omitempty"` // 充值金额平台币(充值金额换算出来的)
	ConvertRatio      string         `json:"convert_ratio,omitempty"`    // 转换比例
}

type AgentLevelGrowthNotify struct {
	UserId                   int64          `json:"user_id"`                               // 用户 ID
	GrowthType               int32          `json:"growth_type"`                           // 成长类型 1-充值 2-投注 3-登录
	RechargePtbAmount        consts.PtbCoin `json:"recharge_ptb_amount,omitempty"`         // 充值金额平台币(充值金额换算出来的) 有效 代理总计
	BetPtbAmount             consts.PtbCoin `json:"bet_ptb_amount,omitempty"`              // 投注金额平台币(投注金额换算出来的) 有效 代理总计
	CurrentRechargePtbAmount consts.PtbCoin `json:"current_recharge_ptb_amount,omitempty"` // 充值金额平台币(充值金额换算出来的) 有效 当次
	CurrentBetPtbAmount      consts.PtbCoin `json:"current_bet_ptb_amount,omitempty"`      // 投注金额平台币(投注金额换算出来的) 有效 当次
	ContinueLoginDay         int64          `json:"continue_login_day,omitempty"`          // 连续登录天数
	LoginDay                 int64          `json:"login_day,omitempty"`                   // 累计登录天数
	SiteId                   int64          `json:"site_id"`                               // 站点 ID
	LabelNo                  string         `json:"label_no"`                              // 消息标识
	DataTime                 int64          `json:"data_time,omitempty"`                   // 数据时间
}

type AgentBetRebatNotify struct {
	SiteId         int64  `json:"site_id"`                    // 站点 ID
	UserId         int64  `json:"user_id,omitempty"`          // 用户 ID
	RecordNo       string `json:"record_no,omitempty"`        // 记录号
	RecordId       int64  `json:"record_id,omitempty"`        // 投注记录表主键id
	BetPtbAmount   string `json:"bet_ptb_amount,omitempty"`   // 投注金额（平台币）
	ValidPtbAmount string `json:"valid_ptb_amount,omitempty"` // 有效投注平台币(有效投注金额换算出来的)
	UserWinAmount  string `json:"user_win_amount,omitempty"`  // 玩家输赢金额 赢为正 输为负
	BetAt          int64  `json:"bet_at,omitempty"`           // 下注时间
	SettledAt      int64  `json:"settled_at,omitempty"`       // 结算时间
	GameCategoryId int64  `json:"game_category_id,omitempty"` // 游戏分类 id
	CurrencyCode   string `json:"currency_code,omitempty"`    // 币种
	ConvertRatio   string `json:"convert_ratio,omitempty"`    // 转换比例
}

// UserRegisterNotify 用户注册通知
type UserRegisterNotify struct {
	SiteID         int64  `json:"site_id"`         // 站点id
	ChannelID      int64  `json:"channel_id"`      // 渠道id
	UserID         int64  `json:"user_id"`         // 用户id
	IP             string `json:"ip"`              // 注册ip
	Domain         string `json:"domain"`          // 注册域名
	DeviceID       string `json:"device_id"`       // 设备id
	DeviceEndpoint string `json:"device_endpoint"` // 设备终端 APP H5 PC
	RegisterType   int32  `json:"register_type"`   // 注册方式
	SourceTarget   int32  `json:"source_target"`   // 注册来源 1-活动 2-游戏 3-代理推广
	SourceID       int32  `json:"source_id"`       // 注册来源id
	CountryCode    string `json:"country_code"`    // 国家码
	RegisterAt     int64  `json:"register_at"`     // 注册时间
}

// WithdrawOrderCreateNotify 发起提现通知
type WithdrawOrderCreateNotify struct {
	OrderId        int64  `json:"order_id,omitempty"`        // 订单ID 主键
	UserId         int64  `json:"user_id,omitempty"`         // 用户ID
	SiteId         int64  `json:"site_id,omitempty"`         // 站点ID
	OrderNo        string `json:"order_no,omitempty"`        // 订单编号
	CurrencyCode   string `json:"currency_code,omitempty"`   // 币种
	WithdrawAmount string `json:"withdraw_amount,omitempty"` // 提现金额
	ActualAmount   string `json:"actual_amount,omitempty"`   // 实际到账金额 = 提现金额 - 手续费金额
	AccountId      int64  `json:"account_id,omitempty"`      // 提现账号ID
	OrderStatus    int32  `json:"order_status,omitempty"`    // 提现订单状态 1-待审核 2-审核驳回 3-待出款/审核通过 4-出款成功 5-出款失败 6-出款已取消 7-提现失败
	ShowStatus     int32  `json:"show_status,omitempty"`     // C端展示状态 1-处理中 2-已到账 3-失败 4-已取消
}
