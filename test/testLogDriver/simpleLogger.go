package testLogDriver

import (
	"fmt"
	"log"
)

const (
	DebugLevel   int8 = 0
	InfoLevel    int8 = 1
	WarningLevel int8 = 2
	ErrorLevel   int8 = 3
	PanicLevel   int8 = 4
	FatalLevel   int8 = 5
)

type SimpleLogger struct{}

func (l *SimpleLogger) Debug(msg string, v ...interface{}) {
	log.Printf("[DEBUG] "+msg, v...)
}

func (l *SimpleLogger) Info(msg string, v ...interface{}) {
	log.Printf("[INFO] "+msg, v...)
}

func (l *SimpleLogger) Warn(msg string, v ...interface{}) {
	log.Printf("[WARNING] "+msg, v...)
}

func (l *SimpleLogger) Error(msg string, v ...interface{}) {
	log.Printf("[ERROR] "+msg, v...)
}

func (l *SimpleLogger) Panic(msg string, v ...interface{}) {
	log.Panicf("[PANIC] "+msg, v...)
}

func (l *SimpleLogger) Fatal(msg string, v ...interface{}) {
	log.Fatalf("[FATAL] "+msg, v...)
}

func (l *SimpleLogger) DebugF(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *SimpleLogger) InfoF(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *SimpleLogger) WarnF(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *SimpleLogger) ErrorF(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *SimpleLogger) PanicF(format string, args ...interface{}) {
	l.Panic(fmt.Sprintf(format, args...))
}

func (l *SimpleLogger) FatalF(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
