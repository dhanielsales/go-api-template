package messaging

import (
	"encoding/json"
	"errors"
	"time"
)

type Meta map[string]string

func (m Meta) ToJSON() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.New("messeging: failed to marshal meta data to JSON: " + err.Error())
	}

	return b, nil
}

func (m *Meta) FromJSON(b []byte) error {
	err := json.Unmarshal(b, m)
	if err != nil {
		return errors.New("messeging: failed to unmarshal meta data from JSON: " + err.Error())
	}

	return nil
}

type ErrorPayload struct {
	PreviousTopic        string     `json:"previous_topic"`
	OriginalOccurrenceAt time.Time  `json:"original_occurrence_at"`
	OriginalOffset       int64      `json:"original_offset"`
	Occourrences         int        `json:"occourrences"`
	IsResolved           bool       `json:"is_resolved"`
	ResolvedAt           *time.Time `json:"resolved_at"`
	StackTrace           string     `json:"stack_trace"`
	Error                string     `json:"error"`
}

func (e ErrorPayload) ToJSON() ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, errors.New("messeging: failed to marshal error payload to JSON: " + err.Error())
	}

	return b, nil
}

func (e *ErrorPayload) FromJSON(b []byte) error {
	err := json.Unmarshal(b, e)
	if err != nil {
		return errors.New("messeging: failed to unmarshal error payload from JSON: " + err.Error())
	}

	return nil
}
