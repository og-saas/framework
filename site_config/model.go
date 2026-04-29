package site_config

import (
	"github.com/og-saas/framework/utils/consts"
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
	Icon     string             `json:"icon"`      // 图标
	Link     string             `json:"link"`      // 跳转链接
	Names    []*LanguageContent `json:"names"`     // 名称列表
	LinkType int                `json:"link_type"` // 链接类型
}

// CurrencyConvertConfig 币种兑换配置
type CurrencyConvertConfig struct {
	Status             consts.StatusType `json:"status"`               // 状态
	FeeType            int64             `json:"fee_type"`             // 手续费模式 1-固定 2-比例
	Multiple           int64             `json:"multiple"`             // 流水倍数
	FeeValue           int64             `json:"fee_value"`            // 手续费值
	MinExchange        int64             `json:"min_exchange"`         // 最小兑换平台币
	DailyFreeNum       int64             `json:"daily_free_num"`       // 每日免费次数
	AutoAuditLimit     int64             `json:"auto_audit_limit"`     // 审核阀值
	SingleMaxLimit     int64             `json:"single_max_limit"`     // 单笔最高限额
	DailyPlayerLimit   int64             `json:"daily_player_limit"`   // 单日玩家兑换上限(全站)
	ExchangeCoolDown   int64             `json:"exchange_cool_down"`   // 兑换冷却时间
	ExchangeOpenHours  []int64           `json:"exchange_open_hours"`  // 兑换开放时间段
	DailyExchangeLimit int64             `json:"daily_exchange_limit"` // 单日兑换上限(单人玩家)
}

type GameCalcBetAmount struct {
	ActiveIndex int                     `json:"active_index"`
	Items       []GameCalcBetAmountItem `json:"items"`
}

type GameCalcBetAmountItem struct {
	Type GameCalcBetAmountType `json:"type"`
	Desc string                `json:"desc"`
}
