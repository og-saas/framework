package redisx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// testKeyPrefix 测试 key 前缀，跑完统一清理，避免污染本地 Redis。
const testKeyPrefix = "test:redisx:decimal:"

func TestMain(m *testing.M) {
	// 连接本地 Redis（用户本地默认无密码）
	Must(Config{
		Addrs: []string{"localhost:6379"},
	})
	code := m.Run()
	// 清理测试 key
	ctx := context.Background()
	rdb := Engine.RDB(ctx)
	if rdb != nil {
		keys, _ := rdb.Keys(ctx, testKeyPrefix+"*").Result()
		if len(keys) > 0 {
			rdb.Del(ctx, keys...)
		}
	}
	m.Run() // noop, 保留 code
	_ = code
}

func key(name string) string {
	return testKeyPrefix + name
}

func clearKey(ctx context.Context, keys ...string) {
	rdb := Engine.RDB(ctx)
	for _, k := range keys {
		rdb.Del(ctx, k)
	}
}

// ===== DecimalIncrBy 场景测试 =====

func TestDecimalIncrBy_Add(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		init string // 预置值（"" 表示不设置）
		incr string
		want string
	}{
		{"zero_plus_zero", "", "0", "0"},
		{"empty_plus_one", "", "1", "1"},
		{"int_plus_int", "100", "200", "300"},
		// 经典浮点坑：IncrByFloat 0.1+0.2=0.30000000000000004
		{"float_classic", "0.1", "0.2", "0.3"},
		{"float_precision_3dp", "0.001", "0.002", "0.003"},
		{"float_precision_misalign", "0.1", "0.02", "0.12"},
		{"int_plus_float", "100", "0.5", "100.5"},
		{"float_plus_int", "0.5", "100", "100.5"},
		{"long_decimals", "1.111111", "2.222222", "3.333333"},
		{"trailing_zeros", "1.10", "0", "1.1"},
		{"leading_zeros_in_init", "0001", "1", "2"},
		{"big_number", "9999999999999999", "1", "10000000000000000"},
		{"big_decimal", "123456789.123456789", "987654321.987654321", "1111111111.111111110"},
		{"one_plus_zero", "1", "0", "1"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := key("add_" + c.name)
			clearKey(ctx, k)
			rdb := Engine.RDB(ctx)
			if c.init != "" {
				rdb.Set(ctx, k, c.init, 0)
			}
			incr, _ := decimal.NewFromString(c.incr)
			got, err := DecimalIncrBy(ctx, k, incr, 0)
			if err != nil {
				t.Fatalf("DecimalIncrBy err: %v", err)
			}
			want, _ := decimal.NewFromString(c.want)
			if !got.Equal(want) {
				t.Errorf("init=%q + incr=%q: got %s, want %s", c.init, c.incr, got.String(), c.want)
			}
		})
	}
}

// TestDecimalIncrBy_Sub 减法场景（传负值）
func TestDecimalIncrBy_Sub(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		init string
		decr string // 正数，测试代码内取负
		want string
	}{
		{"sub_basic", "100", "30", "70"},
		{"sub_to_zero", "100", "100", "0"},
		{"sub_to_negative", "50", "100", "-50"},
		{"sub_float", "0.3", "0.1", "0.2"},
		{"sub_float_precision", "1.0", "0.1", "0.9"},
		{"sub_from_zero", "0", "50", "-50"},
		{"sub_from_negative", "-50", "30", "-80"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := key("sub_" + c.name)
			clearKey(ctx, k)
			rdb := Engine.RDB(ctx)
			if c.init != "" {
				rdb.Set(ctx, k, c.init, 0)
			}
			decr, _ := decimal.NewFromString(c.decr)
			got, err := DecimalIncrBy(ctx, k, decr.Neg(), 0)
			if err != nil {
				t.Fatalf("DecimalIncrBy err: %v", err)
			}
			want, _ := decimal.NewFromString(c.want)
			if !got.Equal(want) {
				t.Errorf("init=%q - decr=%q: got %s, want %s", c.init, c.decr, got.String(), c.want)
			}
		})
	}
}

