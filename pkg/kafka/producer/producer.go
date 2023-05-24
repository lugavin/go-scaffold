package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"

	"github.com/lugavin/go-scaffold/pkg/log"
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

func New(logger log.Logger, brokers []string) *producer {
	var errLogger kafka.LoggerFunc = func(msg string, args ...interface{}) {
		logger.Error(msg, args)
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

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
