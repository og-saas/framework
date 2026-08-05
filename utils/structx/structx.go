package structx

import (
	"strings"

	"github.com/og-saas/framework/utils"
	"github.com/og-saas/framework/utils/consts"
)

type MultiLanguage []struct {
	Lang    string `json:"lang"`
	Content string `json:"content"`
}

type MultiLanguageLink []struct {
	Lang     string `json:"lang"`
	Link     string `json:"link"`
	LinkArgs string `json:"link_args"`
	LinkType int    `json:"link_type"`
}

func (ml MultiLanguage) Get(lang string) string {
	var defaultContent string
	for _, item := range ml {
		if item.Lang == lang {
			return item.Content
		}
		if strings.ToLower(item.Lang) == strings.ToLower(consts.DefaultLanguage) {
			defaultContent = item.Content
		}
	}
	return defaultContent
}

func (ml MultiLanguageLink) GetLink(lang string) string {
	var defaultLink string
	for _, item := range ml {
		if item.Lang == lang {
			return utils.GetLinkUrl(item.Link, item.LinkArgs)
		}
		if strings.ToLower(item.Lang) == strings.ToLower(consts.DefaultLanguage) {
			defaultLink = utils.GetLinkUrl(item.Link, item.LinkArgs)
		}
	}
	return defaultLink
}