// TestDecimalIncrBy_NegativeStart 负数起步场景
func TestDecimalIncrBy_NegativeStart(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		init string
		incr string
		want string
	}{
		{"neg_plus_pos_to_pos", "-50", "100", "50"},
		{"neg_plus_pos_still_neg", "-50", "30", "-20"},
		{"neg_plus_pos_to_zero", "-50", "50", "0"},
		{"neg_plus_neg", "-50", "-30", "-80"},
		{"neg_decimal_plus_pos", "-0.5", "1", "0.5"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := key("neg_" + c.name)
			clearKey(ctx, k)
			rdb := Engine.RDB(ctx)
			rdb.Set(ctx, k, c.init, 0)
			incr, _ := decimal.NewFromString(c.incr)
			got, err := DecimalIncrBy(ctx, k, incr, 0)
			if err != nil {
				t.Fatalf("DecimalIncrBy err: %v", err)
			}
			want, _ := decimal.NewFromString(c.want)
			if !got.Equal(want) {
				t.Errorf("init=%q + incr=%q: got %s, want %s", c.init, c.incr, got.String(), c.want)
			}
		})
	}
}

// TestDecimalIncrBy_MultiAccumulate 多次累加场景（验证累加精度）
func TestDecimalIncrBy_MultiAccumulate(t *testing.T) {
	ctx := context.Background()

	t.Run("accumulate_one_tenth_ten_times", func(t *testing.T) {
		k := key("multi_one_tenth")
		clearKey(ctx, k)
		one := decimal.NewFromFloat(0.1)
		for i := 0; i < 10; i++ {
			_, err := DecimalIncrBy(ctx, k, one, 0)
			if err != nil {
				t.Fatalf("iter %d err: %v", i, err)
			}
		}
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		want := decimal.NewFromInt(1)
		if !got.Equal(want) {
			t.Errorf("0.1 × 10: got %s, want 1", got.String())
		}
	})

	t.Run("accumulate_mixed", func(t *testing.T) {
		k := key("multi_mixed")
		clearKey(ctx, k)
		// 0.1 + 0.2 + 0.3 + 0.4 = 1.0
		vals := []string{"0.1", "0.2", "0.3", "0.4"}
		for _, v := range vals {
			d, _ := decimal.NewFromString(v)
			_, err := DecimalIncrBy(ctx, k, d, 0)
			if err != nil {
				t.Fatalf("incr %s err: %v", v, err)
			}
		}
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		want := decimal.NewFromInt(1)
		if !got.Equal(want) {
			t.Errorf("0.1+0.2+0.3+0.4: got %s, want 1", got.String())
		}
	})

	t.Run("accumulate_with_sub", func(t *testing.T) {
		k := key("multi_with_sub")
		clearKey(ctx, k)
		// +100 -30 +50 -20 = 100
		ops := []string{"100", "-30", "50", "-20"}
		for _, v := range ops {
			d, _ := decimal.NewFromString(v)
			_, err := DecimalIncrBy(ctx, k, d, 0)
			if err != nil {
				t.Fatalf("incr %s err: %v", v, err)
			}
		}
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		want := decimal.NewFromInt(100)
		if !got.Equal(want) {
			t.Errorf("100-30+50-20: got %s, want 100", got.String())
		}
	})
}

