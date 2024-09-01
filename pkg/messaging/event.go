package messaging

import (
	"strings"
)

// Event represents a fully qualified event name
type Event string

func (e Event) String() string {
	return string(e)
}

func (e Event) Namespace() string {
	strArr := strings.Split(e.String(), ".")
	return strArr[0]
}

func (e Event) Group() string {
	strArr := strings.Split(e.String(), ".")
	return strArr[1]
}

func (e Event) Name() string {
	strArr := strings.Split(e.String(), ".")
	return strArr[2]
}

func (e Event) Validate() error {
	if e == "" {
		return ErrInvalidEvent
	}

	strArr := strings.Split(e.String(), ".")

	if len(strArr) != 3 {
		return ErrInvalidEvent
	}

	if strArr[0] == "" || strArr[1] == "" || strArr[2] == "" {
		return ErrInvalidEvent
	}

	return nil
}
