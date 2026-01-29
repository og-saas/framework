package interceptor

import (
	"context"
	"strconv"

	"github.com/og-saas/framework/utils/metadatakey"
	"github.com/og-saas/framework/utils/tenant"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TenantUnaryServerInterceptor() grpc.UnaryServerInterceptor {
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

		values := md.Get(metadatakey.TenantIdMetadataKey)
		if len(values) == 0 {
			return handler(ctx, req)
		}

		tenantId, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid tenant-id")
		}

		return handler(tenant.SetTenantId(ctx, tenantId), req)
	}
}

func TenantStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv any,
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return handler(srv, ss)
		}

		values := md.Get(metadatakey.TenantIdMetadataKey)
		if len(values) == 0 {
			return handler(srv, ss)
		}

		tenantId, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid tenant-id")
		}

		wrapped := &tenantServerStream{
			ServerStream: ss,
			ctx:          tenant.SetTenantId(ss.Context(), tenantId),
		}

		return handler(srv, wrapped)
	}
}

type tenantServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (t *tenantServerStream) Context() context.Context {
	return t.ctx
}

// TenantUnaryClientInterceptor 返回一个 Unary 客户端拦截器，自动在请求中注入 tenantId
func TenantUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy() // 避免污染原 context
		}

		md.Set(metadatakey.TenantIdMetadataKey, cast.ToString(tenant.GetTenantId(ctx)))
		ctx = metadata.NewOutgoingContext(ctx, md)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// TenantStreamClientInterceptor 返回一个 Stream 客户端拦截器，自动在请求中注入 tenantId
func TenantStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		md.Set(metadatakey.TenantIdMetadataKey, cast.ToString(tenant.GetTenantId(ctx)))
		ctx = metadata.NewOutgoingContext(ctx, md)

		return streamer(ctx, desc, cc, method, opts...)
	}
}
