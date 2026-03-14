package consts

import "github.com/shopspring/decimal"

// StatusType 状态
// StatusTypeEnable 启用
// StatusTypeDisable 禁用
type StatusType uint32

const (
	StatusTypeEnable  StatusType = 1 // 启用
	StatusTypeDisable StatusType = 2 // 禁用
)

// Uint32 状态转换成uint32
func (t StatusType) Uint32() uint32 {
	return uint32(t)
}

// Bool 状态转换成bool
func (t StatusType) Bool() bool {
	return t == StatusTypeEnable
}

const DefaultLanguage = "en-US"

// TransferType 转入类型
type TransferType uint32

const (
	TransferTypeIn  TransferType = 1 // 转入
	TransferTypeOut TransferType = 2 // 转出
)

type PtbCoin struct {
	decimal.Decimal
} // 平台币

func (p PtbCoin) Code() string {
	return "PTB"
}

func (p PtbCoin) ToDecimal() decimal.Decimal {
	return p.Decimal
}

func (p PtbCoin) FromDecimal(dcl decimal.Decimal) PtbCoin {
	p.Decimal = dcl
	return p
}

func NewPtbCoinFromInt(val int64) PtbCoin {
	return PtbCoin{}.FromDecimal(decimal.NewFromInt(val))
}

func NewPtbCoinFromString(val string) (PtbCoin, error) {
	dc, err := decimal.NewFromString(val)
	if err != nil {
		return PtbCoin{}, err
	}
	return PtbCoin{}.FromDecimal(dc), nil
}
func NewPtbCoinFromDecimal(val decimal.Decimal) PtbCoin {
	return PtbCoin{}.FromDecimal(val)
}

type OrderPrefix string

const (
	DefaultOrder   OrderPrefix = "OR"
	RechargeOrder  OrderPrefix = "RO"
	WithdrawOrder  OrderPrefix = "WO"
	OrderPrefixVip OrderPrefix = "VIP"
)
