package cachex

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/og-saas/framework/stores/redisx"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/singleflight"
)

const (
	cacheEncodingRawJSON  byte = 0
	cacheEncodingGzipJSON byte = 1

	gzipThreshold = 1024 // 可按需调整
)

var cacheSf singleflight.Group

func gzipCompress(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(src); err != nil {
		_ = zw.Close()
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gzipDecompress(src []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	return io.ReadAll(zr)
}

func encodeCacheValue(v any) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// 小于阈值，不压缩
	if len(raw) <= gzipThreshold {
		buf := make([]byte, 1+len(raw))
		buf[0] = cacheEncodingRawJSON
		copy(buf[1:], raw)
		return buf, nil
	}

	// 大于阈值，gzip 压缩
	compressed, err := gzipCompress(raw)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1+len(compressed))
	buf[0] = cacheEncodingGzipJSON
	copy(buf[1:], compressed)
	return buf, nil
}

func decodeCacheValue[T any](bs []byte, ret *T) error {
	if len(bs) == 0 {
		return fmt.Errorf("empty cache data")
	}

	flag := bs[0]
	body := bs[1:]

	switch flag {
	case cacheEncodingRawJSON:
		return json.Unmarshal(body, ret)

	case cacheEncodingGzipJSON:
		raw, err := gzipDecompress(body)
		if err != nil {
			return err
		}
		return json.Unmarshal(raw, ret)

	default:
		return fmt.Errorf("unknown cache encoding flag: %d", flag)
	}
}

// CacheGet 从缓存中读取数据，支持自动解压
// 返回值:
//   - T: 泛型返回类型，缓存的数据
//   - error: 缓存未命中返回 redis.Nil，其他为真实错误
func CacheGet[T any, K KeyType](ctx context.Context, keyT K, args ...any) (T, error) {
	var (
		ret      T
		log      = logx.WithContext(ctx)
		redisCli = redisx.Engine.RDB(ctx)
		key      = KeyString(ctx, keyT, args...)
	)

	bs, err := redisCli.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Errorf("CacheGet error, key=%s err=%v", key, err)
		}
		return ret, err
	}

	if err = decodeCacheValue(bs, &ret); err != nil {
		log.Errorf("CacheGet decodeCacheValue error, key=%s err=%v", key, err)
		_ = redisCli.Del(ctx, key).Err()
		return ret, err
	}

	return ret, nil
}

// CacheSet 将数据写入缓存，支持自动压缩
func CacheSet[T any, K KeyType](ctx context.Context, keyT K, expire time.Duration, val T, args ...any) error {
	var (
		log      = logx.WithContext(ctx)
		redisCli = redisx.Engine.RDB(ctx)
		key      = KeyString(ctx, keyT, args...)
	)

	data, err := encodeCacheValue(val)
	if err != nil {
		log.Errorf("CacheSet encodeCacheValue error, key=%s err=%v", key, err)
		return err
	}

	if err = redisCli.Set(ctx, key, data, expire).Err(); err != nil {
		log.Errorf("CacheSet error, key=%s err=%v", key, err)
		return err
	}

	return nil
}

// CacheFn 是一个带缓存的通用函数包装器，支持泛型类型
// 该函数实现了缓存读写、防击穿（singleflight）、自动压缩等功能
// 参数:
//   - ctx: 上下文对象，用于控制超时和取消操作
//   - keyT: 键类型参数，用于生成缓存键的前缀
//   - expire: 缓存过期时间
//   - fn: 数据获取函数，当缓存未命中时执行此函数回源
//   - args: 可变参数，用于生成缓存键的具体值
//
// 返回值:
//   - T: 泛型返回类型，缓存的数据
//   - error: 错误信息，如果发生错误则返回
func CacheFn[T any, K KeyType](ctx context.Context, keyT K, expire time.Duration, fn func() (T, error), args ...any) (T, error) {
	var (
		ret              T
		log              = logx.WithContext(ctx)
		key              = KeyString(ctx, keyT, args...)
		disableCacheRead = false
	)
	cache := Engine(ctx)
	if cache != nil {
		disableCacheRead = cache.Options.DisableCacheRead
	}
	if disableCacheRead {
		return fn()
	}

	// 1. 先读缓存
	if val, err := CacheGet[T](ctx, keyT, args...); err == nil {
		return val, nil
	}

	// 2. singleflight 防击穿
	v, err, _ := cacheSf.Do(key, func() (any, error) {
		// 双检缓存
		if val, err := CacheGet[T](ctx, keyT, args...); err == nil {
			return val, nil
		}

		// 回源
		innerRet, e := fn()
		if e != nil {
			return innerRet, e
		}

		// 写缓存
		if e = CacheSet(ctx, keyT, expire, innerRet, args...); e != nil {
			log.Errorf("CacheFn CacheSet error, key=%s err=%v", key, e)
		}

		return innerRet, nil
	})

	data, ok := v.(T)
	if !ok {
		return ret, fmt.Errorf("CacheFn singleflight type assert failed, key=%s", key)
	}

	return data, err
}
