package time

import "time"

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
