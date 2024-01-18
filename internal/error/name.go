package error

import "net/http"

type ErrorName uint8

const (
	ServerError ErrorName = iota
	UnauthorizedError
	NotFoundError
	BadRequestError
	UnprocessableEntityError
	ForbiddenRequestError
)

func (n ErrorName) String() string {
	switch n {
	case ServerError:
		return "ServerError"
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
	return "unknown"
}

func (n ErrorName) Status() int {
	switch n {
	case ServerError:
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
	return 500
}

func (n ErrorName) Level() ErrorLevel {
	switch n {
	case ServerError:
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
