package rocketmqx

import (
	"context"
	"strings"
	"time"

	rmqClient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	v2 "github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
	config2 "github.com/og-saas/framework/mq/rocketmqx/config"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// 无消息时的休眠时间
	noMessageSleepDuration = 200 * time.Millisecond
	// ACK 操作的超时时间
	ackTimeout = 5 * time.Second
	// 表达式类型：消息未找到
	messageNotFoundCode = v2.Code_MESSAGE_NOT_FOUND
)

type RocketMqx struct {
	config config2.Config
}

func NewRocketMqx(config config2.Config) *RocketMqx {
	rmqClient.ResetLogger()
	return &RocketMqx{config: config}
}

// 创建基础配置，减少重复代码
func (r *RocketMqx) createBaseConfig() *rmqClient.Config {
	return &rmqClient.Config{
		Endpoint:      r.config.Endpoint,
		NameSpace:     r.config.NameSpace,
		ConsumerGroup: r.config.ConsumerConfig.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:     r.config.AccessKey,
			AccessSecret:  r.config.AccessSecret,
			SecurityToken: r.config.SecurityToken,
		},
	}
}

func (r *RocketMqx) NewProducer(options ...ProducerOption) (producer rmqClient.Producer, err error) {
	var rocketmqOpts []rmqClient.ProducerOption
	for _, opt := range options {
		rocketmqOpts = opt(rocketmqOpts)
	}

	producer, err = rmqClient.NewProducer(
		r.createBaseConfig(),
		rocketmqOpts...,
	)
	if err != nil {
		logx.Errorf("NewProducer failed: %s", err.Error())
		return
	}
	// 启动生产者
	if err = producer.Start(); err != nil {
		logx.Errorf("Start producer failed: %s", err.Error())
	}
	return
}

func (r *RocketMqx) NewPullConsumer(handler config2.PullMessageHandler) (simpleConsumer rmqClient.SimpleConsumer, err error) {
	simpleConsumer, err = rmqClient.NewSimpleConsumer(
		r.createBaseConfig(),
		rmqClient.WithSimpleAwaitDuration(time.Duration(r.config.ConsumerConfig.AwaitDuration)*time.Second),
		rmqClient.WithSimpleSubscriptionExpressions(r.buildSubscriptionRelations()),
	)
	if err != nil {
		logx.Errorf("Initialize pull consumer failed: %s", err.Error())
		return
	}

	if err = simpleConsumer.Start(); err != nil {
		logx.Errorf("Start pull consumer failed: %s", err.Error())
		return
	}

	// 将消息处理逻辑提取到单独的函数中
	go r.processMessages(simpleConsumer, handler, r.getTopicNames())

	return
}

func (r *RocketMqx) NewPushConsumer(handler config2.PushMessageHandler) (pushConsumer rmqClient.PushConsumer, err error) {
	pushConsumer, err = rmqClient.NewPushConsumer(r.createBaseConfig(),
		rmqClient.WithPushAwaitDuration(time.Duration(r.config.ConsumerConfig.AwaitDuration)*time.Second),
		rmqClient.WithPushSubscriptionExpressions(r.buildSubscriptionRelations()),
		rmqClient.WithPushMessageListener(&rmqClient.FuncMessageListener{
			Consume: handler,
		}),
		rmqClient.WithPushConsumptionThreadCount(r.config.ConsumerConfig.PushConsumptionThreadCount),
		rmqClient.WithPushMaxCacheMessageCount(r.config.ConsumerConfig.PushMaxCacheMessageCount),
	)
	if err != nil {
		logx.Errorf("NewPushConsumer err: %s", err.Error())
		return
	}
	// start pushConsumer
	if err = pushConsumer.Start(); err != nil {
		logx.Errorf("Start pushConsumer err: %s", err.Error())
		return
	}
	return
}

// processMessages 处理消息的逻辑
func (r *RocketMqx) processMessages(consumer rmqClient.SimpleConsumer, handler config2.PullMessageHandler, topics string) {
	for {
		// 1. 拉取消息 - Receive超时设置为 AwaitDuration + 5秒buffer
		receiveCtx, receiveCancel := context.WithTimeout(
			context.Background(),
			time.Duration(r.config.ConsumerConfig.AwaitDuration+5)*time.Second,
		)
		mvs, err := consumer.Receive(
			receiveCtx,
			int32(r.config.ConsumerConfig.PullBatchSize),
			time.Duration(r.config.ConsumerConfig.InvisibleDuration)*time.Second,
		)
		receiveCancel()

		// 2. 处理拉取错误
		if err != nil {
			if strings.Contains(err.Error(), v2.Code_name[int32(messageNotFoundCode)]) {
				// 无消息时短暂休眠
				time.Sleep(noMessageSleepDuration)
				continue
			}
			logx.Errorf("Pull message failed, topics: %s, error: %s", topics, err.Error())
			continue
		}

		// 3. 打印收到的消息
		for _, mv := range mvs {
			logx.Debugf("Received message: consumerGroup=%s, topic=%s, msgId=%s, tag=%s, body=%s",
				consumer.GetGroupName(), mv.GetTopic(), mv.GetMessageId(), mv.GetTag(), string(mv.GetBody()))
		}

		// 4. 处理消息 - 使用 InvisibleDuration 作为处理超时
		handlerCtx, handlerCancel := context.WithTimeout(
			context.Background(),
			time.Duration(r.config.ConsumerConfig.InvisibleDuration)*time.Second,
		)

		res, err := handler(handlerCtx, mvs...)
		handlerCancel()

		// 4. ACK确认
		if res && err == nil {
			// 确认ACK - 5秒超时
			ackCtx, ackCancel := context.WithTimeout(context.Background(), ackTimeout)
			for _, mv := range mvs {
				if ackErr := consumer.Ack(ackCtx, mv); ackErr != nil {
					logx.Errorf("Ack message failed, messageID: %s, error: %s", mv.GetMessageId(), ackErr.Error())
				}
			}
			ackCancel()
		} else if err != nil {
			logx.Errorf("Process message failed, topics: %s, error: %s", topics, err.Error())
		}
	}
}

// getTopicNames 获取所有订阅的 Topic 名称（用于日志）
func (r *RocketMqx) getTopicNames() string {
	names := make([]string, 0, len(r.config.ConsumerConfig.TopicRelationList))
	for _, tr := range r.config.ConsumerConfig.TopicRelationList {
		names = append(names, tr.Topic)
	}
	return strings.Join(names, ",")
}

// buildSubscriptionRelations 构建订阅关系映射
func (r *RocketMqx) buildSubscriptionRelations() map[string]*rmqClient.FilterExpression {
	result := make(map[string]*rmqClient.FilterExpression, len(r.config.ConsumerConfig.TopicRelationList))
	for _, tr := range r.config.ConsumerConfig.TopicRelationList {
		result[tr.Topic] = rmqClient.NewFilterExpressionWithType(
			tr.Expression,
			rmqClient.FilterExpressionType(tr.ExpressionType),
		)
	}
	return result
}
