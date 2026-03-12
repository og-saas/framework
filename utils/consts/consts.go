package consts

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

type PtbCoin int64 // 平台币

func (p PtbCoin) Code() string {
	return "PTB"
}

type OrderPrefix string

const (
	DefaultOrder   OrderPrefix = "OR"
	RechargeOrder  OrderPrefix = "RO"
	WithdrawOrder  OrderPrefix = "WO"
	OrderPrefixVip OrderPrefix = "VIP"
)
