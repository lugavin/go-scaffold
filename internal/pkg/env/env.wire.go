//go:build wireinject
// +build wireinject

package env

import (
	"context"

	"github.com/google/wire"

	"github.com/lugavin/go-scaffold/config"
)

func initEnvironment(ctx context.Context, cfg *config.Config) (*Environment, error) {
	panic(wire.Build(
		wire.Struct(
			new(Environment),
			"logger",
			"mysql",
			"redisCli",
			"kafkaProducer",
			"kafkaConsumer",
			"translationRepo",
			"translationWebAPI",
			"translationUseCase",
			"authTokenRepo",
			"authTokenUseCase",
		),
		provideLogger,
		provideMysql,
		provideRedis,
		provideKafkaProducer,
		provideKafkaConsumer,
		provideTranslationRepo,
		provideTranslationWebAPI,
		provideTranslationUseCase,
		provideAuthTokenRepo,
		provideAuthTokenUseCase,
	))
}
