package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"go.uber.org/zap"
)

const (
	_writerReadTimeout  = 10 * time.Second
	_writerWriteTimeout = 10 * time.Second
	_writerRequiredAcks = -1
	_writerMaxAttempts  = 3
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	w *kafka.Writer
}

func New(logger *zap.Logger, brokers []string) *producer {
	var errLogger kafka.LoggerFunc = func(msg string, args ...interface{}) {
		logger.Error(fmt.Sprintf(msg, args))
	}
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  _writerReadTimeout,
		WriteTimeout: _writerWriteTimeout,
		MaxAttempts:  _writerMaxAttempts,
		RequiredAcks: _writerRequiredAcks,
		Async:        false,
	}
	return &producer{w}
}

// PublishMessage Kafka producers will cache messages in memory when sending messages,
// and send the cached messages to Kafka brokers in batches once certain conditions are met to improve sending efficiency.
func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

// Close Call the writer.Close() method to close the Kafka producer connection.
// This ensures that all buffered messages are flushed to the Kafka broker and frees network connection resources.
func (p *producer) Close() error {
	return p.w.Close()
}
