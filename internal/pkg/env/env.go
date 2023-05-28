package env

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase/repo"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase/webapi"
	"github.com/lugavin/go-scaffold/pkg/kafka/consumer"
	"github.com/lugavin/go-scaffold/pkg/kafka/producer"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

type Environment struct {
	rootContext        context.Context
	config             *config.Config
	logger             *zap.Logger
	mysql              *mysql.Mysql
	redisCli           redis.UniversalClient
	kafkaProducer      producer.Producer
	kafkaConsumer      consumer.Consumer
	translationRepo    *repo.TranslationRepo
	translationWebAPI  *webapi.TranslationWebAPI
	translationUseCase *usecase.TranslationUseCase
}

func InitEnvironment(ctx context.Context, cfg *config.Config) (*Environment, error) {
	env, err := initEnvironment(ctx, cfg)
	if err != nil {
		return nil, err
	}
	env.rootContext = ctx
	env.config = cfg
	return env, nil
}

func (e *Environment) Close() {
	if e.kafkaProducer != nil {
		e.kafkaProducer.Close()
	}
	if e.redisCli != nil {
		e.redisCli.Close()
	}
	if e.mysql != nil {
		e.mysql.Close()
	}
	//if e.pg != nil {
	//	e.pg.Close()
	//}
	if e.logger != nil {
		_ = e.logger.Sync()
	}
}

func (e *Environment) RootContext() context.Context {
	return e.rootContext
}

func (e *Environment) Config() *config.Config {
	return e.config
}

func (e *Environment) Logger() *zap.Logger {
	return e.logger
}

func (e *Environment) Mysql() *mysql.Mysql {
	return e.mysql
}

func (e *Environment) RedisCli() redis.UniversalClient {
	return e.redisCli
}

func (e *Environment) KafkaProducer() producer.Producer {
	return e.kafkaProducer
}

func (e *Environment) KafkaConsumer() consumer.Consumer {
	return e.kafkaConsumer
}

func (e *Environment) TranslationRepo() *repo.TranslationRepo {
	return e.translationRepo
}

func (e *Environment) TranslationWebAPI() *webapi.TranslationWebAPI {
	return e.translationWebAPI
}

func (e *Environment) TranslationUseCase() *usecase.TranslationUseCase {
	return e.translationUseCase
}
