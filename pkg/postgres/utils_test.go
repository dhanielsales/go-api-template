package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		page     string
		perPage  string
		expected postgres.PaginationResult
	}{
		{"ValidInput", "2", "20", postgres.PaginationResult{Limit: 20, Offset: 20}},
		{"InvalidPage", "invalid", "20", postgres.PaginationResult{Limit: 20, Offset: 0}},
		{"InvalidPerPage", "2", "invalid", postgres.PaginationResult{Limit: 10, Offset: 10}},
		{"ZeroPage", "0", "20", postgres.PaginationResult{Limit: 20, Offset: 0}},
		{"ZeroPerPage", "2", "0", postgres.PaginationResult{Limit: 10, Offset: 10}},
		{"LargePerPage", "2", "150", postgres.PaginationResult{Limit: 100, Offset: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := postgres.Pagination(tt.page, tt.perPage)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSorting(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		field     string
		direction string
		expected  string
	}{
		{"ValidInput", "name", "ASC", "name ASC"},
		{"EmptyField", "", "DESC", "created_at DESC"},
		{"EmptyDirection", "name", "", "name DESC"},
		{"InvalidDirection", "name", "invalid", "name DESC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := postgres.Sorting(tt.field, tt.direction)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCallTxCommitSuccess(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectCommit()

	mockResult := 1
	result, err := postgres.CallTx(context.Background(), db, func(tx *sql.Tx) (*int, error) {
		return &mockResult, nil
	})

	require.NoError(t, err)
	assert.Equal(t, &mockResult, result)
}

func TestCallTxRollbackOnError(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback()

	mockErr := errors.New("mock-error")
	result, err := postgres.CallTx(context.Background(), db, func(tx *sql.Tx) (*int, error) {
		return nil, mockErr
	})

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestCallTxBeginWithError(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectBegin().WillReturnError(errors.New("mock-error"))
	mock.ExpectCommit()

	result, err := postgres.CallTx(context.Background(), db, func(tx *sql.Tx) (*int, error) {
		n := 0
		return &n, nil
	})

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestCallTxRollbackWithError(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback().WillReturnError(errors.New("mock-error"))

	result, err := postgres.CallTx(context.Background(), db, func(tx *sql.Tx) (*int, error) {
		n := 0
		return &n, nil
	})

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestCallTxWithErrorAndRollbackWithError(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback().WillReturnError(errors.New("mock-error"))

	mockErr := errors.New("mock-error")
	result, err := postgres.CallTx(context.Background(), db, func(tx *sql.Tx) (*int, error) {
		return nil, mockErr
	})

	require.Error(t, err)
	assert.Nil(t, result)
}
