package cachex

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/og-saas/framework/stores/redisx"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/singleflight"
)

// Fetch2
// Key 只支持 CacheKey | TenantCacheKey | string 类型
// key为string 不会带入参数
// 如果是 TenantCacheKey, 不用传递tenant参数
func Fetch2[T KeyType](ctx context.Context, key T, expire time.Duration, fn func() (string, error), args ...any) (string, error) {
	return Engine(ctx).Fetch2(ctx, KeyString(ctx, key, args...), expire, fn)
}

// TagAsDeleted2 Key 只支持 CacheKey | TenantCacheKey | string 类型, string不会带入参数
func TagAsDeleted2[T KeyType](ctx context.Context, key T, args ...any) error {
	return Engine(ctx).TagAsDeleted2(ctx, KeyString(ctx, key, args...))
}

var cacheSf singleflight.Group

// CacheFn 通用缓存读取方法，支持泛型类型
// 实现了缓存读取、防击穿、回源、写缓存的完整流程
// Key 支持 CacheKey | TenantCacheKey | string 类型
// 使用 gob 编码存储减少redis内存占用，使用 singleflight 防止缓存击穿
func CacheFn[T any, K KeyType](ctx context.Context, keyT K, expire time.Duration, fn func() (T, error), args ...any) (T, error) {
	var (
		ret      T
		log      = logx.WithContext(ctx)
		redisCli = redisx.Engine.RDB(ctx)
		key      = KeyString(ctx, keyT, args...)
	)
	// 1. 先读缓存
	bs, err := redisCli.Get(ctx, key).Bytes()
	if err == nil {
		if err = gob.NewDecoder(bytes.NewReader(bs)).Decode(&ret); err == nil {
			return ret, nil
		}
		log.Errorf("CacheFn NewDecoder error, key=%s err=%v", key, err)
		_ = redisCli.Del(ctx, key).Err()
	} else if !errors.Is(err, redis.Nil) {
		log.Errorf("CacheFn Get error, key=%s err=%v", key, err)
	}

	// 2. 防击穿
	v, err, _ := cacheSf.Do(key, func() (any, error) {
		var innerRet T

		// 双检缓存
		bs, e := redisCli.Get(ctx, key).Bytes()
		if e == nil {
			if e = gob.NewDecoder(bytes.NewReader(bs)).Decode(&innerRet); e == nil {
				return innerRet, nil
			}
			log.Errorf("CacheFn singleflight NewDecoder error, key=%s err=%v", key, e)
			_ = redisCli.Del(ctx, key).Err()
		} else if !errors.Is(e, redis.Nil) {
			log.Errorf("CacheFn singleflight Get error, key=%s err=%v", key, e)
		}

		// 回源
		innerRet, e = fn()
		if e != nil {
			return innerRet, e
		}

		// 写缓存
		var buf bytes.Buffer
		if e = gob.NewEncoder(&buf).Encode(innerRet); e == nil {
			if e = redisCli.Set(ctx, key, buf.Bytes(), expire).Err(); e != nil {
				log.Errorf("CacheFn Set error, key=%s err=%v", key, e)
			}
		} else {
			log.Errorf("CacheFn NewEncoder error, key=%s err=%v", key, e)
		}

		return innerRet, nil
	})
	if err != nil {
		return ret, err
	}

	data, ok := v.(T)
	if !ok {
		return ret, fmt.Errorf("CacheFn singleflight type assert failed, key=%s", key)
	}

	return data, nil
}
