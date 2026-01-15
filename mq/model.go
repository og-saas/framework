package mq

// UserRechargeNotify 用户成功充值推送
type UserRechargeNotify struct {
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
	TransactionId   int64  `json:"transaction_id,omitempty"`   // 主键id
}
