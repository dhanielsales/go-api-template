package error

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
	return "unknown"
}
