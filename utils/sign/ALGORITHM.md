# HMAC-SHA256 签名算法

## 算法规则

1. 过滤 `nil` / `null` / `undefined` 值参数
2. 将剩余参数按 **key 字典序升序** 排列
3. 拼接为 `key1=value1&key2=value2` 格式
4. 对拼接字符串做 **HMAC-SHA256** 签名
5. 输出 **hex** 编码字符串

### value 序列化规则

| 类型 | 序列化方式 | 示例 |
|------|-----------|------|
| nil / null | 跳过，不参与签名 | — |
| string | 原值，**不带引号** | `alice` |
| number | 直接转字符串 | `25` / `1.5` |
| bool | 小写字符串 | `true` / `false` |
| struct / class | JSON 序列化 | `{"name":"alice","age":25}` |
| array / List | JSON 序列化 | `["vip","active"]` |
| map / Object | **递归**按 key 字典序拼接 | `foo=bar&num=100` |

### 示例

```
输入参数:
{
  "name": "alice",
  "age": 25,
  "extra": { "foo": "bar", "num": 100 },
  "tags": ["vip", "active"]
}

拼接结果:
age=25&extra=foo=bar&num=100&name=alice&tags=["vip","active"]

签名:
HMAC-SHA256(signingKey, "age=25&extra=foo=bar&num=100&name=alice&tags=[\"vip\",\"active\"]") → hex
```

---

## Go 实现

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "sort"
    "strconv"
    "strings"
)

func HmacSha256(key, data []byte) string {
    mac := hmac.New(sha256.New, key)
    mac.Write(data)
    return hex.EncodeToString(mac.Sum(nil))
}

// SignParams 对数据签名，key 支持 string/[]byte，data 支持任意类型
func SignParams(key, data any) string {
    keyBytes := toKeyBytes(key)

    switch d := data.(type) {
    case map[string]any:
        return HmacSha256(keyBytes, []byte(buildSignString(d)))
    case nil:
        return HmacSha256(keyBytes, nil)
    default:
        b, _ := json.Marshal(d)
        return HmacSha256(keyBytes, b)
    }
}

func VerifyParams(key, data any, signature string) bool {
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
    keys := make([]string, 0, len(params))
    for k, v := range params {
        if v != nil {
            keys = append(keys, k)
        }
    }
    sort.Strings(keys)

    var buf strings.Builder
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
    default:
        b, _ := json.Marshal(val)
        return string(b)
    }
}
```

---

## JavaScript / TypeScript 实现

```typescript
import CryptoJS from "crypto-js";

function hmacSha256(key: string, data: string): string {
  return CryptoJS.HmacSHA256(data, key).toString(CryptoJS.enc.Hex);
}

function marshalValue(v: any): string {
  if (v === null || v === undefined) return "";
  switch (typeof v) {
    case "string":
      return v;
    case "number":
    case "bigint":
      return String(v);
    case "boolean":
      return String(v);
    default:
      if (typeof v === "object" && !Array.isArray(v)) {
        return buildSignString(v);
      }
      return JSON.stringify(v);
  }
}

function buildSignString(params: Record<string, any>): string {
  return Object.keys(params)
    .filter((k) => params[k] !== null && params[k] !== undefined)
    .sort()
    .map((k) => `${k}=${marshalValue(params[k])}`)
    .join("&");
}

export function signParams(key: string, params: Record<string, any>): string {
  return hmacSha256(key, buildSignString(params));
}

export function verifyParams(
  key: string,
  params: Record<string, any>,
  signature: string
): boolean {
  return signParams(key, params) === signature;
}
```

---

## Flutter / Dart 实现

```dart
import 'dart:convert';
import 'package:crypto/crypto.dart';

String hmacSha256(String key, String data) {
  final hmac = Hmac(sha256, utf8.encode(key));
  return hmac.convert(utf8.encode(data)).toString();
}

String marshalValue(dynamic v) {
  if (v == null) return '';
  if (v is String) return v;
  if (v is num) return v.toString();
  if (v is bool) return v.toString();
  if (v is Map<String, dynamic>) return buildSignString(v);
  return jsonEncode(v);
}

String buildSignString(Map<String, dynamic> params) {
  final keys = params.entries
      .where((e) => e.value != null)
      .map((e) => e.key)
      .toList();
  keys.sort();
  return keys.map((k) => '$k=${marshalValue(params[k])}').join('&');
}

String signParams(String key, Map<String, dynamic> params) {
  return hmacSha256(key, buildSignString(params));
}

bool verifyParams(String key, Map<String, dynamic> params, String signature) {
  return signParams(key, params) == signature;
}
```

> **注意**: Dart 端字符串比较非恒定时间。如需防止时序攻击，可转字节列表逐字节比较。

---

## 跨语言验证

以下测试用例用于验证三种语言实现是否一致，key = `my-secret-key`：

| 用例 | 参数 | 拼接字符串 | 签名 |
|------|------|-----------|------|
| Basic | `{name:"alice",age:25,admin:false}` | `admin=false&age=25&name=alice` | `d28fd59f20bbc56f4f35389f9a167506...` |
| Nested | `{name:"test",extra:{foo:"bar",num:100}}` | `extra=foo=bar&num=100&name=test` | `11ca33b4c085a1549f8d6787453940a7...` |
| Array | `{tags:["vip","active"],user:{name:"bob",age:30}}` | `tags=["vip","active"]&user=age=30&name=bob` | `7fb6ddcbb5da33865f3dbd9d9d87bac8...` |
| Nil | `{a:"hello",b:null,c:1}` | `a=hello&c=1` | `9b5d8c1c1c11160d7142b4073e798aa5...` |
