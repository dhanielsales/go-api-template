package testutils

import (
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
