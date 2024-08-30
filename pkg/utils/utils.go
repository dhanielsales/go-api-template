package utils

import (
	"reflect"
	"time"
)

// ToPtr returns a pointer to the given value
func ToPtr[Arg any](v Arg) *Arg {
	return &v
}

// IsPtr checks if the given value is a pointer
func IsPtr(v any) bool {
	return reflect.ValueOf(v).Type().Kind() == reflect.Pointer
}

// TimeNow returns the current time in milliseconds
func TimeNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

const (
	DAY   = 24 * time.Hour
	WEEK  = 7 * DAY
	MONTH = 30 * DAY
)
