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

type DefaultLogger struct {
	level  LoggerLevel
	prefix string
}

func NewDefaultLogger(level LoggerLevel) *DefaultLogger {
	return &DefaultLogger{
		level: level,
	}
}

func (l DefaultLogger) WithPrefix(prefix string) *DefaultLogger {
	return &DefaultLogger{
		level:  l.level,
		prefix: prefix,
	}
}

func (l DefaultLogger) WithTimePrefix() *DefaultLogger {
	return l.WithPrefix(time.Now().Local().Format("[2006-01-02 15:04:05]" + " "))
}

func (l *DefaultLogger) SetLevel(level LoggerLevel) {
	l.level = level
}

func (l *DefaultLogger) Info(msg string) {
	if l.level <= Info {
		println(l.prefix + "INFO: " + msg)
	}
}

func (l *DefaultLogger) Debug(msg string) {
	if l.level <= Debug {
		println(l.prefix + "DEBUG: " + msg)
	}

}

func (l *DefaultLogger) Warning(msg string) {
	if l.level <= Warning {
		println(l.prefix + "WARNING: " + msg)
	}
}

func (l *DefaultLogger) Error(msg string) {
	if l.level <= Error {
		println(l.prefix + "ERROR: " + msg)
	}
}
