package utils

import (
	"reflect"
	"time"
)

func IsPointer(v any) bool {
	return reflect.ValueOf(v).Type().Kind() == reflect.Pointer
}

func TimeNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

const (
	DAY   = 24 * time.Hour
	WEEK  = 7 * DAY
	MONTH = 30 * DAY
)
