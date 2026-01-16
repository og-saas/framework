package mq

import (
	"os"
)

var (
	// TopicUserWalletTransferNotify 用户钱包交易通知
	TopicUserWalletTransferNotify = "user_wallet_transfer_notify"
)

func UpdateTopicPrefix(prefixes ...string) (prefix string) {
	if len(prefixes) > 0 && prefixes[0] != "" {
		prefix = prefixes[0]
	}
	if prefix == "" {
		prefix = os.Getenv("ROCKETMQ_TOPIC_PREFIX")
	}

	TopicUserWalletTransferNotify = prefix + TopicUserWalletTransferNotify

	return
}
