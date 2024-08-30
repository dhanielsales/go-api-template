package log

import (
	"fmt"
	"strings"

	"github.com/dhanielsales/golang-scaffold/pkg/error"
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
	// TODO Add observability here
	// TODO Convert logger to singleton
	p := fmt.Sprintf("[%s] ", strings.ToUpper(prefix))
	return newStdoutLogger(p)
}
