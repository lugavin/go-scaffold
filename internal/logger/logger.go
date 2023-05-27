package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"

	"github.com/lugavin/go-scaffold/config"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func New(c config.Logger) *zap.Logger {
	once.Do(func() {
		var (
			cfg zap.Config
			err error
		)
		// 创建一个配置实例
		if c.Dev {
			cfg = zap.NewDevelopmentConfig()
		} else {
			cfg = zap.NewProductionConfig()
		}

		// 设置日志级别
		level := zap.NewAtomicLevel()
		if err = level.UnmarshalText([]byte(c.Level)); err != nil {
			log.Fatalf("Failed to parse log level: %s", err)
		}
		cfg.Level = level

		// 设置日志输出路径
		cfg.OutputPaths = c.Paths

		// 创建一个Logger实例
		if logger, err = cfg.Build(); err != nil {
			log.Fatalf("logger.New error: %s", err)
		}

		// 将标准库日志重定向到ZapLogger
		zap.RedirectStdLog(logger)
	})
	return logger
}
