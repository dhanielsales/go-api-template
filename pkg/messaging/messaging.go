package messaging

import "time"

type Options struct {
	meta map[string]string
}

type Emitter interface {
	Emit(event Event, data []byte, options Options) error
}

type Listener interface {
	Listen(event Event, handler Handler) error
}

func GetWaitingTime(occurrences int) time.Duration {
	return time.Duration(INITIAL_RETRY_DELAY * occurrences)
}
