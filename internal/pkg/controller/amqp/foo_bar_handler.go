package amqp

import (
	"context"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type fooBarMessageHandler struct {
	logger *zap.Logger
}

func newFooBarMessageHandler(logger *zap.Logger) *fooBarMessageHandler {
	return &fooBarMessageHandler{logger}
}

func (h *fooBarMessageHandler) HandleTopicMessage(ctx context.Context, r *kafka.Reader, msg kafka.Message) {
	h.logger.Info("handleFooBarMessage...")
}
