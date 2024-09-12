package postgres_test

import (
	"testing"

	"github.com/lib/pq"

	"github.com/stretchr/testify/assert"

	"github.com/dhanielsales/go-api-template/pkg/postgres"
)

func TestIsUniqueViolation(t *testing.T) {
	t.Parallel()
	mockErr := &pq.Error{Message: "duplicate key value violates unique constraint"}
	assert.True(t, postgres.IsUniqueViolation(mockErr))
}

func TestIsUniqueViolationByField(t *testing.T) {
	t.Parallel()
	mockErr := &pq.Error{Message: "duplicate key value violates unique constraint on field"}
	assert.True(t, postgres.IsUniqueViolationByField(mockErr, "field"))
}
