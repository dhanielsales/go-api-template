package apperror

import "net/http"

// ErrorName represents different types of errors in the application.
type ErrorName uint8

// Constants for different error names.
const unknown = "unknown"

const (
	InternalServerError      ErrorName = iota // Represents internal server errors (500).
	UnauthorizedError                         // Represents unauthorized errors (401).
	NotFoundError                             // Represents not found errors (404).
	BadRequestError                           // Represents bad request errors (400).
	UnprocessableEntityError                  // Represents unprocessable entity errors (422).
	ForbiddenRequestError                     // Represents forbidden request errors (403).
)

// String returns a string representation of the ErrorName.
func (n ErrorName) String() string {
	switch n {
	case InternalServerError:
		return "InternalServerError"
	case UnauthorizedError:
		return "UnauthorizedError"
	case NotFoundError:
		return "NotFoundError"
	case BadRequestError:
		return "BadRequestError"
	case UnprocessableEntityError:
		return "UnprocessableEntityError"
	case ForbiddenRequestError:
		return "ForbiddenRequestError"
	}

	return unknown
}

// Status returns the HTTP status code corresponding to the ErrorName.
func (n ErrorName) Status() int {
	switch n {
	case InternalServerError:
		return http.StatusInternalServerError
	case UnauthorizedError:
		return http.StatusUnauthorized
	case NotFoundError:
		return http.StatusNotFound
	case BadRequestError:
		return http.StatusBadRequest
	case UnprocessableEntityError:
		return http.StatusUnprocessableEntity
	case ForbiddenRequestError:
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

// Level returns the error level corresponding to the ErrorName.
func (n ErrorName) Level() ErrorLevel {
	switch n {
	case InternalServerError:
		return Error
	case UnauthorizedError:
		return Warn
	case NotFoundError:
		return Info
	case BadRequestError:
		return Info
	case UnprocessableEntityError:
		return Info
	case ForbiddenRequestError:
		return Warn
	}
	return Error
}
