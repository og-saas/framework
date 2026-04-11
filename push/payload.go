package push

// 站点端会员列表信息上报
type SiteUserListPayload struct {
	SiteId          int64    `json:"site_id"`                    // 站点ID
	UserId          int64    `json:"user_id"`                    // 用户ID
	ChannelId       int64    `json:"channel_id,omitempty"`       // 渠道ID
	Username        string   `json:"username,omitempty"`         // 用户名
	RegisterTime    int64    `json:"register_time,omitempty"`    // 注册时间
	RegisterCountry string   `json:"register_country,omitempty"` // 注册国家 code_2
	RegisterIp      string   `json:"register_ip,omitempty"`      // 注册IP
	RegisterSource  string   `json:"register_source,omitempty"`  // 注册来源 H5 PC APP
	UserTags        []string `json:"user_tags,omitempty"`        // 会员标签
	Phone           string   `json:"phone,omitempty"`            // 手机号
	Email           string   `json:"email,omitempty"`            // 邮箱
	Status          int32    `json:"status,omitempty"`           // 账号状态
	SubStatus       []int32  `json:"sub_status,omitempty"`       // 账号子状态
	VipLevel        int32    `json:"vip_level,omitempty"`        // Vip等级
	LastLoginTime   int64    `json:"last_login_time,omitempty"`  // 最后登录时间
	LastLoginIp     string   `json:"last_login_ip,omitempty"`    // 最后登录IP
}

func (p *SiteUserListPayload) GetFilterableAttributes() []any {
	return []any{
		"site_id", "user_id", "channel_id", "username", "register_country", "register_ip", "register_source", "user_tags", "phone", "email", "status", "sub_status", "vip_level", "last_login_time",
	}
}
