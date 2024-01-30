package gql_test

import (
	"context"
	"testing"

	"github.com/dhanielsales/golang-scaffold/internal/gql"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Buffer(t *testing.T) {
	req := gql.NewRequest(context.Background(), "query", map[string]any{"var": "value"})
	buffer, err := req.Buffer()
	assert.NoError(t, err)
	assert.NotNil(t, buffer)
}

func TestRequest_BufferWithError(t *testing.T) {
	invalidVariables := map[string]interface{}{
		"foo": make(chan int),
	}

	req := gql.NewRequest(context.Background(), "query", invalidVariables)
	buffer, err := req.Buffer()
	assert.Error(t, err)
	assert.Nil(t, buffer)
}

func TestSanitize(t *testing.T) {
	input := "   query    with\n\n  \t whitespaces    "
	expected := "query with whitespaces"
	result := gql.Sanitize(input)
	assert.Equal(t, expected, result)
}
