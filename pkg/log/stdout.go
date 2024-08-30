package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type stdoutLogger struct {
	logger *log.Logger
}

type logPayload struct {
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Stack   string         `json:"stack"`
	Level   string         `json:"level"`
	Meta    map[string]any `json:"meta"`
}

// TODO Add Zap logger option

func newStdoutLogger(prefix string) Logger { // TODO change to default go logger
	return &stdoutLogger{
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

func formatPayload(params Params) string { // TODO change to default go logger
	p := logPayload{}

	if params.Error != nil {
		p.Error = params.Error.Error()
		p.Level = params.Error.Level.String()
	} else {
		p.Level = "info"
	}

	p.Message = params.Message
	p.Meta = params.Meta

	f, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(f)
}

func (l *stdoutLogger) Info(params Params) { // TODO change to default go logger
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Unwrap())
	fmt.Println(params.Error.Stack())
}

func (l *stdoutLogger) Warn(params Params) { // TODO change to default go logger
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Unwrap())
	fmt.Println(params.Error.Stack())
}

func (l *stdoutLogger) Error(params Params) { // TODO change to default go logger
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Unwrap())
	fmt.Println(params.Error.Stack())
}
