package amqp

import (
	"context"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/internal/pkg/env"
	"github.com/lugavin/go-scaffold/pkg/kafka/consumer"
)

type TopicMessageHandler interface {
	HandleTopicMessage(ctx context.Context, r *kafka.Reader, msg kafka.Message)
}

type MessageConsumer struct {
	logger   *zap.Logger
	consumer consumer.Consumer
	handlers map[string]TopicMessageHandler
}

func NewMessageConsumer(e *env.Environment) *MessageConsumer {
	var (
		logger = e.Logger()
		cfg    = e.Config()
	)
	return &MessageConsumer{
		logger:   logger,
		consumer: e.KafkaConsumer(),
		handlers: map[string]TopicMessageHandler{
			cfg.FooBarTopic.TopicName: newFooBarMessageHandler(logger, cfg),
		},
	}
}

func (c *MessageConsumer) Start() {
	for topic, handler := range c.handlers {
		go c.consumer.ConsumeTopic(
			context.Background(),
			[]string{topic},
			1,
			newMessageHandler(c.logger, handler).HandleMessage,
		)
	}
}
