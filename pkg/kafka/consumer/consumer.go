package consumer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	_minBytes               = 10e3 // 10KB
	_maxBytes               = 10e6 // 10MB
	_queueCapacity          = 100
	_heartbeatInterval      = 3 * time.Second
	_commitInterval         = 0
	_partitionWatchInterval = 5 * time.Second
	_maxAttempts            = 3
	_dialTimeout            = 3 * time.Minute
	_maxWait                = 1 * time.Second
)

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type Consumer interface {
	ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker)
}

type consumer struct {
	logger  *zap.Logger
	Brokers []string
	GroupID string
}

func New(logger *zap.Logger, brokers []string, groupID string) *consumer {
	return &consumer{logger, brokers, groupID}
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumer) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	var errLogger kafka.LoggerFunc = func(msg string, args ...interface{}) {
		c.logger.Error(fmt.Sprintf(msg, args))
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                c.Brokers,
		GroupID:                c.GroupID,
		GroupTopics:            groupTopics,
		ErrorLogger:            errLogger,
		MinBytes:               _minBytes,
		MaxBytes:               _maxBytes,
		QueueCapacity:          _queueCapacity,
		HeartbeatInterval:      _heartbeatInterval,
		CommitInterval:         _commitInterval,
		PartitionWatchInterval: _partitionWatchInterval,
		MaxAttempts:            _maxAttempts,
		MaxWait:                _maxWait,
		Dialer: &kafka.Dialer{
			Timeout: _dialTimeout,
		},
	})

	defer func() {
		if err := r.Close(); err != nil {
			c.logger.Error("consumerGroup.r.Close", zap.Error(err))
		}
	}()

	c.logger.Info("Starting consumer",
		zap.String("groupID", c.GroupID),
		zap.Strings("topics", groupTopics),
		zap.Int("poolSize", poolSize))

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
