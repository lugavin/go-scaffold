package kafka

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"

	"github.com/lugavin/go-scaffold/pkg/log"
)

// MessageProcessor processor methods must implement kafka.Worker func method interface
type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type messageProcessor struct {
	log log.Logger
}

func NewMessageProcessor(log log.Logger) *messageProcessor {
	return &messageProcessor{log: log}
}

func (s *messageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Error("workerID: %v, err: %v", workerID, err)
			continue
		}

		s.log.Info("process messages, topic: %s, key: %s, val: %s", m.Topic, string(m.Key), string(m.Value))

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
