package redis_test

import (
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/redis"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Parallel()

	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	s, err := redis.New(db)
	require.NoError(t, err)

	err = s.Cleanup()
	assert.NoError(t, err)
}

func TestStoragePingError(t *testing.T) {
	t.Parallel()

	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetErr(errors.New("ping error"))

	_, err := redis.New(db)
	require.Error(t, err)
}
