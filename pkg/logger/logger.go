package logger

import (
	"log/slog"
	"sync"
)

type Field struct {
	Key   string
	Value any
}

type FieldOption func(c *Field)

type Logger interface {
	Info(message string, fields ...FieldOption)
	Warn(message string, fields ...FieldOption)
	Error(message string, fields ...FieldOption)
	Debug(message string, fields ...FieldOption)
}

var (
	once     sync.Once
	instance Logger
)

// GetInstance returns the singleton instance of the logger
func GetInstance() Logger {
	if instance == nil {
		once.Do(func() {
			instance = NewStdoutLogger(slog.LevelDebug)
		})
	}

	return instance
}

func Info(message string, fields ...FieldOption) {
	GetInstance().Info(message, fields...)
}

func Warn(message string, fields ...FieldOption) {
	GetInstance().Warn(message, fields...)
}

func Error(message string, fields ...FieldOption) {
	GetInstance().Error(message, fields...)
}

func Debug(message string, fields ...FieldOption) {
	GetInstance().Debug(message, fields...)
}
