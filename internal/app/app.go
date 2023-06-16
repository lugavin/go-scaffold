// Package app configures and runs application.
package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/pkg/controller/amqp"
	http "github.com/lugavin/go-scaffold/internal/pkg/controller/http/v1"
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

	amqp.NewMessageConsumer(env).Start()

	// HTTP Server
	httpServer := httpserver.New(http.NewRouter(env), httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// If none of the channels in the case statement are available, the select statement will block the current goroutine until a channel is available.
	select {
	case s := <-interrupt:
		// Server was interrupted
		env.Logger().Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		// Server failed to start
		env.Logger().Error("app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		env.Logger().Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}
}
