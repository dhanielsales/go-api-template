package time

import "time"

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

const (
	DAY   = 24 * time.Hour
	WEEK  = 7 * DAY
	MONTH = 30 * DAY
)
