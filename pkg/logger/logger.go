package logger

type Logger interface {
	Printf(format string, args ...interface{})
}
