package push

import "github.com/shopspring/decimal"

// 站点端会员信息上报
type SiteUserListPayload struct {
	SiteId          string   `json:"site_id"`                    // 站点ID
	UserId          string   `json:"user_id"`                    // 用户ID
	ChannelId       string   `json:"channel_id,omitempty"`       // 渠道ID
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

func (p SiteUserListPayload) GetFilterableAttributes() []any {
	return []any{
		"site_id", "user_id", "channel_id", "username", "register_country", "register_ip", "register_source", "user_tags", "phone", "email", "status", "sub_status", "vip_level", "last_login_time",
	}
}

// 站点端代理信息上报
type SiteAgentListPayload struct {
	SiteId                   string          `json:"site_id"`                               // 站点ID
	UserId                   string          `json:"user_id"`                               // 用户ID/代理ID
	Username                 string          `json:"username,omitempty"`                    // 用户名/代理用户名
	PromotionStatus          int32           `json:"promotion_status,omitempty"`            // 推广开关状态：1-启用，2-禁用
	ParentUserID             string          `json:"parent_user_id,omitempty"`              // 父代理ID 直属上级
	TopUserID                string          `json:"top_user_id,omitempty"`                 // 顶部代理ID
	Level                    int32           `json:"level,omitempty"`                       // 所处层级
	FeeAmountClaimed         decimal.Decimal `json:"fee_amount_claimed,omitempty"`          // 已领取佣金金额
	FeeAmountDirectClaimed   decimal.Decimal `json:"fee_amount_direct_claimed,omitempty"`   // 直属佣金已领取金额
	FeeAmountIndirectClaimed decimal.Decimal `json:"fee_amount_indirect_claimed,omitempty"` // 非直属佣金已领取金额
	ContributeFeeValid       decimal.Decimal `json:"contribute_fee_valid,omitempty"`        // 贡献上级佣金 有效
	ContributeFeeInvalid     decimal.Decimal `json:"contribute_fee_invalid,omitempty"`      // 贡献上级佣金 无效
	UpgradeTeamAt            int64           `json:"upgraded_team_at,omitempty"`            // 成为代理时间
	SubNumDirect             int32           `json:"sub_num_direct,omitempty"`              // 直属下级总数 绑定
	SubNumIndirect           int32           `json:"sub_num_indirect,omitempty"`            // 非直属下级总数 绑定
	ChannelId                string          `json:"channel_id,omitempty"`                  // 渠道ID
	RegisterSource           string          `json:"register_source,omitempty"`             // 注册来源 H5 PC APP
	LevelSubMax              int32           `json:"level_sub_max,omitempty"`               // 下级最大层级 顶层代理数据
	AgentModeId              string          `json:"agent_mode_id,omitempty"`               // 代理模式id 顶层代理数据
}

func (p SiteAgentListPayload) GetFilterableAttributes() []any {
	return []any{
		"site_id", "user_id", "username", "promotion_status", "parent_user_id", "top_user_id", "level", "fee_amount_claimed",
		"fee_amount_direct_claimed", "fee_amount_indirect_claimed", "contribute_fee_valid", "contribute_fee_invalid", "upgraded_team_at", "sub_num_direct", "sub_num_indirect", "channel_id", "register_source", "level_sub_max", "agent_mode_id",
	}
}
