package schedule

import (
	"github.com/og-saas/framework/utils/consts"
	"github.com/og-saas/proto/rpc/promotionservice"
	"github.com/shopspring/decimal"
)

type CheckGameTransferPayload struct {
	RecordId           int64               `json:"record_id"`             // 转账记录ID
	SiteId             int64               `json:"site_id"`               // 站点ID
	UserId             int64               `json:"user_id"`               // 用户ID
	CurrencyCode       string              `json:"currency_code"`         // 币种
	TransferInOrderNo  string              `json:"transfer_in_order_no"`  // 转入订单号
	TransferOutOrderNo string              `json:"transfer_out_order_no"` // 转出订单号
	WalletType         int32               `json:"wallet_type"`           // 钱包类型
	Amount             decimal.Decimal     `json:"amount"`                // 金额
	TransferType       consts.TransferType `json:"transfer_type"`         // 转账标记 1-转入 2-转出
}

type AgentCommissionStatRetryPayload struct {
	SiteId      int64   `json:"site_id"`
	UserIds     []int64 `json:"user_ids"`     // 用户id列表
	StartTime   int64   `json:"start_time"`   // 开始时间
	EndTime     int64   `json:"end_time"`     // 结束时间
	SettleCycle int32   `json:"settle_cycle"` // 结算周期 1-日 2-周 3-月
}

type RewardReleaseRetryPayload struct {
	SiteId  int64                               `json:"site_id"`
	Rewards []*promotionservice.CreateRewardReq `json:"rewards"`
}
