package testutils

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ErrorEqual(t *testing.T, want, got error) bool {
	t.Helper()
	if want == nil {
		return assert.Equal(t, want, got)
	}

	return assert.EqualError(t, got, want.Error())
}

func BytesEqual(t *testing.T, want, got []byte) bool {
	t.Helper()

	wantStr := strings.ReplaceAll(string(want), "\n", "")
	wantStr = strings.ReplaceAll(wantStr, " ", "")
	gotStr := strings.ReplaceAll(string(got), "\n", "")
	gotStr = strings.ReplaceAll(gotStr, " ", "")

	return assert.EqualValues(t, wantStr, gotStr)
}

// Int64ToByte converts an int64 value to a byte slice.
func Int64ToByte(i int64) []byte {
	return []byte(strconv.FormatInt(i, 10))
}

// ToByte converts an struct value to a byte's slice. It will return a empty byte's slice in error
func ToByte(val any) []byte {
	b, err := json.Marshal(val)
	if err != nil {
		return []byte{}
	}

	return b
}
