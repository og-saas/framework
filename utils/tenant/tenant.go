package tenant

import (
	"context"
	"github.com/og-saas/framework/utils/contextkey"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

const Default int64 = 0
const MetaKey = "tenant-id"

func GetTenantId(ctx context.Context) int64 {
	val := contextkey.GetContext[any](ctx, contextkey.TenantKey)
	return cast.ToInt64(val)
}

func SetTenantId(ctx context.Context, tenantId int64) context.Context {
	return contextkey.SetContext(ctx, contextkey.TenantKey, tenantId)
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		values := md.Get(MetaKey)
		if len(values) == 0 {
			return handler(ctx, req)
		}

		id, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid tenant-id")
		}

		return handler(SetTenantId(ctx, id), req)
	}
}
