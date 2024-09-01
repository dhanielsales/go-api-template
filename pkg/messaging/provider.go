package messaging

import (
	"errors"
	"time"
)

var (
	ErrInvalidEvent           = errors.New("invalid event: event must be in the format '<namespace>.<group>.<event>'")
	ErrEventAlreadySubscribed = errors.New("event already subscribed")
)

const (
	FANOUT_GROUP        string = "fanout"
	META_HEADER_KEY     string = "x-private-meta"
	ERROR_HEADER_KEY    string = "x-private-error"
	RETRY_TOPIC         string = "retry-topic"
	RETRY_MAX_ATTEMPTS  int    = 5
	INITIAL_RETRY_DELAY int    = int(500 * time.Millisecond)
)

type EventPayload struct {
	Data       []byte
	Event      Event
	Meta       Meta
	ErrPayload ErrorPayload
}

type Handler func(EventPayload) error

type Provider interface {
	Subscribe(Subject Event, handler Handler) error
	SubscribeFanOut(Subject Event, handler Handler) error
	Publish(Subject Event, data []byte, meta Meta) error
	Cleanup()
	Start()
}
