package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// PrettyJSON 美化打印
func PrettyJSON(v interface{}) {
	// 使用 json.MarshalIndent 进行格式化和美化打印
	prettyJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	// 打印格式化后的 JSON 字符串
	fmt.Println(string(prettyJSON))
}

func ToPrettyJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}

// Ternary 三元运算符的模拟函数
func Ternary[T any](condition bool, value1, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

// Base62Encode id转base62
func Base62Encode(id int64) string {
	if id == 0 {
		return "0"
	}
	base := int64(len(base62Chars))
	result := ""

	for id > 0 {
		remainder := id % base
		result = string(base62Chars[remainder]) + result
		id /= base
	}
	return result
}

// Base62Decode base62转id
func Base62Decode(code string) int64 {
	if code == "" {
		return 0
	}

	base := int64(len(base62Chars))
	var id int64

	// 创建字符到索引的映射
	charIndex := make(map[rune]int64)
	for i, ch := range base62Chars {
		charIndex[ch] = int64(i)
	}

	// 从左到右遍历编码字符串
	for _, ch := range code {
		index, ok := charIndex[ch]
		if !ok {
			return 0
		}
		id = id*base + index
	}

	return id
}
