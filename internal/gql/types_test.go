package gql_test

import (
	"testing"

	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := gql.NewID(42)
	assert.Equal(t, gql.ID("42"), *id)

	idStr := gql.NewID("123")
	assert.Equal(t, gql.ID("123"), *idStr)

	idZero := gql.NewID(0)
	assert.Equal(t, gql.ID(""), *idZero)
}

func TestToID(t *testing.T) {
	id := gql.ToID(42)
	assert.Equal(t, gql.ID("42"), id)

	idStr := gql.ToID("123")
	assert.Equal(t, gql.ID("123"), idStr)

	idZero := gql.ToID(0)
	assert.Equal(t, gql.ID(""), idZero)
}
