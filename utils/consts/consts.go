package consts

import (
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/stringx"
)

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

func (p PtbCoin) FromString(val string) PtbCoin {
	if stringx.HasEmpty(val) {
		return p.FromDecimal(decimal.Zero)
	}
	valDecimal, err := decimal.NewFromString(val)
	if err != nil {
		return p.FromDecimal(decimal.Zero)
	}
	p.Decimal = valDecimal
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
	DefaultOrder                 OrderPrefix = "OR"   // 默认订单
	RechargeOrder                OrderPrefix = "RO"   // 充值订单
	WithdrawOrder                OrderPrefix = "WO"   // 提现订单
	OrderPrefixVip               OrderPrefix = "VIP"  // VIP奖励订单
	OrderPrefixAgentInviteReward OrderPrefix = "AIRO" // 代理邀请奖励订单
	OrderPrefixAgentRebate       OrderPrefix = "ARO"  // 代理返佣订单
)

// Trace 上下文 key
const Trace = "trace"

// DeviceType 设备类型
type DeviceType string

const (
	DeviceTypeWindowsPC      DeviceType = "Windows-PC"      // Windows PC
	DeviceTypeWindowsPWA     DeviceType = "Windows-PWA"     // Windows PWA
	DeviceTypeMacPC          DeviceType = "Mac-PC"          // Mac PC
	DeviceTypeMacPWA         DeviceType = "Mac-PWA"         // Mac PWA
	DeviceTypeLinuxPC        DeviceType = "Linux-PC"        // Linux PC
	DeviceTypeLinuxPWA       DeviceType = "Linux-PWA"       // Linux PWA
	DeviceTypeIOSWebApp      DeviceType = "iOS-WebApp"      // iOS WebApp（原生浏览器）
	DeviceTypeIOSPWA         DeviceType = "iOS-PWA"         // iOS PWA（添加到主屏幕）
	DeviceTypeIOSWebView     DeviceType = "iOS-WebView"     // iOS WebView（App 内嵌浏览器）
	DeviceTypeAndroidWebApp  DeviceType = "Android-WebApp"  // Android WebApp（原生浏览器）
	DeviceTypeAndroidPWA     DeviceType = "Android-PWA"     // Android PWA（添加到主屏幕）
	DeviceTypeAndroidWebView DeviceType = "Android-WebView" // Android WebView（App 内嵌浏览器）
	DeviceTypeWeb            DeviceType = "Web"             // Web 浏览器（无法识别具体设备时）
	DeviceTypeUnknown        DeviceType = "Unknown"         // 未知设备
	DeviceTypeAndroid        DeviceType = "Android"         // Android
	DeviceTypeIOS            DeviceType = "iOS"             // iOS
)

func (d DeviceType) EndpointType() EndpointType {
	switch d {
	case DeviceTypeWindowsPC, DeviceTypeMacPC, DeviceTypeLinuxPC,
		DeviceTypeWindowsPWA, DeviceTypeMacPWA, DeviceTypeLinuxPWA:
		return EndpointTypePC
	case DeviceTypeIOSWebApp, DeviceTypeIOSPWA, DeviceTypeIOSWebView, DeviceTypeAndroidWebApp,
		DeviceTypeAndroidPWA, DeviceTypeAndroidWebView, DeviceTypeWeb, DeviceTypeUnknown:
		return EndpointTypeH5
	case DeviceTypeAndroid, DeviceTypeIOS:
		return EndpointTypeApp
	}

	return EndpointTypeH5
}

// EndpointType 终端信息
type EndpointType string

const (
	EndpointTypeH5  EndpointType = "H5"  // H5
	EndpointTypeApp EndpointType = "APP" // APP
	EndpointTypePC  EndpointType = "PC"  // PC
)
