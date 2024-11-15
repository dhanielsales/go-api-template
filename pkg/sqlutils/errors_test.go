package sqlutils_test

import (
	"testing"

	"github.com/lib/pq"

	"github.com/stretchr/testify/assert"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

func TestIsUniqueViolation(t *testing.T) {
	t.Parallel()
	mockErr := &pq.Error{Message: "duplicate key value violates unique constraint"}
	assert.True(t, sqlutils.IsUniqueViolation(mockErr))
}

func TestIsUniqueViolationByField(t *testing.T) {
	t.Parallel()
	mockErr := &pq.Error{Message: "duplicate key value violates unique constraint on field"}
	assert.True(t, sqlutils.IsUniqueViolationByField(mockErr, "field"))
}
