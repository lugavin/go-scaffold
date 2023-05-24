package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Logger -.
type Logger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type logger struct {
	log *zerolog.Logger
}

var _ Logger = (*logger)(nil)

var loggerLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
}

// New -.
func New(level string) *logger {
	globalLevel := zerolog.InfoLevel

	if l, exist := loggerLevelMap[strings.ToLower(level)]; exist {
		globalLevel = l
	}

	zerolog.SetGlobalLevel(globalLevel)

	skipFrameCount := 3
	log := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &logger{
		log: &log,
	}
}

// Debug -.
func (l *logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *logger) Info(message string, args ...interface{}) {
	l.logf(message, args...)
}

// Warn -.
func (l *logger) Warn(message string, args ...interface{}) {
	l.logf(message, args...)
}

// Error -.
func (l *logger) Error(message interface{}, args ...interface{}) {
	if l.log.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("error", message, args...)
}

// Fatal -.
func (l *logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *logger) logf(message string, args ...interface{}) {
	if len(args) == 0 {
		l.log.Info().Msg(message)
	} else {
		l.log.Info().Msgf(message, args...)
	}
}

func (l *logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.logf(msg.Error(), args...)
	case string:
		l.logf(msg, args...)
	default:
		l.logf(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
