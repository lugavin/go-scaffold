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
	"github.com/lugavin/go-scaffold/internal/controller/amqp"
	"github.com/lugavin/go-scaffold/internal/controller/http/v1"
	"github.com/lugavin/go-scaffold/internal/logger"
	"github.com/lugavin/go-scaffold/internal/usecase"
	"github.com/lugavin/go-scaffold/internal/usecase/repo"
	"github.com/lugavin/go-scaffold/internal/usecase/webapi"
	"github.com/lugavin/go-scaffold/pkg/httpserver"
	"github.com/lugavin/go-scaffold/pkg/kafka/consumer"
	"github.com/lugavin/go-scaffold/pkg/kafka/producer"
	"github.com/lugavin/go-scaffold/pkg/mysql"
	"github.com/lugavin/go-scaffold/pkg/redis"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger)
	defer l.Sync()

	// Repository
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	log.Fatalf("app - Run - postgres.New: %s", err)
	//}
	//defer pg.Close()

	ms, err := mysql.New(cfg.Mysql.URL, mysql.MaxIdleConns(cfg.Mysql.PoolMax))
	if err != nil {
		log.Fatalf("app - Run - mysql.New: %s", err)
	}
	defer ms.Close()

	redisCli := redis.New(cfg.Redis.Addrs, redis.PoolSize(cfg.Redis.PoolSize))
	defer redisCli.Close() // nolint: errcheck

	p := producer.New(l, cfg.Kafka.Brokers)
	defer p.Close()

	c := consumer.New(l, cfg.Kafka.Brokers, cfg.App.Name)
	go c.ConsumeTopic(context.Background(), []string{cfg.KafkaTopics.FooBarTopic.TopicName}, 1, amqp.NewKafkaProcessor(l).ProcessMessages)

	// Use case
	translationUseCase := usecase.NewTranslationUseCase(
		//repo.NewTranslationPgRepo(pg),
		repo.NewTranslationRepo(ms),
		webapi.New(),
	)

	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)
	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	log.Fatalf("app - Run - rmqServer - server.New: %s", err)
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
		l.Error("app - Run - httpServer.Notify", zap.Error(err))
		//case err = <-rmqServer.Notify():
		//	l.Error("app - Run - rmqServer.Notify", zap.Error(err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		l.Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}

	//if err = rmqServer.Shutdown(); err != nil {
	//	l.Error("app - Run - rmqServer.Shutdown", zap.Error(err))
	//}
}
