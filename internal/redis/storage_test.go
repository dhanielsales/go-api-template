package redis_test

import (
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dhanielsales/golang-scaffold/internal/redis"
)

func TestStorage(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	s, err := redis.New(db)
	require.NoError(t, err)

	err = s.Cleanup()
	assert.NoError(t, err)
}

func TestStoragePingError(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetErr(errors.New("ping error"))

	_, err := redis.New(db)
	require.Error(t, err)
}
