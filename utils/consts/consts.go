package consts

import (
	comV1 "github.com/og-saas/proto/pb/common/v1"
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

// ActionType 用户旅程动作类型
type JourneyActionType uint32

const (
	_                       JourneyActionType = iota
	ActionTypePopup                           // 弹窗
	ActionTypeToast                           // toast
	ActionTypePush                            // push
	ActionTypeStationLetter                   // 站内信
	ActionTypeFloatIcon                       // 悬浮图标
	ActionTypeTips                            // tips
	ActionTypeBanner                          // banner
	ActionTypeActivity                        // 活动
	ActionTypeReward                          // 奖励
)

const DefaultLanguage = "en-US"

// TransferType 转入类型
type TransferType uint32

const (
	TransferTypeIn  TransferType = 1 // 转入
	TransferTypeOut TransferType = 2 // 转出
)

// PlatformCurrencyCode 平台币
const PlatformCurrencyCode = "PTB"

type PtbCoin struct {
	decimal.Decimal
} // 平台币

func (p PtbCoin) Code() string {
	return PlatformCurrencyCode
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
	OrderPrefixCurrencyConvert   OrderPrefix = "CO"   // 币种兑换订单
	GameBet                      OrderPrefix = "GB"   // 游戏打码
	TransferIn                   OrderPrefix = "TI"   // 转入操作【中台资金转入用户钱包】
	TransferOut                  OrderPrefix = "TO"   // 转出操作【用户资金转入中台】
	OrderPrefixActivity          OrderPrefix = "AO"   // 活动奖励订单
	OrderPrefixJourney           OrderPrefix = "JY"   // 用户 Journey 奖励订单
)
const (
	// OpenTelemetry 标准字段名（推荐）
	TraceID = "trace_id"
	SpanID  = "span_id"
)

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
	DeviceTypeAndroid        DeviceType = "Android-APP"     // Android APP
	DeviceTypeIOS            DeviceType = "iOS-APP"         // iOS APP
)

// EndpointType 终端信息
type EndpointType string

const (
	EndpointTypeH5  EndpointType = "H5"  // H5
	EndpointTypeApp EndpointType = "APP" // APP
	EndpointTypePC  EndpointType = "PC"  // PC
)

// EndpointSubType 子终端信息
type EndpointSubType string

const (
	EndpointSubTypeAndroidApp EndpointSubType = "android-app" // Android APP
	EndpointSubTypeAndroidH5  EndpointSubType = "android-h5"  // Android H5
	EndpointSubTypeIOSH5      EndpointSubType = "ios-h5"      // iOS H5
	EndpointSubTypeIOSApp     EndpointSubType = "ios-app"     // iOS APP
	EndpointSubTypePC         EndpointSubType = "pc"          // PC
	EndpointSubTypeOther      EndpointSubType = "other"       // 其他
)

func (d DeviceType) EndpointType() EndpointType {
	switch d {
	case DeviceTypeWindowsPC, DeviceTypeMacPC, DeviceTypeLinuxPC, DeviceTypeWindowsPWA:
		return EndpointTypePC
	case DeviceTypeMacPWA, DeviceTypeLinuxPWA,
		DeviceTypeIOSWebApp, DeviceTypeIOSPWA, DeviceTypeIOSWebView, DeviceTypeAndroidWebApp,
		DeviceTypeAndroidPWA, DeviceTypeAndroidWebView, DeviceTypeWeb, DeviceTypeUnknown:
		return EndpointTypeH5
	case DeviceTypeAndroid, DeviceTypeIOS:
		return EndpointTypeApp
	}

	return EndpointTypeH5
}

func (d DeviceType) EndpointSubType() EndpointSubType {
	switch d {
	case DeviceTypeWindowsPC, DeviceTypeMacPC, DeviceTypeLinuxPC, DeviceTypeWindowsPWA:
		return EndpointSubTypePC
	case DeviceTypeIOSWebApp, DeviceTypeIOSPWA, DeviceTypeIOSWebView:
		return EndpointSubTypeIOSH5
	case DeviceTypeAndroidWebApp, DeviceTypeAndroidPWA, DeviceTypeAndroidWebView:
		return EndpointSubTypeAndroidH5
	case DeviceTypeAndroid:
		return EndpointSubTypeAndroidApp
	case DeviceTypeIOS:
		return EndpointSubTypeIOSApp
	}

	return EndpointSubTypeOther
}

// ThirdPartyOauthType 第三方登录类型
type ThirdPartyOauthType string

const (
	ThirdPartyOauthTypeAccount  ThirdPartyOauthType = "account"
	ThirdPartyOauthTypePhone    ThirdPartyOauthType = "phone"
	ThirdPartyOauthTypeEmail    ThirdPartyOauthType = "email"
	ThirdPartyOauthTypeGoogle   ThirdPartyOauthType = "google"   // Google
	ThirdPartyOauthTypeTelegram ThirdPartyOauthType = "telegram" // Telegram
	ThirdPartyOauthTypeFacebook ThirdPartyOauthType = "facebook" // Facebook
	ThirdPartyOauthTypeX        ThirdPartyOauthType = "x"        // X
)

func (t ThirdPartyOauthType) String() string {
	return string(t)
}

func (t ThirdPartyOauthType) ToThirdAuthType() comV1.ThirdAccountType {
	switch t {
	case ThirdPartyOauthTypeAccount:
		return comV1.ThirdAccountType(1) // 账号无对应三方类型枚举，按枚举类型转换
	case ThirdPartyOauthTypePhone:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_PHONE
	case ThirdPartyOauthTypeEmail:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_EMAIL
	case ThirdPartyOauthTypeGoogle:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_GOOGLE
	case ThirdPartyOauthTypeTelegram:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_TELEGRAM
	case ThirdPartyOauthTypeFacebook:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_FACEBOOK
	case ThirdPartyOauthTypeX:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_X
	default:
		return comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_UNSPECIFIED
	}
}

func ToThirdAuthTypeString(t comV1.ThirdAccountType) ThirdPartyOauthType {
	switch t {
	case comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_GOOGLE:
		return ThirdPartyOauthTypeGoogle
	case comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_TELEGRAM:
		return ThirdPartyOauthTypeTelegram
	case comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_FACEBOOK:
		return ThirdPartyOauthTypeFacebook
	case comV1.ThirdAccountType_THIRD_ACCOUNT_TYPE_X:
		return ThirdPartyOauthTypeX
	default:
		return "unknown"
	}
}

// RechargeTargetType 充值目标类型
type RechargeTargetType int

const (
	RechargeTargetTypeActivity RechargeTargetType = iota + 1 // 充值目标类型活动
)
