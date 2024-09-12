package postgres_test

import (
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorageBootstrapSuccess(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.New(db)

	mock.ExpectPing()

	assert.NotNil(t, s.Client)
	assert.NoError(t, s.Client.Ping())
}

func TestStorageCleanupSuccess(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.New(db)

	mock.ExpectClose()

	err = s.Cleanup()
	assert.NoError(t, err)
}

func TestStorageCleanupError(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	require.NoError(t, err)

	s := postgres.New(db)
	mock.ExpectClose().WillReturnError(errors.New("Error on close"))

	err = s.Cleanup()
	assert.Error(t, err)
	assert.EqualError(t, err, "Error on close")
}
