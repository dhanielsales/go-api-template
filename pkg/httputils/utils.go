package httputils

import "strconv"

// Int64ToByte converts an int64 value to a byte slice.
func Int64ToByte(i int64) []byte {
	return []byte(strconv.FormatInt(i, 10))
}
