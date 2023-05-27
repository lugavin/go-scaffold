package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Logger `yaml:"logger"`
		//PG   `yaml:"postgres"`
		Mysql `yaml:"mysql"`
		Redis `yaml:"redis"`
		//RMQ  `yaml:"rabbitmq"`
		Kafka       `yaml:"kafka"`
		KafkaTopics `yaml:"kafka_topics"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"    env-required:"true" `
		Version string `yaml:"version" env:"APP_VERSION" env-required:"true"`
	}

	// HTTP -.
	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT" env-required:"true"`
	}

	// Logger -.
	Logger struct {
		Dev   bool     `yaml:"dev"   env:"LOGGER_DEV"`
		Level string   `yaml:"level" env:"LOGGER_LEVEL" env-required:"true"`
		Paths []string `yaml:"paths" env:"LOGGER_PATHS" env-required:"true"`
	}

	// PG -.
	PG struct {
		URL     string `yaml:"url"      env:"PG_URL"      env-required:"true"`
		PoolMax int    `yaml:"pool_max" env:"PG_POOL_MAX" env-required:"true"`
	}

	// Mysql -.
	Mysql struct {
		URL     string `yaml:"url"      env:"MS_URL"      env-required:"true"`
		PoolMax int    `yaml:"pool_max" env:"MS_POOL_MAX" env-required:"true"`
	}

	// Redis -.
	Redis struct {
		Addrs    []string `yaml:"addrs"     env:"REDIS_ADDRS"     env-required:"true"`
		PoolSize int      `yaml:"pool_size" env:"REDIS_POOL_SIZE" env-required:"true"`
	}

	// RMQ -.
	RMQ struct {
		URL            string `yaml:"url"             env:"RMQ_URL"             env-required:"true"`
		ServerExchange string `yaml:"server_exchange" env:"RMQ_SERVER_EXCHANGE" env-required:"true"`
		ClientExchange string `yaml:"client_exchange" env:"RMQ_CLIENT_EXCHANGE" env-required:"true"`
	}

	// Kafka -.
	Kafka struct {
		Brokers []string `yaml:"brokers" env:"KAFKA_BROKERS" env-required:"true"`
	}

	// KafkaTopics -.
	KafkaTopics struct {
		FooBarTopic KafkaTopic `yaml:"foo_bar_topic"`
	}

	KafkaTopic struct {
		TopicName  string `yaml:"name"`
		Partitions int    `yaml:"partitions"`
		Replicas   int    `yaml:"replicas"`
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
