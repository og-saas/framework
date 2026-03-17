package site_config

import (
	"github.com/zeromicro/go-zero/core/stringx"
)

func GetContentByLanguage(data []*LanguageContent, language, defaultLanguage string) string {
	langData := GetLanguageObject(data, language, defaultLanguage)
	if langData == nil {
		return ""
	}
	return langData.Content
}

func GetLanguageObject[T LangAware](items []T, language, defaultLanguage string) T {
	var zero T
	if items == nil || len(items) == 0 {
		return zero
	}
	var (
		defaultContent       T // 站点默认语种内容
		firstContent         T // 第一个有值内容
		hasDefault, hasFirst bool
	)
	for _, item := range items {
		lang, hasContent := item.GetLangData()
		// 优先精确匹配
		if lang == language {
			return item
		}
		// 获取站点默认语言内容
		if stringx.NotEmpty(defaultLanguage) && lang == defaultLanguage {
			defaultContent = item
			hasDefault = true
		}
		// 获取第一个有值数据
		if !hasFirst && hasContent {
			firstContent = item
			hasFirst = true
		}
	}
	if hasDefault {
		return defaultContent
	}
	if hasFirst {
		return firstContent
	}
	return zero
}
