package gormx

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/og-saas/framework/utils/contextkey"
	"github.com/og-saas/framework/utils/tenant"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dbRowsField     = "rows"
	dbDurationField = "duration"
)

const skipCaller = 3

type GormLogger struct {
	logger.Config
}

func NewLogger(cfg logger.Config) *GormLogger {
	return &GormLogger{
		Config: cfg,
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		log(ctx).Infow(fmt.Sprintf(msg, data...), logx.Field(contextkey.TenantKey.Name(), tenant.GetTenantId(ctx)))
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		log(ctx).Sloww(fmt.Sprintf(msg, data...), logx.Field(contextkey.TenantKey.Name(), tenant.GetTenantId(ctx)))
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		log(ctx).Errorw(fmt.Sprintf(msg, data...), logx.Field(contextkey.TenantKey.Name(), tenant.GetTenantId(ctx)))
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []logx.LogField{
		logx.Field(dbRowsField, rows),
		logx.Field(dbDurationField, float64(elapsed.Nanoseconds())/1e6),
	}

	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		fields = append(fields, logx.Field("err", err))
		log(ctx).Errorw(sql, fields...)
	case elapsed > l.SlowThreshold && l.SlowThreshold > 0 && l.LogLevel >= logger.Warn:
		log(ctx).Sloww(sql, fields...)
	case l.LogLevel == logger.Info:
		log(ctx).Infow(sql, fields...)
	}
}

func log(ctx context.Context) logx.Logger {
	return logx.WithContext(ctx).WithCallerSkip(fileIndex())
}

func fileIndex() int {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	length := runtime.Callers(4, pcs[:])
	frames := runtime.CallersFrames(pcs[:length])
	for i := 0; i < length; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()
		if (!strings.Contains(frame.File, "/gorm.io/") ||
			strings.HasSuffix(frame.File, "_test.go")) && !strings.HasSuffix(frame.File, ".gen.go") && !strings.Contains(frame.File, "/dao/") {
			return i + 1
		}
	}

	return 1
}
