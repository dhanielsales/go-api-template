package postgres

import (
	"strings"

	"github.com/lib/pq"
)

// IsUniqueViolation check if error is a unique constraint violation error
func IsUniqueViolation(err error) bool {
	return strings.Contains(err.(*pq.Error).Message, "duplicate key value violates unique constraint")
}

// IsUniqueViolationByField check if error is a unique constraint violation error by field
func IsUniqueViolationByField(err error, field string) bool {
	return IsUniqueViolation(err) && strings.Contains(err.(*pq.Error).Message, field)
}
