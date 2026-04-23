package message_center

import "fmt"

type Topic string

// 用户消息
const (
	TopicGlobalUser      Topic = "/global/user"        // 全局所有用户
	TopicSiteUser        Topic = "/site/%d/user"       // 站点所有用户
	TopicSiteUserSingle  Topic = "/site/%d/user/%d"    // 站点指定用户
	TopicSiteUserTag     Topic = "/site/%d/tag/%s"     // 站点用户标签（如 VIP、新用户等）
	TopicSiteUserChannel Topic = "/site/%d/channel/%d" // 站点渠道
)

// 管理员消息
const (
	TopicGlobalAdmin     Topic = "/global/admin"     // 全局所有管理员
	TopicSiteAdmin       Topic = "/site/%d/admin"    // 站点所有管理员
	TopicSiteAdminSingle Topic = "/site/%d/admin/%s" // 站点指定管理员
)

// TopicCustom 自定义消息
const TopicCustom Topic = "/custom/%s"

// BuildTopic 构建 Topic（通用方法）
// topic: Topic 模板
// args: 格式化参数
func (c *Client) BuildTopic(topic Topic, args ...interface{}) string {
	return c.config.AppKey + fmt.Sprintf(string(topic), args...)
}

// ========== 全局消息 Topic ==========

// BuildGlobalUserTopic 全局用户消息
// 生成: {appKey}/global/user
func (c *Client) BuildGlobalUserTopic() string {
	return c.BuildTopic(TopicGlobalUser)
}

// BuildGlobalAdminTopic 全局管理员消息
// 生成: {appKey}/global/admin
func (c *Client) BuildGlobalAdminTopic() string {
	return c.BuildTopic(TopicGlobalAdmin)
}

// ========== 站点用户消息 Topic ==========

// BuildSiteUserTopic 站点所有用户
// 生成: {appKey}/site/{siteId}/user
func (c *Client) BuildSiteUserTopic(siteId int64) string {
	return c.BuildTopic(TopicSiteUser, siteId)
}

// BuildSiteUserSingleTopic 站点指定用户
// 生成: {appKey}/site/{siteId}/user/{userId}
func (c *Client) BuildSiteUserSingleTopic(siteId int64, userId int64) string {
	return c.BuildTopic(TopicSiteUserSingle, siteId, userId)
}

// BuildSiteUserTagTopic 站点用户标签
// 生成: {appKey}/site/{siteId}/tag/{tag}
// tag 示例: "vip", "new_user", "active"
func (c *Client) BuildSiteUserTagTopic(siteId int64, tag string) string {
	return c.BuildTopic(TopicSiteUserTag, siteId, tag)
}

// BuildSiteUserChannelTopic 站点渠道
// 生成: {appKey}/site/{siteId}/channel/{channelId}
// channelId: 渠道 ID
func (c *Client) BuildSiteUserChannelTopic(siteId int64, channelId int64) string {
	return c.BuildTopic(TopicSiteUserChannel, siteId, channelId)
}

// ========== 站点管理员消息 Topic ==========

// BuildSiteAdminTopic 站点所有管理员
// 生成: {appKey}/site/{siteId}/admin
func (c *Client) BuildSiteAdminTopic(siteId int64) string {
	return c.BuildTopic(TopicSiteAdmin, siteId)
}

// BuildSiteAdminSingleTopic 站点指定管理员
// 生成: {appKey}/site/{siteId}/admin/{adminId}
func (c *Client) BuildSiteAdminSingleTopic(siteId int64, adminId string) string {
	return c.BuildTopic(TopicSiteAdminSingle, siteId, adminId)
}

// ========== 自定义消息 Topic ==========

// BuildCustomTopic 自定义消息
// 生成: {appKey}/custom/{customKey}
// customKey: 自定义标识，如 "chat/room001", "notification/order"
func (c *Client) BuildCustomTopic(customKey string) string {
	return c.BuildTopic(TopicCustom, customKey)
}
