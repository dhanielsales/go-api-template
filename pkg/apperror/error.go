package apperror

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

// Map represents a custom type for storing key-value pairs, typically used to provide additional details about the error.
type Map map[string]any

// AppError represents an application-specific error with additional context such as description, details, and stack trace.
type AppError struct {
	Name        string     `json:"name"`                  // The name of the error.
	Description string     `json:"description,omitempty"` // A description providing more details about the error.
	Details     any        `json:"details,omitempty"`     // Additional details related to the error.
	Level       ErrorLevel `json:"-"`                     // The severity or level of the error (e.g., Error, Warning).
	err         error      `json:"-"`                     // The underlying error.
	status      int        `json:"-"`                     // The HTTP status code associated with the error.
	stack       []byte     `json:"-"`                     // A stack trace for debugging purposes.
}

// Error implements the error interface for AppError, returning a formatted error string.
func (e *AppError) Error() string {
	return fmt.Sprintf("error %v: err %v - %v", e.Name, e.err, e.Description)
}

// Merge combines the current error with another error, accumulating the details.
func (e *AppError) Merge(err error) {
	e.err = errors.Join(e.err, err)
}

// Stack returns the stack trace of the error as a string.
func (e *AppError) Stack() string {
	return string(e.stack)
}

// Unwrap returns the underlying error wrapped in AppError.
func (e *AppError) Unwrap() error {
	return e.err
}

// StatusCode returns the HTTP status code associated with the error.
// If no specific code is set, it defaults to internal server error (500).
func (e *AppError) StatusCode() int {
	if e.status == 0 {
		e.Level = Error
		return http.StatusInternalServerError
	}
	return e.status
}

// New creates a new AppError with the provided description and internal server error as the default status.
func New(description string) *AppError {
	return &AppError{
		Name:        http.StatusText(http.StatusInternalServerError),
		status:      http.StatusInternalServerError,
		Level:       Error,
		Description: description,
		stack:       debug.Stack(),
	}
}

// FromError creates a new AppError from an existing error, setting the status to internal server error.
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

// WithDetails sets additional details to the AppError.
func (e *AppError) WithDetails(value any) *AppError {
	e.Details = value
	return e
}

// WithStatusCode sets a specific HTTP status code for the AppError.
func (e *AppError) WithStatusCode(statusCode int) *AppError {
	e.status = statusCode
	e.Name = http.StatusText(statusCode)
	e.Level = ErrorLevelFromStatus(statusCode)
	return e
}

// WithDescription sets a custom description for the AppError.
func (e *AppError) WithDescription(description string) *AppError {
	e.Description = description
	return e
}
