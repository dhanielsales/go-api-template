package gql_test

import (
	"testing"

	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/stretchr/testify/assert"
)

func TestResponse_To(t *testing.T) {
	response := &gql.Response{
		Data: ExampleStruct{
			ID:   "42",
			Name: "John",
			Age:  25,
			Flag: true,
		},
	}

	var target ExampleStruct
	err := response.To(&target)
	assert.NoError(t, err)
	assert.Equal(t, ExampleStruct{ID: "42", Name: "John", Age: 25, Flag: true}, target)
}

func TestResponse_HasErrors(t *testing.T) {
	response := &gql.Response{
		Errors: []gqlerrors.FormattedError{{Message: "Error 1"}, {Message: "Error 2"}},
	}

	assert.True(t, response.HasErrors())

	response.Errors = nil
	assert.False(t, response.HasErrors())
}

func TestResponse_Err(t *testing.T) {
	response := &gql.Response{
		Errors: []gqlerrors.FormattedError{{Message: "Error 1"}, {Message: "Error 2"}},
	}

	err := response.Err()
	assert.Error(t, err)
	assert.EqualError(t, response.Err(), err.Error())

	response.Errors = nil
	assert.NoError(t, response.Err())
}
