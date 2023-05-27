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

func New(logCfg config.Logger) *zap.Logger {
	once.Do(func() {
		var (
			zapCfg zap.Config
			err    error
		)
		// 创建一个配置实例
		if logCfg.Development {
			zapCfg = zap.NewDevelopmentConfig()
		} else {
			zapCfg = zap.NewProductionConfig()
		}

		// 设置日志级别
		level := zap.NewAtomicLevel()
		if err = level.UnmarshalText([]byte(logCfg.Level)); err != nil {
			log.Fatalf("Failed to parse log level: %s", err)
		}
		zapCfg.Level = level

		// 设置日志输出路径
		zapCfg.OutputPaths = logCfg.OutputPaths

		// 创建一个Logger实例
		if logger, err = zapCfg.Build(); err != nil {
			log.Fatalf("logger.New error: %s", err)
		}

		// 将标准库日志重定向到 ZapLogger
		zap.RedirectStdLog(logger)
	})
	return logger
}
