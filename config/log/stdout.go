package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type stdoutLogger struct {
	logger *log.Logger
}

type logPayload struct {
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Stack   string         `json:"stack"`
	LogId   string         `json:"logId"`
	Level   string         `json:"level"`
	Meta    map[string]any `json:"meta"`
}

func newStdoutLogger(prefix string) Logger {
	return &stdoutLogger{
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

func formatPayload(params Params) string {
	p := logPayload{}

	if params.Error != nil {
		p.Error = params.Error.Error()
		p.Level = params.Error.Name.Level().String()
		p.LogId = params.Error.Id
	} else {
		p.Level = "info"
		p.LogId = uuid.New().String()
	}

	p.Message = params.Message
	p.Meta = params.Meta

	f, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(f)
}

func (l *stdoutLogger) Info(params Params) {
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Stack())
}

func (l *stdoutLogger) Warn(params Params) {
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Stack())
}

func (l *stdoutLogger) Error(params Params) {
	formated := formatPayload(params)
	l.logger.Println(formated)
	fmt.Println(params.Error.Stack())
}
