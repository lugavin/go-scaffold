// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/lugavin/go-scaffold/config"
	v1 "github.com/lugavin/go-scaffold/internal/controller/http/v1"
	"github.com/lugavin/go-scaffold/internal/usecase"
	"github.com/lugavin/go-scaffold/internal/usecase/repo"
	"github.com/lugavin/go-scaffold/internal/usecase/webapi"
	"github.com/lugavin/go-scaffold/pkg/httpserver"
	"github.com/lugavin/go-scaffold/pkg/logger"
	"github.com/lugavin/go-scaffold/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	//ms, err := mysql.New(cfg.Mysql.URL, mysql.MaxIdleConns(cfg.Mysql.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - mysql.New: %w", err))
	//}
	//defer ms.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)
	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	// HTTP Server
	router := chi.NewRouter()
	v1.NewRouter(router, l, translationUseCase)
	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
