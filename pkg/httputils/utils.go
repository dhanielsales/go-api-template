package httputils

import "strconv"

func Int64ToByte(i int64) []byte {
	return []byte(strconv.FormatInt(i, 10))
}
