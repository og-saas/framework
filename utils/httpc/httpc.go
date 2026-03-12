package httpc

import (
	"context"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
)

// 定义http 引擎
var engine *resty.Client
var once sync.Once

func Do(ctx context.Context) *resty.Request {
	once.Do(func() {
		engine = MustClient()
	})
	return engine.R().SetContext(ctx).SetDebug(true).SetLogger(newCtxLogger(ctx))
}

func New(ctx context.Context, fs ...func(cli *resty.Client)) *resty.Request {
	client := MustClient()
	for _, f := range fs {
		f(client)
	}
	return client.R().SetContext(ctx)
}

// MustClient new http client
func MustClient() *resty.Client {
	return resty.New()
}

// ctxLogger 将 resty debug 日志桥接到 go-zero logx，自动携带 trace_id
type ctxLogger struct {
	ctx context.Context
	l   logx.Logger
}

func newCtxLogger(ctx context.Context) *ctxLogger {
	return &ctxLogger{ctx: ctx, l: logx.WithContext(ctx)}
}

func (c *ctxLogger) Errorf(format string, v ...interface{}) {
	c.l.Errorf("[httpc] "+format, v...)
}

func (c *ctxLogger) Warnf(format string, v ...interface{}) {
	c.l.Slowf("[httpc] "+format, v...)
}

func (c *ctxLogger) Debugf(format string, v ...interface{}) {
	c.l.Infof("[httpc] trace_id=%s "+format, append([]interface{}{trace.TraceIDFromContext(c.ctx)}, v...)...)
}
