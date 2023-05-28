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
	"github.com/lugavin/go-scaffold/internal/pkg/env"
	"github.com/lugavin/go-scaffold/pkg/httpserver"
)

// Run creates objects via constructors.
func Run(c *config.Config) {
	e, err := env.InitEnvironment(context.Background(), c)
	if err != nil {
		log.Fatalf("app - Run - env.InitEnvironment: %s", err)
	}
	defer e.Close()

	l := e.Logger()

	go e.KafkaConsumer().ConsumeTopic(
		context.Background(),
		[]string{c.KafkaTopics.FooBarTopic.TopicName},
		1,
		amqp.NewMessageHandler(l, c).HandleMessage,
	)

	// HTTP Server
	router := chi.NewRouter()
	v1.NewRouter(router, e)
	httpServer := httpserver.New(router, httpserver.Port(c.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		l.Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}
}
