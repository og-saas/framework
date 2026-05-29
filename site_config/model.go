package site_config

import (
	"github.com/og-saas/framework/utils/consts"
	commonv1 "github.com/og-saas/proto/pb/common/v1"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/stringx"
)

type LangAware interface {
	GetLangData() (string, bool)
}

// LanguageContent 语言内容信息
type LanguageContent struct {
	Lang    string `json:"lang"`    // 语言编码
	Content string `json:"content"` // 内容
}

func (l *LanguageContent) GetLangData() (string, bool) {
	var hasContent bool
	if stringx.NotEmpty(l.Content) {
		hasContent = true
	}
	return l.Lang, hasContent
}

// VipPromoteInfo vip推广信息
type VipPromoteInfo struct {
	DefaultLanguage string                  `json:"default_language"`
	H5              *VipEndpointPromoteInfo `json:"h5"`
	PC              *VipEndpointPromoteInfo `json:"pc"`
	APP             *VipEndpointPromoteInfo `json:"app"`
}

// VipEndpointPromoteInfo vip终端推广信息
type VipEndpointPromoteInfo struct {
	Image    string             `json:"image"`    // 图片
	Contents []*LanguageContent `json:"contents"` // 内容列表
}

// VipCustomerService vip客服信息
type VipCustomerService struct {
	DefaultLanguage string                        `json:"default_language"`
	H5              []*VipEndpointCustomerService `json:"h5"`
	PC              []*VipEndpointCustomerService `json:"pc"`
	APP             []*VipEndpointCustomerService `json:"app"`
}

// VipEndpointCustomerService vip终端客服信息
type VipEndpointCustomerService struct {
	VipLevels         []int32            `json:"vip_level"`               // vip等级
	Icon              string             `json:"icon"`                    // 图标
	Link              string             `json:"link"`                    // 跳转链接
	Names             []*LanguageContent `json:"names"`                   // 名称列表
	LinkType          int                `json:"link_type"`               // 链接类型
	CustomerServiceID int64              `json:"customer_service,string"` // 客服ID
}

// CurrencyConvertConfig 币种兑换配置
type CurrencyConvertConfig struct {
	AutoConvert int `json:"auto_convert"` // 客户端自动换算开关1=开2=关
	FeeEnabled  int `json:"fee_enabled"`  // 是否启用手续费1=开2=关,启用后,需要配置fee_setting
	FeeSetting  struct {
		FeeType int     `json:"fee_type"` // 手续费类型percent或fixed
		FeeRate float64 `json:"fee_rate"` // 手续费值
	} `json:"fee_setting"` // 手续费配置
	ExchangeMin            int64   `json:"exchange_min"`              // 单笔最小兑换
	ExchangeMax            int64   `json:"exchange_max"`              // 单笔最大兑换
	ExchangeDailyLimit     int64   `json:"exchange_daily_limit"`      // 每日兑换上限
	UserDailyTotalLimit    int64   `json:"user_daily_total_limit"`    // 单日玩家总上限
	ExchangeDailyFreeCount int64   `json:"exchange_daily_free_count"` // 每日免费兑换次数,0=无免费次数
	ExchangeCooldown       int64   `json:"exchange_cooldown"`         // 兑换冷却时间(秒),0=无冷却
	OperationTimeUnlimited int     `json:"operation_time_unlimited"`  // 是否不限制运营时间,1=开 2=关
	AutoCredit             int     `json:"auto_credit"`               // 自动入账开关,1=开 2=关
	OperatingHours         []int64 `json:"operating_hours"`           // 运营时段配置
}

type GameCalcBetAmount struct {
	ActiveIndex int                     `json:"active_index"`
	Items       []GameCalcBetAmountItem `json:"items"`
}

type GameCalcBetAmountItem struct {
	Type GameCalcBetAmountType `json:"type"`
	Desc string                `json:"desc"`
}

type PayoutMonitor struct {
	Rules   []PayoutMonitorRule `json:"rules"`
	Enabled bool                `json:"enabled"`
}

type PayoutMonitorRule struct {
	Remark            string          `json:"remark"`
	CurrencyCode      string          `json:"currency_code"`
	BigWinAmount      decimal.Decimal `json:"big_win_amount"`
	HighPowerAmount   decimal.Decimal `json:"high_power_amount"`
	HighPowerMultiple decimal.Decimal `json:"high_power_multiple"`
}

// PlatformCurrencyMode 币种模式
type PlatformCurrencyMode struct {
	CurrencyCode string               `json:"currency_code"`
	Mode         PlatformCurrencyType `json:"mode"`
}

type AppIcon struct {
	Src  string `json:"src"`
	Type string `json:"type"`
}

type AppIconMap struct {
	Icon192 AppIcon `json:"icon192"`
	Icon512 AppIcon `json:"icon512"`
}

type DownGuide struct {
	Windows []*LanguageContent `json:"windows"`
	Mac     []*LanguageContent `json:"mac"`
	Ios     []*LanguageContent `json:"ios"`
	Android []*LanguageContent `json:"android"`
}

type APPInstall struct {
	DefaultLanguage string             `json:"default_language"`
	Status          consts.StatusType  `json:"status"`           // 状态 1-开启 2-关闭
	Name            []*LanguageContent `json:"name"`             // 标题
	Remark          []*LanguageContent `json:"remark"`           // 描述
	PwaName         string             `json:"pwa_name"`         // 应用名称
	PwaShortName    string             `json:"pwa_short_name"`   // 应用短名称
	StartURL        string             `json:"start_url"`        // 启动地址
	Display         string             `json:"display"`          // 显示模式
	ThemeColor      string             `json:"theme_color"`      // 主题色
	BackgroundColor string             `json:"background_color"` // 背景色
	Icons           *AppIconMap        `json:"icons"`            // PWA icon
	DownGuide       *DownGuide         `json:"down_guide"`       // 下载引导
}

// SidebarVisualMenu 侧边栏可视化配置
type SidebarVisualMenu struct {
	DefaultLanguage  string                  `json:"default_language"`
	Type             int                     `json:"type"`              // 类型：1-链接 2-组件 3-分组
	TitlesLang       []*LanguageContent      `json:"titles_lang"`       // 多语言标题
	Icon             string                  `json:"icon"`              // 图标
	LoginStatus      []int                   `json:"login_status"`      // 登录状态：1-登录前 2-登录后
	EndpointTypes    []commonv1.EndpointType `json:"endpoint_types"`    // 终端：1-h5 2-app 3-pc
	VipLevels        []int32                 `json:"vip_levels"`        // VIP等级限制
	UserIds          []int64                 `json:"user_ids"`          // 用户ID限制
	ChannelIds       []int64                 `json:"channel_ids"`       // 渠道ID限制
	Link             *SidebarLink            `json:"link"`              // 链接信息
	Component        *SidebarComponent       `json:"component"`         // 组件信息
	InteractionTypes []int                   `json:"interaction_types"` // 交互类型：1-点击 2-悬浮
	Children         []*SidebarVisualMenu    `json:"children"`          // 子集
}

type SidebarLink struct {
	LinkType commonv1.LinkType `json:"link_type"` // 链接类型
	Link     string            `json:"link"`      // 链接
	LinkArgs string            `json:"link_args"` // 链接参数
}

type SidebarComponent struct {
	ComponentType int                `json:"component_type"` // 组件类型： 1-个人中心 2-广告
	TipsLang      []*LanguageContent `json:"tips_lang"`      // 多语言提示语
	Image         string             `json:"image"`          // 图片
}
