// Package app configures and runs application.
package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/pkg/controller/amqp"
	"github.com/lugavin/go-scaffold/internal/pkg/controller/http/v1"
	envt "github.com/lugavin/go-scaffold/internal/pkg/env"
	"github.com/lugavin/go-scaffold/pkg/httpserver"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	env, err := envt.InitEnvironment(context.Background(), cfg)
	if err != nil {
		log.Fatalf("app - Run - env.InitEnvironment: %s", err)
	}
	defer env.Close()

	logger := env.Logger()

	go env.KafkaConsumer().ConsumeTopic(
		context.Background(),
		[]string{cfg.KafkaTopics.FooBarTopic.TopicName},
		1,
		amqp.NewMessageHandler(logger, cfg).HandleMessage,
	)

	// HTTP Server
	router := chi.NewRouter()
	v1.NewRouter(router, env)
	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// If none of the channels in the case statement are available, the select statement will block the current goroutine until a channel is available
	select {
	case s := <-interrupt:
		// Server was interrupted
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		// Server failed to start
		logger.Error("app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		logger.Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}
}
