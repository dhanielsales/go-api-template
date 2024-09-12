package postgres

import (
	"strings"

	"github.com/lib/pq"
)

// IsUniqueViolation check if error is a unique constraint violation error
func IsUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return strings.Contains(pqErr.Message, "duplicate key value violates unique constraint")
	}

	return false
}

// IsUniqueViolationByField check if error is a unique constraint violation error by field
func IsUniqueViolationByField(err error, field string) bool {
	if pqErr, ok := err.(*pq.Error); ok && IsUniqueViolation(err) {
		return strings.Contains(pqErr.Message, field)
	}

	return false
}
