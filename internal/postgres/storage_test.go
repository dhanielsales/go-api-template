package postgres_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorageBootstrapSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.Bootstrap(db)

	mock.ExpectPing()

	assert.NotNil(t, s.Client)
	assert.NoError(t, s.Client.Ping())
}

func TestStorageCleanupSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.Bootstrap(db)

	mock.ExpectClose()

	err = s.Cleanup()
	assert.NoError(t, err)
}

func TestStorageCleanupError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.Bootstrap(db)

	err = errors.New("Error on close")
	mock.ExpectClose().WillReturnError(err)

	err = s.Cleanup()
	assert.Error(t, err)
	assert.EqualError(t, err, "Error on close")
}
