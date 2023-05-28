package amqprpc

import (
	"github.com/lugavin/go-scaffold/internal/pkg/usecase"
	"github.com/lugavin/go-scaffold/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
