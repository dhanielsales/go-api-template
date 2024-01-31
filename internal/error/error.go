package error

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/google/uuid"
)

type Map map[string]any

type AppError struct {
	Name        ErrorName  `json:"name"`
	Level       ErrorLevel `json:"-"`
	Description string     `json:"description,omitempty"`
	Id          string     `json:"logId,omitempty"`
	Err         error      `json:"-"`
	Details     Map        `json:"details,omitempty"`
	stack       []byte     `json:"-"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("App error %v: err %v - %v", e.Name.String(), e.Err, e.Description)
}

func (e *AppError) Merge(err error) {
	e.Err = errors.Join(e.Err, err)
}

func (e *AppError) AddDetail(key string, value any) {
	if e.Details == nil {
		e.Details = make(Map)
	}
	e.Details[key] = value
}

func (e *AppError) Stack() string {
	return string(e.stack)
}

func Is(err error) (*AppError, bool) {
	appError, ok := err.(*AppError)
	return appError, ok
}

func New(err error, name ErrorName, description string) *AppError {
	return &AppError{
		Name:        name,
		Id:          uuid.New().String(),
		Level:       name.Level(),
		Description: description,
		Err:         err,
		stack:       debug.Stack(),
	}
}
