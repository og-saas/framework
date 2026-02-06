package structx

import (
	"strings"

	"github.com/og-saas/framework/utils/consts"
)

type MultiLanguage []struct {
	Lang    string `json:"lang"`
	Content string `json:"content"`
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
