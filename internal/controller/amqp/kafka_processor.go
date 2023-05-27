package amqp

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// KafkaProcessor processor methods must implement kafka.Worker func method interface
type KafkaProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type kafkaProcessor struct {
	logger *zap.Logger
}

func NewKafkaProcessor(logger *zap.Logger) *kafkaProcessor {
	return &kafkaProcessor{logger}
}

func (s *kafkaProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.logger.Error("kafkaProcessor.ProcessMessages", zap.Int("workerID", workerID), zap.Error(err))
			continue
		}

		s.logger.Info("kafkaProcessor.ProcessMessages",
			zap.String("topic", m.Topic),
			zap.ByteString("key", m.Key),
			zap.ByteString("value", m.Value))

		switch m.Topic {
		//case s.cfg.KafkaTopics.ProductCreated.TopicName:
		//	s.processProductCreated(ctx, r, m)
		//case s.cfg.KafkaTopics.ProductUpdated.TopicName:
		//	s.processProductUpdated(ctx, r, m)
		//case s.cfg.KafkaTopics.ProductDeleted.TopicName:
		//	s.processProductDeleted(ctx, r, m)
		}
	}
}
