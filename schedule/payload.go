package schedule

import "github.com/shopspring/decimal"

type CheckGameTransferPayload struct {
	SiteId             int64           `json:"site_id"`               // 站点ID
	UserId             int64           `json:"user_id"`               // 用户ID
	CurrencyCode       string          `json:"currency_code"`         // 币种
	TransferInOrderNo  string          `json:"transfer_in_order_no"`  // 转入订单号
	TransferOutOrderNo string          `json:"transfer_out_order_no"` // 转出订单号
	WalletType         int32           `json:"wallet_type"`           // 钱包类型
	Amount             decimal.Decimal `json:"amount"`                // 金额
}
