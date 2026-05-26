package push

// 站点端会员信息上报

// SiteUserFullPayload 会员信息全量上报（首次创建时使用，所有字段都会上报，包括零值）
type SiteUserFullPayload struct {
	SiteId          string   `json:"site_id"`          // 站点ID
	UserId          string   `json:"user_id"`          // 用户ID
	ChannelId       string   `json:"channel_id"`       // 渠道ID
	Username        string   `json:"username"`         // 用户名
	RegisterTime    int64    `json:"register_time"`    // 注册时间
	RegisterCountry string   `json:"register_country"` // 注册国家 code_2
	RegisterIp      string   `json:"register_ip"`      // 注册IP
	RegisterSource  string   `json:"register_source"`  // 注册来源 H5 PC APP
	UserTags        []string `json:"user_tags"`        // 会员标签
	Phone           string   `json:"phone"`            // 手机号
	Email           string   `json:"email"`            // 邮箱
	Status          int32    `json:"status"`           // 账号状态
	SubStatus       []int32  `json:"sub_status"`       // 账号子状态
	VipLevel        int32    `json:"vip_level"`        // Vip等级
	LastLoginTime   int64    `json:"last_login_time"`  // 最后登录时间
	LastLoginIp     string   `json:"last_login_ip"`    // 最后登录IP
}

func (p SiteUserFullPayload) GetFilterableAttributes() []any {
	return []any{
		"site_id", "user_id", "channel_id", "username", "register_time", "register_country", "register_ip", "register_source", "user_tags", "phone", "email", "status", "sub_status", "vip_level", "last_login_time", "last_login_ip",
	}
}

// SiteUserListPayload 会员信息列表上报（增量更新时使用）
type SiteUserListPayload struct {
	UserId        string   `json:"user_id"`                   // 用户ID
	Username      string   `json:"username,omitempty"`        // 用户名
	UserTags      []string `json:"user_tags,omitempty"`       // 会员标签
	Phone         string   `json:"phone,omitempty"`           // 手机号
	Email         string   `json:"email,omitempty"`           // 邮箱
	Status        int32    `json:"status,omitempty"`          // 账号状态
	SubStatus     []int32  `json:"sub_status,omitempty"`      // 账号子状态
	VipLevel      int32    `json:"vip_level,omitempty"`       // Vip等级
	LastLoginTime int64    `json:"last_login_time,omitempty"` // 最后登录时间
	LastLoginIp   string   `json:"last_login_ip,omitempty"`   // 最后登录IP
}

// 站点端代理信息上报

// SiteAgentFullPayload 代理信息全量上报（首次创建时使用，所有字段都会上报，包括零值）
type SiteAgentFullPayload struct {
	SiteId          string `json:"site_id"`          // 站点ID
	UserId          string `json:"user_id"`          // 用户ID/代理ID
	Username        string `json:"username"`         // 用户名/代理用户名
	PromotionStatus int32  `json:"promotion_status"` // 推广开关状态：1-启用，2-禁用
	ParentUserID    string `json:"parent_user_id"`   // 父代理ID 直属上级
	TopUserID       string `json:"top_user_id"`      // 顶部代理ID
	Level           int32  `json:"level"`            // 所处层级
	Type            int32  `json:"type"`             // 代理类型：1-代理成员（无下级），2-代理团队（有下级）
	UpgradeTeamAt   int64  `json:"upgraded_team_at"` // 成为代理时间
	SubNumDirect    int32  `json:"sub_num_direct"`   // 直属下级总数 绑定
	SubNumIndirect  int32  `json:"sub_num_indirect"` // 非直属下级总数 绑定
	ChannelId       string `json:"channel_id"`       // 渠道ID
	RegisterSource  string `json:"register_source"`  // 注册来源 H5 PC APP
	LevelSubMax     int32  `json:"level_sub_max"`    // 下级最大层级 顶层代理数据
	AgentModeId     string `json:"agent_mode_id"`    // 代理模式id 顶层代理数据
	CreatedAt       int64  `json:"created_at"`       // 创建时间
}

func (p SiteAgentFullPayload) GetFilterableAttributes() []any {
	return []any{
		"site_id", "user_id", "username", "promotion_status", "parent_user_id", "top_user_id", "level", "type", "upgraded_team_at", "sub_num_direct", "sub_num_indirect", "channel_id", "register_source", "level_sub_max", "agent_mode_id", "created_at",
	}
}

// SiteAgentListPayload 代理信息列表上报（增量更新时使用）
type SiteAgentListPayload struct {
	UserId          string `json:"user_id"`                    // 用户ID/代理ID
	Username        string `json:"username,omitempty"`         // 用户名/代理用户名
	PromotionStatus int32  `json:"promotion_status,omitempty"` // 推广开关状态：1-启用，2-禁用
	Type            int32  `json:"type,omitempty"`             // 代理类型：1-代理成员（无下级），2-代理团队（有下级）
	SubNumDirect    int32  `json:"sub_num_direct,omitempty"`   // 直属下级总数 绑定
	SubNumIndirect  int32  `json:"sub_num_indirect,omitempty"` // 非直属下级总数 绑定
	LevelSubMax     int32  `json:"level_sub_max,omitempty"`    // 下级最大层级 顶层代理数据
	AgentModeId     string `json:"agent_mode_id,omitempty"`    // 代理模式id 顶层代理数据
}
