package shared

import "time"
import commonv1 "github.com/og-saas/proto/pb/common/v1"

// 风控指标查询请求和响应结构
// API: POST /api/risk/metric

const (
	AccessTokenPath    = "/user-server/open/api/token"                // AccessTokenPath 用户服务获取token接口路径
	RiskMetricPath     = "/data-api-server/api/risk/metric"           // RiskMetricPath 大数据风控指标查询接口路径
	RiskMetricPagePath = "/data-api-server/api/risk/metric/page/user" // RiskMetricPagePath 大数据风控指标人员分页列表接口路径
)

// 大数据请求基础配置
type BigDataConfig struct {
	BaseURL      string `json:"base_url"`
	AppID        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	TokenExpired int64  `json:"token_expired,optional"` // 过期时间 秒
}

// 大数据接口特殊响应code
const (
	BigDataCodeSuccess      = 0
	BigDataCodeTokenExpired = 11012
)

func (c BigDataConfig) TokenTTL() time.Duration {
	if c.TokenExpired <= 0 {
		return 24 * time.Hour
	}
	return time.Duration(c.TokenExpired) * time.Second
}

// RiskMetricCode 指标code枚举
type RiskMetricCode string

const (
	RiskMetricRegisterIP       RiskMetricCode = "register_ip_cnt"        // 注册IP频次
	RiskMetricRegisterDevice   RiskMetricCode = "register_device_cnt"    // 注册设备频次
	RiskMetricRegisterUsername RiskMetricCode = "register_user_name_cnt" // 注册用户名频次
	RiskMetricLoginIP          RiskMetricCode = "login_ip_cnt"           // 登录IP频次
	RiskMetricLoginDevice      RiskMetricCode = "login_device_cnt"       // 登录设备频次
	RiskMetricLoginUsername    RiskMetricCode = "login_user_name_cnt"    // 登录用户名频次
	RiskMetricWithdrawIP       RiskMetricCode = "withdraw_ip_cnt"        // 提现IP频次
	RiskMetricWithdrawDevice   RiskMetricCode = "withdraw_device_cnt"    // 提现设备频次
	RiskMetricWithdrawUsername RiskMetricCode = "withdraw_user_name_cnt" // 提现用户名频次
)

func (c RiskMetricCode) String() string {
	return string(c)
}

// 根据行为和限制对象获取对应的指标code
func GetRiskMetricCode(behavior commonv1.RiskBehavior, targetType commonv1.AclTargetType) RiskMetricCode {
	switch behavior {
	case commonv1.RiskBehavior_RISK_BEHAVIOR_REGISTER:
		switch targetType {
		case commonv1.AclTargetType_ACL_TARGET_TYPE_IP:
			return RiskMetricRegisterIP
		case commonv1.AclTargetType_ACL_TARGET_TYPE_DEVICE:
			return RiskMetricRegisterDevice
		case commonv1.AclTargetType_ACL_TARGET_TYPE_USER:
			return RiskMetricRegisterUsername
		}
	case commonv1.RiskBehavior_RISK_BEHAVIOR_LOGIN:
		switch targetType {
		case commonv1.AclTargetType_ACL_TARGET_TYPE_IP:
			return RiskMetricLoginIP
		case commonv1.AclTargetType_ACL_TARGET_TYPE_DEVICE:
			return RiskMetricLoginDevice
		case commonv1.AclTargetType_ACL_TARGET_TYPE_USER:
			return RiskMetricLoginUsername
		}
	case commonv1.RiskBehavior_RISK_BEHAVIOR_WITHDRAW:
		switch targetType {
		case commonv1.AclTargetType_ACL_TARGET_TYPE_IP:
			return RiskMetricWithdrawIP
		case commonv1.AclTargetType_ACL_TARGET_TYPE_DEVICE:
			return RiskMetricWithdrawDevice
		case commonv1.AclTargetType_ACL_TARGET_TYPE_USER:
			return RiskMetricWithdrawUsername
		}
	}
	return ""
}

// RiskQueryField 维度属性枚举
type RiskQueryField string

const (
	RiskQueryFieldIP       RiskQueryField = "ip"        // IP地址
	RiskQueryFieldDeviceID RiskQueryField = "device_id" // 设备ID
	RiskQueryFieldUsername RiskQueryField = "username"  // 用户名
)

