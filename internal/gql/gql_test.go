package gql_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dhanielsales/golang-scaffold/internal/gql"
)

type ExampleStruct struct {
	ID   gql.ID      `json:"id"`
	Name gql.String  `json:"name"`
	Age  gql.Int     `json:"age"`
	Flag gql.Boolean `json:"flag"`
}

func TestClient_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": {"id": "123", "name": "Alice", "age": 30, "flag": true}}`))
	}))
	defer server.Close()

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, ExampleStruct{ID: "123", Name: "Alice", Age: 30, Flag: true}, target)
}

func TestClient_Do_Error(t *testing.T) {
	client := gql.NewClient("invalid-url", nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.ErrorContains(t, err, "unsupported protocol scheme")
}

func TestClient_Do_ErrorTargetIsNotPointer(t *testing.T) {
	client := gql.NewClient("http://example.com", nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, target)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.ErrorIs(t, err, gql.ErrTargetIsNotPointer)
}

func TestClient_Do_ErrorBuffer(t *testing.T) {
	invalidVariables := map[string]interface{}{
		"foo": make(chan int),
	}
	client := gql.NewClient("http://example.com", nil)
	req := gql.NewRequest("query", invalidVariables)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestClient_Do_ErrorServerResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(nil))
	}))
	defer server.Close()

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestClient_Do_ErrorNewRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data": {"id": "123", "name": "Alice", "age": 30, "flag": true}}`))
	}))
	defer server.Close()

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(nil, req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestClient_Do_ErrorIoReadAll(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("Content-Length", "50")
		w.Write([]byte("Payload smaller than Content-Length"))
	}))

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestClient_Do_ErrorResponseErr(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(`{ "errors": [ {"message": "Error 1" } ] }`))
	}))
	defer server.Close()

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target ExampleStruct
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestClient_Do_ErrorResponseTo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(`{"data": {"id": "123", "name": "Alice", "age": 30, "flag": true}}`))
	}))
	defer server.Close()

	client := gql.NewClient(server.URL, nil)
	req := gql.NewRequest("query", nil)

	var target string
	response, err := client.Do(context.Background(), req, &target)
	assert.Error(t, err)
	assert.Nil(t, response)
}
