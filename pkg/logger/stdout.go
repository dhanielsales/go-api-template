package logger

import (
	"log/slog"
	"os"
)

type StdoutLogger struct {
	logger *slog.Logger
}

func NewStdoutLogger(level slog.Leveler) Logger {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: level},
	))

	return &StdoutLogger{logger: logger}
}

func (l *StdoutLogger) Info(message string, fields ...FieldOption) {
	formatedFields := formatFields(fields)
	l.logger.Info(message, formatedFields...)
}

func (l *StdoutLogger) Warn(message string, fields ...FieldOption) {
	formatedFields := formatFields(fields)
	l.logger.Warn(message, formatedFields...)
}

func (l *StdoutLogger) Error(message string, fields ...FieldOption) {
	formatedFields := formatFields(fields)
	l.logger.Error(message, formatedFields...)
}

func (l *StdoutLogger) Debug(message string, fields ...FieldOption) {
	formatedFields := formatFields(fields)
	l.logger.Debug(message, formatedFields...)
}
