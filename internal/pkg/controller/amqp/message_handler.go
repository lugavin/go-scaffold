package amqp

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
)

// MessageHandler handler methods must implement kafka.Worker func method interface
type MessageHandler interface {
	HandleMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type messageHandler struct {
	logger *zap.Logger
	config *config.Config
}

func NewMessageHandler(logger *zap.Logger, config *config.Config) *messageHandler {
	return &messageHandler{logger, config}
}

func (h *messageHandler) HandleMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	// 循环读取 Kafka 消息
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// 从 Kafka 中读取一条消息
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			h.logger.Error("messageHandler.HandleMessage", zap.Int("workerID", workerID), zap.Error(err))
			continue
		}

		h.logger.Info("messageHandler.HandleMessage",
			zap.String("topic", msg.Topic),
			zap.ByteString("key", msg.Key),
			zap.ByteString("value", msg.Value))

		switch msg.Topic {
		case h.config.KafkaTopics.FooBarTopic.TopicName:
			h.handleFooBarMessage(ctx, r, msg)
		default:
		}

		// 手动提交消息的偏移量
		if err = r.CommitMessages(ctx, msg); err != nil {
			h.logger.Error("messageHandler.r.CommitMessages", zap.Error(err))
		}
	}
}

func (h *messageHandler) handleFooBarMessage(ctx context.Context, r *kafka.Reader, msg kafka.Message) {
	h.logger.Info("handleFooBarMessage...")
}
