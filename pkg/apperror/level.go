package apperror

// ErrorLevel represents the severity level of an error.
type ErrorLevel uint8

const (
	Info  ErrorLevel = iota // Info level indicates informational messages.
	Warn                    // Warn level indicates warning messages.
	Error                   // Error level indicates error messages.
)

// String returns a string representation of the ErrorLevel.
func (l ErrorLevel) String() string {
	switch l {
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	}

	return "unknown" // Default if the error level is unknown.
}

// ErrorLevelFromStatus maps HTTP status codes to error levels.
// It defaults to `Error` level if no match is found.
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