// TestDecimalIncrBy_Expire 验证过期时间设置
func TestDecimalIncrBy_Expire(t *testing.T) {
	ctx := context.Background()
	rdb := Engine.RDB(ctx)

	t.Run("expire_set_seconds", func(t *testing.T) {
		k := key("expire_seconds")
		clearKey(ctx, k)
		_, err := DecimalIncrBy(ctx, k, decimal.NewFromInt(1), 100)
		if err != nil {
			t.Fatal(err)
		}
		ttl, err := rdb.TTL(ctx, k).Result()
		if err != nil {
			t.Fatal(err)
		}
		if ttl <= 0 || ttl > 100*time.Second {
			t.Errorf("TTL not in (0, 100s]: got %v", ttl)
		}
	})

	t.Run("expire_zero_means_no_ttl", func(t *testing.T) {
		k := key("expire_zero")
		clearKey(ctx, k)
		_, err := DecimalIncrBy(ctx, k, decimal.NewFromInt(1), 0)
		if err != nil {
			t.Fatal(err)
		}
		ttl, err := rdb.TTL(ctx, k).Result()
		if err != nil {
			t.Fatal(err)
		}
		// TTL = -1 表示 key 存在但无过期时间
		if ttl != -1 {
			t.Errorf("expected no TTL (-1), got %v", ttl)
		}
	})

	t.Run("expire_refresh_on_each_call", func(t *testing.T) {
		k := key("expire_refresh")
		clearKey(ctx, k)
		// 第一次设 100s
		_, _ = DecimalIncrBy(ctx, k, decimal.NewFromInt(1), 100)
		time.Sleep(2 * time.Second)
		ttl1, _ := rdb.TTL(ctx, k).Result()
		// 第二次重置 TTL 到 100s
		_, _ = DecimalIncrBy(ctx, k, decimal.NewFromInt(1), 100)
		ttl2, _ := rdb.TTL(ctx, k).Result()
		if ttl2 <= ttl1 {
			t.Errorf("TTL should refresh: before=%v, after=%v", ttl1, ttl2)
		}
	})
}

// TestDecimalIncrByAt 绝对过期时间
func TestDecimalIncrByAt(t *testing.T) {
	ctx := context.Background()
	rdb := Engine.RDB(ctx)

	t.Run("future_expire", func(t *testing.T) {
		k := key("at_future")
		clearKey(ctx, k)
		// 60s 后过期
		expireAt := time.Now().Add(60 * time.Second).Unix()
		_, err := DecimalIncrByAt(ctx, k, decimal.NewFromInt(1), expireAt)
		if err != nil {
			t.Fatal(err)
		}
		ttl, err := rdb.TTL(ctx, k).Result()
		if err != nil {
			t.Fatal(err)
		}
		if ttl <= 0 || ttl > 60*time.Second {
			t.Errorf("TTL not in (0, 60s]: got %v", ttl)
		}
	})

	t.Run("past_expire_immediate", func(t *testing.T) {
		k := key("at_past")
		clearKey(ctx, k)
		// 过期时间已过 → expireSec 被 clamp 到 0，key 永久存在
		expireAt := time.Now().Add(-10 * time.Second).Unix()
		_, err := DecimalIncrByAt(ctx, k, decimal.NewFromInt(1), expireAt)
		if err != nil {
			t.Fatal(err)
		}
		ttl, _ := rdb.TTL(ctx, k).Result()
		if ttl != -1 {
			t.Errorf("past expireAt should clamp to no-TTL (-1), got %v", ttl)
		}
	})
}

// TestDecimalGet 读取场景
func TestDecimalGet(t *testing.T) {
	ctx := context.Background()
	rdb := Engine.RDB(ctx)

	t.Run("nonexistent_key_returns_zero", func(t *testing.T) {
		k := key("get_nonexistent")
		clearKey(ctx, k)
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		if !got.IsZero() {
			t.Errorf("nonexistent key: got %s, want 0", got.String())
		}
	})

	t.Run("existing_key", func(t *testing.T) {
		k := key("get_existing")
		clearKey(ctx, k)
		rdb.Set(ctx, k, "123.456", 0)
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		want, _ := decimal.NewFromString("123.456")
		if !got.Equal(want) {
			t.Errorf("got %s, want 123.456", got.String())
		}
	})

	t.Run("existing_negative", func(t *testing.T) {
		k := key("get_negative")
		clearKey(ctx, k)
		rdb.Set(ctx, k, "-99.9", 0)
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		want, _ := decimal.NewFromString("-99.9")
		if !got.Equal(want) {
			t.Errorf("got %s, want -99.9", got.String())
		}
	})

	t.Run("empty_string_returns_zero", func(t *testing.T) {
		k := key("get_empty")
		clearKey(ctx, k)
		rdb.Set(ctx, k, "", 0)
		got, err := DecimalGet(ctx, k)
		if err != nil {
			t.Fatal(err)
		}
		if !got.IsZero() {
			t.Errorf("empty value: got %s, want 0", got.String())
		}
	})
}

