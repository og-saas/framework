package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// HmacSha256 对数据进行 HMAC-SHA256 签名，返回 hex 编码字符串
func HmacSha256(key, data []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha256Verify 验证 HMAC-SHA256 签名，使用恒定时间比较防止时序攻击
func HmacSha256Verify(key, data []byte, signature string) bool {
	expected := HmacSha256(key, data)
	return hmac.Equal([]byte(expected), []byte(signature))
}

// SignParams 对数据签名
//   - key: 支持 string / []byte
//   - data: map[string]any 按 key 排序拼接; 其他类型 json.Marshal 后签名
func SignParams(key any, data any) string {
	keyBytes := toKeyBytes(key)

	switch d := data.(type) {
	case map[string]any:
		return HmacSha256(keyBytes, []byte(buildSignString(d)))
	case nil:
		return HmacSha256(keyBytes, nil)
	default:
		b, err := json.Marshal(d)
		if err != nil {
			return ""
		}
		return HmacSha256(keyBytes, b)
	}
}

// VerifyParams 验证参数签名
func VerifyParams(key any, data any, signature string) bool {
	return hmac.Equal([]byte(SignParams(key, data)), []byte(signature))
}

func toKeyBytes(key any) []byte {
	switch k := key.(type) {
	case string:
		return []byte(k)
	case []byte:
		return k
	default:
		return []byte(fmt.Sprintf("%v", k))
	}
}

func buildSignString(params map[string]any) string {
	keys := sortedKeys(params)

	var buf strings.Builder
	buf.Grow(len(keys) * 32)

	for i, k := range keys {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(marshalValue(params[k]))
	}

	return buf.String()
}

func sortedKeys(params map[string]any) []string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != nil {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

func marshalValue(v any) string {
	switch val := v.(type) {
	case nil:
		return ""
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case map[string]any:
		return buildSignString(val)
	case []byte:
		return string(val)
	case json.Marshaler:
		b, _ := val.MarshalJSON()
		return string(b)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}