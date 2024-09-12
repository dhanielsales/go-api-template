package apperror

type ErrorLevel uint8

const (
	Info ErrorLevel = iota
	Warn
	Error
)

func (l ErrorLevel) String() string {
	switch l {
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	}

	return unknown
}

func ErrorLevelFromStatus(status int) ErrorLevel {
	if status > 100 && status < 300 {
		return Info
	}

	if status >= 400 && status < 500 {
		return Warn
	}

	if status >= 500 {
		return Error
	}

	return Error
}
