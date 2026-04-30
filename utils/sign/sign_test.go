package sign

import (
	"testing"
)

func TestHmacSha256(t *testing.T) {
	key := []byte("secret-key")
	data := []byte("data to sign")

	sig := HmacSha256(key, data)
	if sig == "" {
		t.Fatal("empty signature")
	}

	if !HmacSha256Verify(key, data, sig) {
		t.Fatal("signature verification failed")
	}

	if HmacSha256Verify(key, data, "invalid") {
		t.Fatal("should not verify invalid signature")
	}

	if HmacSha256Verify([]byte("wrong-key"), data, sig) {
		t.Fatal("should not verify with wrong key")
	}
}

func TestSignParams(t *testing.T) {
	key := []byte("signing-key")
	params := map[string]any{
		"name":  "张三",
		"age":   25,
		"city":  "北京",
		"token": "abc123",
	}

	sig := SignParams(key, params)
	if !VerifyParams(key, params, sig) {
		t.Fatal("verification failed")
	}

	// 修改参数后验证失败
	modified := map[string]any{
		"name":  "张三",
		"age":   26,
		"city":  "北京",
		"token": "abc123",
	}
	if VerifyParams(key, modified, sig) {
		t.Fatal("should not verify modified params")
	}
}

func TestSignParamsWithNil(t *testing.T) {
	key := []byte("signing-key")
	params := map[string]any{
		"name":  "test",
		"extra": nil,
	}

	sig := SignParams(key, params)
	if !VerifyParams(key, params, sig) {
		t.Fatal("nil value should be skipped")
	}

	// nil 值不影响签名
	paramsWithNil := map[string]any{
		"name":  "test",
		"extra": nil,
		"flag":  true,
	}
	sig2 := SignParams(key, paramsWithNil)

	paramsWithoutNil := map[string]any{
		"name": "test",
		"flag": true,
	}
	sig3 := SignParams(key, paramsWithoutNil)

	if sig2 != sig3 {
		t.Fatal("nil values should be skipped, signatures should match")
	}
}

func TestSignParamsWithNestedMap(t *testing.T) {
	key := []byte("signing-key")
	params := map[string]any{
		"name": "test",
		"extra": map[string]any{
			"foo": "bar",
			"num": 100,
		},
	}

	sig := SignParams(key, params)
	if !VerifyParams(key, params, sig) {
		t.Fatal("nested map verification failed")
	}

	// key 顺序不同但值相同，签名一致
	sameParams := map[string]any{
		"name": "test",
		"extra": map[string]any{
			"num": 100,
			"foo": "bar",
		},
	}
	if !VerifyParams(key, sameParams, sig) {
		t.Fatal("different key order should produce same signature")
	}
}

func TestSignParamsWithStruct(t *testing.T) {
	key := []byte("signing-key")

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	params := map[string]any{
		"user": User{Name: "张三", Age: 25},
	}

	sig := SignParams(key, params)
	if !VerifyParams(key, params, sig) {
		t.Fatal("struct verification failed")
	}

	modified := map[string]any{
		"user": User{Name: "张三", Age: 26},
	}
	if VerifyParams(key, modified, sig) {
		t.Fatal("should not verify modified struct")
	}
}

func TestSignParamsWithSlice(t *testing.T) {
	key := []byte("signing-key")
	params := map[string]any{
		"items": []string{"a", "b", "c"},
	}

	sig := SignParams(key, params)
	if !VerifyParams(key, params, sig) {
		t.Fatal("slice verification failed")
	}

	modified := map[string]any{
		"items": []string{"a", "b", "d"},
	}
	if VerifyParams(key, modified, sig) {
		t.Fatal("should not verify modified slice")
	}
}

func TestSignParamsStable(t *testing.T) {
	key := []byte("signing-key")

	sig1 := SignParams(key, map[string]any{"b": 2, "a": 1, "c": 3})
	sig2 := SignParams(key, map[string]any{"c": 3, "a": 1, "b": 2})

	if sig1 != sig2 {
		t.Fatal("same params with different iteration order should produce same signature")
	}
}
