package xhttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/og-saas/framework/metadata"
	"github.com/og-saas/framework/utils/sign"
	"github.com/og-saas/framework/utils/xerr"
	v1 "github.com/og-saas/proto/pb/user/v1"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/status"
)

const (
	BusinessCodeOK = 0
	BusinessMsgOk  = "ok"

	serverTimeHeader = "X-Server-Time"
)

type BaseResponse[T any] struct {
	Code      int            `json:"code" xml:"code"`
	CodeGroup xerr.CodeGroup `json:"code_group,omitempty" xml:"code_group,omitempty"`
	Message   string         `json:"message" xml:"message"`
	Data      T              `json:"data,omitempty" xml:"data,omitempty"`
	TraceID   string         `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	Sign      string         `json:"sign,omitempty" xml:"sign,omitempty"`
}

// JsonBaseResponseCtx writes v into w with appropriate http status code.
func JsonBaseResponseCtx(ctx context.Context, w http.ResponseWriter, v any) {
	w.Header().Set(serverTimeHeader, fmt.Sprintf("%d", time.Now().Unix()))
	resp := wrapBaseResponse(ctx, v)
	if ok := metadata.DataEncrypt.GetBool(ctx); ok {
		resp.Sign = sign.SignParams(resp.TraceID, resp.Data)
		resp.Data = sign.AesEncrypt(resp.TraceID, resp.Data)
	}
	// 使用error 防止关闭info后看不见
	if resp.Code != BusinessCodeOK {
		logx.WithContext(ctx).Errorf("JsonBaseResponseCtx Code: %d response: %+v", resp.Code, resp)
	} else {
		logx.WithContext(ctx).Errorf("JsonBaseResponseCtx OK")
	}

	httpx.OkJsonCtx(ctx, w, resp)
}
func wrapBaseResponse(ctx context.Context, v any) BaseResponse[any] {
	var resp BaseResponse[any]
	var formatArgs []any
	switch data := v.(type) {
	case xerr.Error:
		resp.Code = data.Code.Int()
		resp.Message = data.Msg
		resp.Data = data.Data
		formatArgs = data.FormatArgs
	case errors.CodeMsg:
		resp.Code = data.Code
		resp.Message = data.Msg
	case *status.Status:
		resp.Code = int(data.Code())
		resp.Message = data.Message()
	case error:
		if st, ok := status.FromError(data); ok {
			resp.Code = int(st.Code())
			resp.Message = st.Message()
			if details := st.Details(); len(details) > 0 {
				if msg, ok1 := details[0].(*v1.StringList); ok1 {
					for _, arg := range msg.Items {
						formatArgs = append(formatArgs, arg)
					}
				}
			}
		} else {
			resp.Code = http.StatusInternalServerError
			resp.Message = data.Error()
		}
	default:
		resp.Code = BusinessCodeOK
		resp.Message = BusinessMsgOk
		resp.Data = v
	}

	if resp.Code != BusinessCodeOK {
		resp.Message = xerr.TransErrMsg(resp.Code, resp.Message, metadata.Language.GetString(ctx))
		if len(formatArgs) > 0 {
			resp.Message = fmt.Sprintf(resp.Message, formatArgs...)
		}
		resp.CodeGroup = xerr.CodeGroupToast
		if group, ok := xerr.ErrCodeGroupMap[xerr.ErrCode(resp.Code)]; ok {
			resp.CodeGroup = group
		}
	}

	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		resp.TraceID = spanCtx.TraceID().String()
	}
	return resp
}
