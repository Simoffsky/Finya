package log

import (
	"fmt"
	"strings"
	"time"
)

type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}

type LoggerLevel int

const (
	Debug LoggerLevel = iota
	Info
	Warning
	Error
)

func LevelFromString(level string) LoggerLevel {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return Debug
	case "INFO":
		return Info
	case "WARNING":
		return Warning
	case "ERROR":
		return Error
	default:
		fmt.Printf("[LOG]Unknown log level \"%s\", setting to DEBUG level\n", level)
		return Debug
	}
}

type defaultLogger struct {
	level  LoggerLevel
	prefix string
}

func NewDefaultLogger(level LoggerLevel) *defaultLogger {
	return &defaultLogger{
		level: level,
	}
}

func (l *defaultLogger) WithPrefix(prefix string) *defaultLogger {
	return &defaultLogger{
		level:  l.level,
		prefix: l.prefix + prefix,
	}
}

func (l defaultLogger) WithTimePrefix() *defaultLogger {
	return l.WithPrefix(time.Now().Local().Format("[2006-01-02 15:04:05]" + " "))
}

func (l *defaultLogger) SetLevel(level LoggerLevel) {
	l.level = level
}

func (l *defaultLogger) Info(msg string) {
	if l.level <= Info {
		println(l.prefix + "INFO: " + msg)
	}
}

func (l *defaultLogger) Debug(msg string) {
	if l.level <= Debug {
		println(l.prefix + "DEBUG: " + msg)
	}

}

func (l *defaultLogger) Warning(msg string) {
	if l.level <= Warning {
		println(l.prefix + "WARNING: " + msg)
	}
}

func (l *defaultLogger) Error(msg string) {
	if l.level <= Error {
		println(l.prefix + "ERROR: " + msg)
	}
}
