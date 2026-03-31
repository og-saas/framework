# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Module Overview

This is `github.com/og-saas/framework` (Go 1.25.1) — a shared internal framework module providing multi-tenant infrastructure for go-zero microservices.

## Common Commands

```bash
# Run tests
go test ./...

# Run a single test
go test ./path/to/package -run TestName -v

# Tidy dependencies
go mod tidy
```

## Architecture

### Multi-Tenancy System

The framework implements pervasive tenant-awareness across all infrastructure layers:

**Tenant Context Flow:**
1. `TenantUnaryServerInterceptor` / `TenantStreamServerInterceptor` extract `tenant-id` from gRPC metadata
2. Tenant ID is stored in context via `tenant.SetTenantId(ctx, tenantId)`
3. All stores (DB/Redis/Cache) use `tenant.GetTenantId(ctx)` to route to tenant-specific resources
4. `TenantUnaryClientInterceptor` / `TenantStreamClientInterceptor` automatically inject tenant-id into outgoing RPC calls

**Skipping Tenant Isolation:**
- Use `tenant.SkipTenant(ctx)` to bypass tenant filtering (returns tenant ID 0)
- Check with `tenant.IsSkipTenant(ctx)`

### Store Engines (stores/)

All stores follow a tenant-aware engine pattern with `Must()` for default tenant and `MustTenant()` for multi-tenant setup:

**gormx (Database):**
- `gormx.Engine.DB(driver).WithContext(ctx)` returns tenant-scoped `*gorm.DB`
- Supports MySQL, PostgreSQL, ClickHouse via driver parameter
- `MultiTenantPlugin` automatically adds `WHERE tenant_id = ?` to queries/updates/deletes
- `ShardingPlugin` modifies table names based on tenant ID: `table_{tenantId % count}`
- Base model: `gormx.Model` (CreatedAt, UpdatedAt, DeletedAt with soft delete)

**redisx (Redis):**
- `redisx.Engine.RDB(ctx)` returns tenant-scoped `redis.UniversalClient`
- Supports standalone, sentinel, and cluster modes

**cachex (Distributed Cache):**
- `cachex.Engine(ctx)` returns tenant-scoped `*rockscache.Client`
- Built on rockscache (防击穿/防穿透)
- Key types: `CacheKey` (global), `TenantCacheKey` (auto-prefixes with tenant ID)
- Use `cachex.Fetch2()` and `cachex.TagAsDeleted2()` with typed keys

**etcdx (Config Center):**
- `etcdx.NewEtcd[T](config, log)` creates typed config loader
- `GetConfig()` retrieves current config, `Listener()` watches for changes

### Interceptors (utils/interceptor/)

**Server-side:**
- `TenantUnaryServerInterceptor()` / `TenantStreamServerInterceptor()` — extract tenant-id from metadata
- `ValidateUnaryServerInterceptor()` / `ValidateStreamServerInterceptor()` — call `Validate()` on requests

**Client-side:**
- `TenantUnaryClientInterceptor()` / `TenantStreamClientInterceptor()` — inject tenant-id into metadata

### Error Handling (utils/xerr/)

- `xerr.Error` struct with `Code`, `Data`, `Msg`
- `xerr.Must(config)` to enable multi-language error messages
- `err.GetMessage(language)` returns localized error text
- Standard codes: `ErrCodeParamError` (400), `ErrCodeUnauthorized` (401), `ErrCodeServerInternalError` (500), etc.
- Domain-specific codes: user (10001+), game (20001+), finance (30001+), agent (40001+)

### Metadata & Context (metadata/, utils/contextkey/)

**metadata package:**
- Defines standard metadata keys: `Authorization`, `UserId`, `TenantId`, `Language`, `Currency`, `DeviceId`, etc.
- `metadata.TenantId.GetInt64(ctx)` / `metadata.UserId.GetString(ctx)` for typed access
- `metadata.SetValues(ctx, map[Metadata]any)` for batch setting

**contextkey package:**
- Generic helpers: `contextkey.SetContext[T](ctx, key, val)` and `contextkey.GetContext[T](ctx, key)`

### Message Queue (mq/rocketmqx/)

- `rocketmqx.NewRocketMqx(config)` creates MQ client
- `NewProducer(options...)` for publishing
- `NewPullConsumer(handler)` for pull-based consumption (manual receive loop)
- `NewPushConsumer(handler)` for push-based consumption (automatic delivery)

### Utilities

**uniqueid:**
- `uniqueid.GenId()` generates Sonyflake-based unique IDs
- `uniqueid.GenOrderNO(prefix)` creates order numbers with prefix (e.g., "RO-123456789")
- `uniqueid.ExtractTime(id)` / `uniqueid.ExtractOrderTime(orderNo)` extract timestamps

**site_config:**
- `site_config.GetLanguageObject[T](items, language, defaultLanguage)` for multi-language content selection

**consts:**
- `StatusType` (Enable/Disable), `TransferType` (In/Out), `DeviceType`, `EndpointType`
- `OrderPrefix` constants: `RechargeOrder`, `WithdrawOrder`, `GameBet`, etc.
- `PtbCoin` type wrapping `decimal.Decimal` for platform currency

**schedule:**
- Defines `JobKey` and `HandlerName` constants for scheduled tasks

## Configuration Pattern

Services typically use two-layer config:
1. Local YAML with env placeholders (loaded via `conf.MustLoad(..., conf.UseEnv())`)
2. Dynamic config from etcd via `etcdx.NewEtcd[T](config, log).GetConfig()`

Tenant-specific DB/Redis configs are loaded via `TenantConfigProvider` interface and passed to `MustTenant()`.
