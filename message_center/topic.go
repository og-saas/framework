package message_center

import (
	"fmt"
)

type Topic string

// 设备消息
const (
	TopicGlobalDevice     Topic = "/global/device"     // 全局所有设备
	TopicSiteDevice       Topic = "/site/%d/device"    // 站点所有设备, /site/{站点ID}/device
	TopicSiteDeviceSingle Topic = "/site/%d/device/%s" // 站点指定设备，/site/{站点ID}/device/{设备ID}
)

// 用户消息
const (
	TopicGlobalUser      Topic = "/global/user"        // 全局所有用户
	TopicSiteUser        Topic = "/site/%d/user"       // 站点所有用户，/site/{站点ID}/user
	TopicSiteUserSingle  Topic = "/site/%d/user/%d"    // 站点指定用户，/site/{站点ID}/user/{用户ID}
	TopicSiteUserChannel Topic = "/site/%d/channel/%d" // 站点渠道，/site/{站点ID}/channel/{渠道ID}
)

// 管理员消息
const (
	TopicGlobalAdmin     Topic = "/global/admin"     // 全局所有管理员，/global/admin
	TopicSiteAdmin       Topic = "/site/%d/admin"    // 站点所有管理员，/site/{站点ID}/admin
	TopicSiteAdminSingle Topic = "/site/%d/admin/%d" // 站点指定管理员，/site/{站点ID}/admin/{管理员ID}
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
