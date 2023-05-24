package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		//PG   `yaml:"postgres"`
		Mysql `yaml:"mysql"`
		Redis `yaml:"redis"`
		//RMQ  `yaml:"rabbitmq"`
		Kafka       `yaml:"kafka"`
		KafkaTopics `yaml:"kafka_topics"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	// Mysql -.
	Mysql struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"MS_POOL_MAX"`
		URL     string `env-required:"true"                 env:"MS_URL"`
	}

	// Redis -.
	Redis struct {
		Addrs    []string `env-required:"true"  yaml:"addrs" env:"REDIS_ADDRS"`
		PoolSize int      `env-required:"true"  yaml:"pool_size" env:"REDIS_POOL_SIZE"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
		URL            string `env-required:"true"                            env:"RMQ_URL"`
	}

	// Kafka -.
	Kafka struct {
		Brokers []string `env-required:"true" yaml:"brokers" env:"KAFKA_BROKERS"`
	}

	// KafkaTopics -.
	KafkaTopics struct {
		FooBarTopic KafkaTopic `env-required:"true" yaml:"foo_bar_topic"`
	}

	KafkaTopic struct {
		TopicName  string `env-required:"true" yaml:"name"`
		Partitions int    `env-required:"true" yaml:"partitions"`
		Replicas   int    `env-required:"true" yaml:"replicas"`
	}
)

// NewConfig returns app config.
func NewConfig(cfgPath string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