// 根据限制对象获取对应的维度属性
func GetRiskQueryField(targetType commonv1.AclTargetType) RiskQueryField {
	switch targetType {
	case commonv1.AclTargetType_ACL_TARGET_TYPE_USER:
		return RiskQueryFieldUsername
	case commonv1.AclTargetType_ACL_TARGET_TYPE_IP:
		return RiskQueryFieldIP
	case commonv1.AclTargetType_ACL_TARGET_TYPE_DEVICE:
		return RiskQueryFieldDeviceID
	default:
		return ""
	}
}

func (f RiskQueryField) QueryMode() QueryMode {
	if f == RiskQueryFieldUsername {
		return QueryModeLeftLike
	}
	return QueryModeEq
}

// QueryMode 查询模式枚举 eq == left_like 左模糊匹配 right_like 右模糊匹配
type QueryMode string

const (
	QueryModeEq        QueryMode = "eq"
	QueryModeLeftLike  QueryMode = "left_like"
	QueryModeRightLike QueryMode = "right_like"
)

// RiskMetricItem 单条指标查询
type RiskMetricItem struct {
	Code            RiskMetricCode `json:"code"`              // 指标code
	QueryFiled      RiskQueryField `json:"query_filed"`       // 维度属性（ip、username、device_id）
	QueryFiledValue string         `json:"query_filed_value"` // 维度值
	QueryMode       QueryMode      `json:"query_mode"`        // 查询模式

}

// StatTimeConfig 统计时间配置
type StatTimeConfig struct {
	FormatType          string `json:"format_type"`            // 指标统计时间粒度（day）
	CloseAutoHandleTime string `json:"close_auto_handle_time"` // 是否关闭自动补齐（true）
	ParamsZoneCode      string `json:"params_zone_code"`       // 查看数据时区（如 UTC-05:00）
	QueryTimeOffset     string `json:"query_time_offset"`      // 是否打开分区偏移量（true）
}

// RiskMetricReq 风控指标总计查询请求
type RiskMetricReq struct {
	StartTime string           `json:"start_time"` // 开始时间（yyyy-MM-dd HH:mm:ss）
	EndTime   string           `json:"end_time"`   // 结束时间（yyyy-MM-dd HH:mm:ss）
	SiteID    string           `json:"site_id"`    // 站点ID
	Metrics   []RiskMetricItem `json:"metric"`     // 指标列表
}

// RiskMetricResp 风控指标查询响应
type RiskMetricResp struct {
	Code    int                `json:"code"`    // 接口状态码（0成功）
	Msg     string             `json:"msg"`     // 接口消息
	Data    RiskMetricRespData `json:"data"`    // 指标数据（key=指标code, value=次数）
	TraceID string             `json:"traceId"` // 链路追踪ID
}
type RiskMetricRespData map[string]int64

// RiskMetricPageReq 风控指标人员分页列表请求
// API: POST /api/risk/metric/page/user
type RiskMetricPageReq struct {
	StartTime string         `json:"start_time"` // 开始时间（yyyy-MM-dd HH:mm:ss）
	EndTime   string         `json:"end_time"`   // 结束时间（yyyy-MM-dd HH:mm:ss）
	SiteID    string         `json:"site_id"`    // 站点ID
	Current   int            `json:"current"`    // 当前页
	Size      int            `json:"size"`       // 分页大小
	Metric    RiskMetricItem `json:"metric"`     // 指标查询条件
}

// RiskMetricPageData 分页数据
type RiskMetricPageData struct {
	Records []RiskMetricPageItem `json:"records"` // 用户ID列表
}
type RiskMetricPageItem struct {
	UserID  int64 `json:"user_id"`  // 用户ID
	FirstAt int64 `json:"first_at"` // 首次时间 毫秒
}

// RiskMetricPageResp 风控指标人员分页列表响应
type RiskMetricPageResp struct {
	Code    int                `json:"code"`    // 接口状态码（0成功）
	Msg     string             `json:"msg"`     // 接口消息
	Data    RiskMetricPageData `json:"data"`    // 分页数据
	Total   int                `json:"total"`   // 总数
	Size    int                `json:"size"`    // 页大小
	Current int                `json:"current"` // 当前页
	Pages   int                `json:"pages"`   // 总页数
	TraceID string             `json:"traceId"` // 链路追踪ID
}