// TestDecimalIncrBy_LargeNumber 大数累加
func TestDecimalIncrBy_LargeNumber(t *testing.T) {
	ctx := context.Background()

	t.Run("beyond_int64", func(t *testing.T) {
		k := key("large_beyond_int64")
		clearKey(ctx, k)
		// 9.2e18 接近 int64 上限，累加后超出但 decimal 可表示
		big, _ := decimal.NewFromString("9223372036854775807") // int64 max
		one := decimal.NewFromInt(1)
		_, err := DecimalIncrBy(ctx, k, big, 0)
		if err != nil {
			t.Fatalf("first incr err: %v", err)
		}
		got, err := DecimalIncrBy(ctx, k, one, 0)
		if err != nil {
			t.Fatalf("second incr err: %v", err)
		}
		want, _ := decimal.NewFromString("9223372036854775808")
		if !got.Equal(want) {
			t.Errorf("int64max + 1: got %s, want %s", got.String(), want.String())
		}
	})

	t.Run("high_precision", func(t *testing.T) {
		k := key("large_precision")
		clearKey(ctx, k)
		// 18 位小数精度
		a, _ := decimal.NewFromString("0.123456789012345678")
		b, _ := decimal.NewFromString("0.876543210987654322")
		_, err := DecimalIncrBy(ctx, k, a, 0)
		if err != nil {
			t.Fatal(err)
		}
		got, err := DecimalIncrBy(ctx, k, b, 0)
		if err != nil {
			t.Fatal(err)
		}
		// 0.123456789012345678 + 0.876543210987654322 = 1.000000000000000000 = 1
		want := decimal.NewFromInt(1)
		if !got.Equal(want) {
			t.Errorf("got %s, want 1", got.String())
		}
	})
}

// TestDecimalIncrBy_Concurrent 并发安全（同一 key 并发累加，结果应正确）
func TestDecimalIncrBy_Concurrent(t *testing.T) {
	ctx := context.Background()
	k := key("concurrent")
	clearKey(ctx, k)

	const goroutines = 50
	const perGoroutine = 20
	done := make(chan error, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			one := decimal.NewFromInt(1)
			for j := 0; j < perGoroutine; j++ {
				_, err := DecimalIncrBy(ctx, k, one, 0)
				if err != nil {
					done <- err
					return
				}
			}
			done <- nil
		}()
	}

	for i := 0; i < goroutines; i++ {
		if err := <-done; err != nil {
			t.Fatalf("goroutine err: %v", err)
		}
	}

	got, err := DecimalGet(ctx, k)
	if err != nil {
		t.Fatal(err)
	}
	want := decimal.NewFromInt(int64(goroutines * perGoroutine))
	if !got.Equal(want) {
		t.Errorf("concurrent %d×%d: got %s, want %s", goroutines, perGoroutine, got.String(), want.String())
	}
}

// TestDecimalIncrBy_BackwardCompatible 旧 IncrByFloat 写入的 float key 兼容性
func TestDecimalIncrBy_BackwardCompatible(t *testing.T) {
	ctx := context.Background()
	rdb := Engine.RDB(ctx)

	t.Run("legacy_float_key", func(t *testing.T) {
		k := key("legacy_float")
		clearKey(ctx, k)
		// 模拟旧 IncrByFloat 写入的 float 字符串
		rdb.Set(ctx, k, "123.45", 0)
		got, err := DecimalIncrBy(ctx, k, decimal.NewFromFloat(0.05), 0)
		if err != nil {
			t.Fatalf("err on legacy float key: %v", err)
		}
		want, _ := decimal.NewFromString("123.5")
		if !got.Equal(want) {
			t.Errorf("legacy float 123.45 + 0.05: got %s, want 123.5", got.String())
		}
	})
}

func ExampleDecimalIncrBy() {
	ctx := context.Background()
	k := fmt.Sprintf("%sexample", testKeyPrefix)
	clearKey(ctx, k)

	// 累加 0.1 十次
	one := decimal.NewFromFloat(0.1)
	for i := 0; i < 10; i++ {
		DecimalIncrBy(ctx, k, one, 0)
	}
	got, _ := DecimalGet(ctx, k)
	fmt.Println(got.String())
	// Output: 1
}
