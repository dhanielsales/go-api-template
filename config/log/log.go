package log

import (
	"fmt"
	"strings"

	"github.com/dhanielsales/golang-scaffold/internal/error"
)

type Params struct {
	Message string
	Error   *error.AppError
	Meta    map[string]any
}

type Logger interface {
	Info(params Params)
	Warn(params Params)
	Error(params Params)
}

func New(prefix string) Logger {
	// TODO Adicionar logger de observabilidade aqui
	p := fmt.Sprintf("[%s] ", strings.ToUpper(prefix))
	return newStdoutLogger(p)
}
