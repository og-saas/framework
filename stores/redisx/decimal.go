package redisx

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

// decimalIncrScript 在 Redis 内用字符串十进制加法做原子累加 + 设置过期。
//
// 设计目的：彻底替代 IncrByFloat，规避两个问题：
//   - IEEE 754 浮点精度丢失（审计 H4）
//   - IncrByFloat + Expire 非原子（审计 H5）
//
// 实现要点：
//   - 全程不使用 Lua 数字运算（Lua 5.1 数字也是 double），改用字节码差值 +
//     字符拼接做"逐位十进制加/减法"，零精度损失。
//   - 支持带符号输入（incr 可正可负），与 IncrByFloat 语义完全对齐，
//     finance-rpc 的 Decr 场景传负值即可。
//   - SET 与 EXPIRE 在同一 Lua 调用内执行，避免非原子窗口。
//
// 入参：
//   - KEYS[1] = 累加 key
//   - ARGV[1] = 增量值（decimal.Decimal.String()，可带前导 -）
//   - ARGV[2] = 过期秒数，<=0 表示不设置过期
//
// 返回：累加后的十进制字符串
var decimalIncrScript = redis.NewScript(`
local function dec_split(s)
	-- 拆分符号 / 整数部分 / 小数部分，返回 (is_neg, int_part, frac_part)
	local neg = false
	if s:sub(1, 1) == '-' then
		neg = true
		s = s:sub(2)
	elseif s:sub(1, 1) == '+' then
		s = s:sub(2)
	end
	if s == '' then s = '0' end

	local intp, fracp = s, ''
	local dot = s:find('%.')
	if dot then
		intp = s:sub(1, dot - 1)
		fracp = s:sub(dot + 1)
	end
	if intp == '' then intp = '0' end

	-- 去除前导零，保留至少一位
	intp = intp:gsub('^0+', '')
	if intp == '' then intp = '0' end

	return neg, intp, fracp
end

local function dec_strip_leading_zeros(s)
	s = s:gsub('^0+', '')
	if s == '' then s = '0' end
	return s
end

-- 绝对值比较：a vs b（均为非负十进制字符串，允许小数点）
-- 返回: 1 (a>b), 0 (a==b), -1 (a<b)
local function dec_abs_cmp(a, b)
	local _, ai, af = dec_split(a)
	local _, bi, bf = dec_split(b)

	if #ai > #bi then return 1 end
	if #ai < #bi then return -1 end
	if ai > bi then return 1 end
	if ai < bi then return -1 end

	-- 整数部分相同，比小数部分（补齐等长后字典序比较）
	local nf = #af
	if #bf > nf then nf = #bf end
	af = af .. string.rep('0', nf - #af)
	bf = bf .. string.rep('0', nf - #bf)
	if af > bf then return 1 end
	if af < bf then return -1 end
	return 0
end

-- 绝对值加法：a + b（均为非负十进制字符串）
local function dec_abs_add(a, b)
	local _, ai, af = dec_split(a)
	local _, bi, bf = dec_split(b)

	local nf = #af
	if #bf > nf then nf = #bf end
	af = af .. string.rep('0', nf - #af)
	bf = bf .. string.rep('0', nf - #bf)

	local ni = #ai
	if #bi > ni then ni = #bi end
	ai = string.rep('0', ni - #ai) .. ai
	bi = string.rep('0', ni - #bi) .. bi

	local sa = ai .. af
	local sb = bi .. bf
	local n = #sa

	local out = {}
	local carry = 0
	for i = n, 1, -1 do
		local sum = (sa:byte(i) - 48) + (sb:byte(i) - 48) + carry
		if sum >= 10 then
			carry = 1
			sum = sum - 10
		else
			carry = 0
		end
		out[i] = string.char(sum + 48)
	end

	if carry > 0 then
		table.insert(out, 1, '1')
		ni = ni + 1
	end

	local all = table.concat(out)
	if nf == 0 then
		return all
	end

	local intPart = all:sub(1, ni)
	local fracPart = all:sub(ni + 1):gsub('0+$', '')
	if fracPart == '' then
		return dec_strip_leading_zeros(intPart)
	end
	return dec_strip_leading_zeros(intPart) .. '.' .. fracPart
end

-- 绝对值减法：a - b（前提 a >= b，均为非负十进制字符串）
local function dec_abs_sub(a, b)
	local _, ai, af = dec_split(a)
	local _, bi, bf = dec_split(b)

	local nf = #af
	if #bf > nf then nf = #bf end
	af = af .. string.rep('0', nf - #af)
	bf = bf .. string.rep('0', nf - #bf)

	local ni = #ai
	if #bi > ni then ni = #bi end
	ai = string.rep('0', ni - #ai) .. ai
	bi = string.rep('0', ni - #bi) .. bi

	local sa = ai .. af
	local sb = bi .. bf
	local n = #sa

	local out = {}
	local borrow = 0
	for i = n, 1, -1 do
		local diff = (sa:byte(i) - 48) - (sb:byte(i) - 48) - borrow
		if diff < 0 then
			diff = diff + 10
			borrow = 1
		else
			borrow = 0
		end
		out[i] = string.char(diff + 48)
	end

	local all = table.concat(out)
	if nf == 0 then
		return dec_strip_leading_zeros(all)
	end

	local intPart = all:sub(1, ni)
	local fracPart = all:sub(ni + 1):gsub('0+$', '')
	intPart = dec_strip_leading_zeros(intPart)
	if fracPart == '' then
		return intPart
	end
	return intPart .. '.' .. fracPart
end

-- 带符号加法：支持 incr 为正/负
local function dec_add(a, b)
	local a_neg, _, _ = dec_split(a)
	local b_neg, _, _ = dec_split(b)
	-- 去符号后传给 abs 函数
	local a_abs = a
	if a_neg then a_abs = a:sub(2) end
	local b_abs = b
	if b_neg then b_abs = b:sub(2) end
	-- 再去一次前导零（dec_split 已做，但保险）
	a_abs = dec_strip_leading_zeros(a_abs:gsub('^%+', ''))
	b_abs = dec_strip_leading_zeros(b_abs:gsub('^%+', ''))

	if a_neg == b_neg then
		local sum = dec_abs_add(a_abs, b_abs)
		if a_neg and sum ~= '0' then
			return '-' .. sum
		end
		return sum
	end

	-- 异号：绝对值大减小
	local cmp = dec_abs_cmp(a_abs, b_abs)
	if cmp == 0 then return '0' end
	if cmp > 0 then
		local diff = dec_abs_sub(a_abs, b_abs)
		if a_neg then return '-' .. diff end
		return diff
	end
	local diff = dec_abs_sub(b_abs, a_abs)
	if b_neg then return '-' .. diff end
	return diff
end

local cur = redis.call('GET', KEYS[1])
if not cur then cur = '0' end
local newval = dec_add(cur, ARGV[1])

local ttl = tonumber(ARGV[2])
if ttl and ttl > 0 then
	redis.call('SET', KEYS[1], newval, 'EX', ttl)
else
	redis.call('SET', KEYS[1], newval)
end
return newval
`)

