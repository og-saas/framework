package message_center

import (
	"fmt"
)

type Topic string

// 用户消息
const (
	TopicGlobalUser      Topic = "/global/user"        // 全局所有用户
	TopicSiteUser        Topic = "/site/%d/user"       // 站点所有用户
	TopicSiteUserSingle  Topic = "/site/%d/user/%v"    // 站点指定用户
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

func (t Topic) String() string {
	return string(t)
}

// BuildTopic 构建 Topic（通用方法）
// topic: Topic 模板
// args: 格式化参数
func (c *Client) BuildTopic(topic Topic, args ...interface{}) string {
	return c.config.AppKey + fmt.Sprintf(string(topic), args...)
}

// BuildCustomTopic 自定义消息
// 生成: {appKey}/custom/{customKey}
// customKey: 自定义标识，如 "chat/room001", "notification/order"
func (c *Client) BuildCustomTopic(customKey string) string {
	return c.BuildTopic(TopicCustom, customKey)
}
