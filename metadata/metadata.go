package metadata

import (
	"context"
	"fmt"

	"github.com/spf13/cast"
	md "google.golang.org/grpc/metadata"
)

type Metadata string

const (
	Authorization   Metadata = "Authorization"    // 授权
	XTenantId       Metadata = "X-Tenant-Id"      // 租户id
	UserId          Metadata = "user_id"          // 用户id
	Username        Metadata = "username"         // 用户名
	ChannelId       Metadata = "channel_id"       // 渠道id
	Language        Metadata = "language"         // 语言
	IP              Metadata = "ip"               // ip
	Currency        Metadata = "currency"         // 币种
	Domain          Metadata = "domain"           // 域名
	DeviceId        Metadata = "Device-Id"        // 设备id
	DeviceType      Metadata = "Device-Type"      // 设备类型
	DeviceOS        Metadata = "Device-Os"        // 设备操作系统
	AppVersion      Metadata = "App-Version"      // app版本
	UserAgent       Metadata = "User-Agent"       // 浏览器用户代理
	DefaultCurrency Metadata = "default_currency" // 钱包默认币种
	DeviceEndpoint  Metadata = "Device-Endpoint"  // 设备终端类型 APP H5 PC
)

// RpcMetadata 同步到下游服务的Metadata
var RpcMetadata = []Metadata{
	XTenantId,
	UserId,
	Username,
	ChannelId,
	Language,
	IP,
	Currency,
	Domain,
	DeviceId,
	DeviceType,
	DeviceOS,
	AppVersion,
	UserAgent,
	DefaultCurrency,
	DeviceEndpoint,
}

// GetKey 获取元数据key
func (s Metadata) GetKey() string {
	return string(s)
}

// GetValue 获取元数据
func (s Metadata) GetValue(ctx context.Context) any {
	return ctx.Value(s)
}

// GetString 获取元数据字符串
func (s Metadata) GetString(ctx context.Context) string {
	return cast.ToString(ctx.Value(s))
}

// GetMetadataKey 防止同名被框架覆盖 添加og前缀
func (s Metadata) GetMetadataKey() string {
	return fmt.Sprintf("og-%s", s)
}

func (s Metadata) GetFromContentOrMetadata(ctx context.Context) string {
	mdData, ok := md.FromIncomingContext(ctx)
	if !ok {
		return cast.ToString(ctx.Value(s))
	}

	values := mdData.Get(s.GetMetadataKey())
	if len(values) == 0 {
		return cast.ToString(ctx.Value(s))
	}

	return values[0]
}

// GetInt64 获取元数据int64
func (s Metadata) GetInt64(ctx context.Context) int64 {
	return cast.ToInt64(ctx.Value(s))
}

// SetValue 设置元数据
func (s Metadata) SetValue(ctx context.Context, val any) context.Context {
	return context.WithValue(ctx, s, val)
}

// SetValues 批量设置元数据
func SetValues(ctx context.Context, vals map[Metadata]any) context.Context {
	for k, v := range vals {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
