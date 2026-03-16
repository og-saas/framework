package site_config

import "github.com/zeromicro/go-zero/core/stringx"

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
