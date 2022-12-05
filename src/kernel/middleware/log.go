package middleware

import "log"

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

var defaultLogger = DefaultLogger{level: LogLevelError, l: log.Default()}

// GetDefaultLogger 返回 中间件的 Default Logger
func GetDefaultLogger() DefaultLogger {
	return defaultLogger
}

type LogLevel int32

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelNone
)

type DefaultLogger struct {
	level LogLevel
	l     *log.Logger
}

func (d DefaultLogger) Debug(msg string) {
	if d.level <= LogLevelDebug {
		d.l.Printf("[DEBUG] %s", msg)
	}
}

func (d DefaultLogger) Info(msg string) {
	if d.level <= LogLevelInfo {
		d.l.Printf("[INFO] %s", msg)
	}
}

func (d DefaultLogger) Warning(msg string) {
	if d.level <= LogLevelWarning {
		d.l.Printf("[WARNING] %s", msg)
	}
}

func (d DefaultLogger) Error(msg string) {
	if d.level <= LogLevelError {
		d.l.Printf("[ERROR] %s", msg)
	}
}
