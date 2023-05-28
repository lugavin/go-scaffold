package env

import (
	rds "github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/pkg/logger"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase/repo"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase/webapi"
	"github.com/lugavin/go-scaffold/pkg/kafka/consumer"
	"github.com/lugavin/go-scaffold/pkg/kafka/producer"
	"github.com/lugavin/go-scaffold/pkg/mysql"
	"github.com/lugavin/go-scaffold/pkg/redis"
)

func provideLogger(c *config.Config) *zap.Logger {
	return logger.New(c.Logger)
}

//func providePostgres(c *config.Config) (*postgres.Postgres, error) {
//	return postgres.New(c.PG.URL, postgres.MaxPoolSize(c.PG.PoolMax))
//}

func provideMysql(c *config.Config) (*mysql.Mysql, error) {
	return mysql.New(c.Mysql.URL, mysql.MaxIdleConns(c.Mysql.PoolMax))
}

func provideRedis(c *config.Config) rds.UniversalClient {
	return redis.New(c.Redis.Addrs, redis.PoolSize(c.Redis.PoolSize))
}

//func provideRmqServer(l *zap.Logger, c *config.Config, translationUseCase *usecase.TranslationUseCase) (*server.Server, error) {
//	return server.New(c.RMQ.URL, c.RMQ.ServerExchange, amqprpc.NewRouter(translationUseCase), l)
//}

func provideKafkaProducer(l *zap.Logger, c *config.Config) producer.Producer {
	return producer.New(l, c.Kafka.Brokers)
}

func provideKafkaConsumer(l *zap.Logger, c *config.Config) consumer.Consumer {
	return consumer.New(l, c.Kafka.Brokers, c.App.Name)
}

func provideTranslationRepo(ms *mysql.Mysql) *repo.TranslationRepo {
	return repo.NewTranslationRepo(ms)
}

func provideTranslationWebAPI(_ *config.Config) *webapi.TranslationWebAPI {
	return webapi.NewTranslationWebAPI()
}

func provideTranslationUseCase(r *repo.TranslationRepo, w *webapi.TranslationWebAPI) *usecase.TranslationUseCase {
	return usecase.NewTranslationUseCase(r, w)
}

func provideAuthTokenRepo(ms *mysql.Mysql) *repo.AuthTokenRepo {
	return repo.NewAuthTokenRepo(ms)
}

func provideAuthTokenUseCase(r *repo.AuthTokenRepo, c *config.Config) (*usecase.AuthTokenUseCase, error) {
	return usecase.NewAuthTokenUseCase(r, c.JWT)
}
