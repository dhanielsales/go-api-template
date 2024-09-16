package apperror

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

type Map map[string]any

type AppError struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Details     any        `json:"details,omitempty"`
	Level       ErrorLevel `json:"-"`
	err         error      `json:"-"`
	status      int        `json:"-"`
	stack       []byte     `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error %v: err %v - %v", e.Name, e.err, e.Description)
}

func (e *AppError) Merge(err error) {
	e.err = errors.Join(e.err, err)
}

func (e *AppError) Stack() string {
	return string(e.stack)
}

func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) StatusCode() int {
	if e.status == 0 {
		e.Level = Error
		return http.StatusInternalServerError
	}

	return e.status
}

func New(description string) *AppError {
	return &AppError{
		Name:        http.StatusText(http.StatusInternalServerError),
		status:      http.StatusInternalServerError,
		Level:       Error,
		Description: description,
		stack:       debug.Stack(),
	}
}

func FromError(err error) *AppError {
	return &AppError{
		Name:        http.StatusText(http.StatusInternalServerError),
		status:      http.StatusInternalServerError,
		Level:       Error,
		Description: err.Error(),
		err:         err,
		stack:       debug.Stack(),
	}
}

func (e *AppError) WithDetails(value any) *AppError {
	e.Details = value

	return e
}

func (e *AppError) WithStatusCode(statusCode int) *AppError {
	e.status = statusCode
	e.Name = http.StatusText(statusCode)
	e.Level = ErrorLevelFromStatus(statusCode)

	return e
}

func (e *AppError) WithDescription(description string) *AppError {
	e.Description = description

	return e
}