// DecimalIncrBy 在 Redis 上对 key 做十进制原子累加，返回累加后的值。
// incr 可正可负（负数即减法），expireSec <= 0 表示不设置过期时间。
//
// 替代 IncrByFloat + Expire 两步操作，同时解决：
//   - 浮点精度丢失（改用字符串逐位加/减法）
//   - 累加与过期非原子（Lua 单次调用原子完成）
func DecimalIncrBy(ctx context.Context, key string, incr decimal.Decimal, expireSec int64) (decimal.Decimal, error) {
	if key == "" {
		return decimal.Zero, errors.New("redisx.DecimalIncrBy: key is empty")
	}
	if Engine == nil {
		return decimal.Zero, errors.New("redisx.DecimalIncrBy: Engine not initialized")
	}
	rdb := Engine.RDB(ctx)
	if rdb == nil {
		return decimal.Zero, errors.New("redisx.DecimalIncrBy: no redis client for tenant")
	}
	res, err := decimalIncrScript.Run(ctx, rdb,
		[]string{key},
		incr.String(), expireSec,
	).Text()
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromString(res)
}

// DecimalIncrByAt 与 DecimalIncrBy 等价，但接受绝对过期时间戳（unix 秒）。
// 内部转换为剩余秒数，兼容原 ExpireAt 语义。
func DecimalIncrByAt(ctx context.Context, key string, incr decimal.Decimal, expireAtUnix int64) (decimal.Decimal, error) {
	expireSec := expireAtUnix - time.Now().Unix()
	if expireSec < 0 {
		expireSec = 0
	}
	return DecimalIncrBy(ctx, key, incr, expireSec)
}

// DecimalGet 读取 Redis 中以字符串存储的十进制值。
// key 不存在或为空时返回 decimal.Zero，避免调用方处理 redis.Nil。
func DecimalGet(ctx context.Context, key string) (decimal.Decimal, error) {
	if key == "" {
		return decimal.Zero, errors.New("redisx.DecimalGet: key is empty")
	}
	if Engine == nil {
		return decimal.Zero, errors.New("redisx.DecimalGet: Engine not initialized")
	}
	rdb := Engine.RDB(ctx)
	if rdb == nil {
		return decimal.Zero, errors.New("redisx.DecimalGet: no redis client for tenant")
	}
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return decimal.Zero, nil
		}
		return decimal.Zero, err
	}
	if res == "" {
		return decimal.Zero, nil
	}
	return decimal.NewFromString(res)
}
