package amqp

import (
	"context"

	"github.com/lugavin/go-scaffold/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type fooBarMessageHandler struct {
	logger *zap.Logger
	config *config.Config
}

func NewFooBarMessageHandler(logger *zap.Logger, config *config.Config) *fooBarMessageHandler {
	return &fooBarMessageHandler{logger, config}
}

func (h *fooBarMessageHandler) HandleTopicMessage(ctx context.Context, r *kafka.Reader, msg kafka.Message) {
	h.logger.Info("handleFooBarMessage...")
}
