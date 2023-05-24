package consumer

import (
	"context"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/lugavin/go-scaffold/pkg/log"
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
	logger  log.Logger
	Brokers []string
	GroupID string
}

func New(logger log.Logger, brokers []string, groupID string) *consumer {
	return &consumer{logger, brokers, groupID}
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumer) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	var errLogger kafka.LoggerFunc = func(msg string, args ...interface{}) {
		c.logger.Error(msg, args)
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
			c.logger.Warn("consumerGroup.r.Close: %v", err)
		}
	}()

	c.logger.Info("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
