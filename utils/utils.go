package utils

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	mathrand "math/rand/v2"
	"strings"

	"github.com/og-saas/framework/utils/consts"
	"github.com/og-saas/framework/utils/tenant"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

const base36Chars = "0123456789abcdefghijklmnopqrstuvwxyz"

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

// Base36Encode id转base36
func Base36Encode(id int64) string {
	if id == 0 {
		return "0"
	}
	base := int64(len(base36Chars))
	result := ""

	for id > 0 {
		remainder := id % base
		result = string(base36Chars[remainder]) + result
		id /= base
	}
	return result
}

// Base36Decode base36转id
func Base36Decode(code string) int64 {
	if code == "" {
		return 0
	}

	base := int64(len(base36Chars))
	var id int64

	charIndex := make(map[rune]int64)
	for i, ch := range base36Chars {
		charIndex[ch] = int64(i)
	}

	for _, ch := range code {
		index, ok := charIndex[ch]
		if !ok {
			return 0
		}
		id = id*base + index
	}

	return id
}

func init() {
	otel.SetTracerProvider(noop.NewTracerProvider())
}

func WithTraceCtx(ctx context.Context) context.Context {
	tracer := otel.Tracer("framework")
	ctx, span := tracer.Start(ctx, "WithTraceCtx")
	defer span.End()
	traceID := span.SpanContext().TraceID().String()
	return logx.ContextWithFields(ctx, logx.Field(consts.TraceID, traceID))
}

func NewFromContext(ctx context.Context) context.Context {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return trace.ContextWithSpan(context.Background(), span)
	}
	return WithTraceCtx(ctx)
}

// DetachContext 创建不受上游取消/超时限制的新上下文，保留 trace 链路和租户信息
// 适用于异步调用、后台任务等需要脱离请求生命周期的场景
// 如果上游没有 trace，自动注入新的 trace ID
func DetachContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	newCtx := context.Background()

	// 继承或创建 trace
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		newCtx = trace.ContextWithSpanContext(newCtx, spanCtx)
	} else {
		// 生成新 trace ID
		var traceID trace.TraceID
		_, _ = rand.Read(traceID[:])
		sc := trace.NewSpanContext(trace.SpanContextConfig{
			TraceID:    traceID,
			TraceFlags: trace.FlagsSampled,
		})
		newCtx = trace.ContextWithSpanContext(newCtx, sc)
	}

	// 继承租户信息
	if tenantID := tenant.GetTenantId(ctx); tenantID != 0 {
		newCtx = tenant.SetTenantId(newCtx, tenantID)
	}

	return newCtx
}

// DetachContextWithSpan 创建新上下文并启动 span，由调用方 defer span.End()
// spanName 建议格式："{service}_{operation}"，例如 "agent_process_bet_notify"
func DetachContextWithSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if stringx.HasEmpty(spanName) {
		spanName = "unknown_operation"
	}

	// 继承 trace 链路
	spanCtx := trace.SpanContextFromContext(ctx)

	// 创建新 span
	name := strings.Split(spanName, "_")[0]
	tracer := otel.Tracer(name)
	newCtx, span := tracer.Start(trace.ContextWithSpanContext(context.Background(), spanCtx), spanName)

	// 注入日志字段
	newCtx = logx.ContextWithFields(newCtx,
		logx.Field(consts.TraceID, span.SpanContext().TraceID().String()),
		logx.Field(consts.SpanID, span.SpanContext().SpanID().String()),
	)

	// 继承租户信息
	if tenantID := tenant.GetTenantId(ctx); tenantID != 0 {
		newCtx = tenant.SetTenantId(newCtx, tenantID)
	}

	return newCtx, span
}

// RandomInt64 生成 [min, max] 范围内的随机整数（包含边界）
func RandomInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return min
	}
	return min + mathrand.Int64N(max-min+1)
}
