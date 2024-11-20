package redisutils_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestStorage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	redisMock := redisutils.NewMockRedisClient(ctrl)

	res := redis.NewStatusCmd(context.Background())
	res.SetVal("pong")

	redisMock.EXPECT().Ping(gomock.Any()).Return(res)
	redisMock.EXPECT().Close().Return(nil)

	s, err := redisutils.New(redisMock)
	require.NoError(t, err)

	err = s.Cleanup()
	assert.NoError(t, err)
}

func TestCleanup(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	redisMock := redisutils.NewMockRedisClient(ctrl)

	res := redis.NewStatusCmd(context.Background())
	res.SetVal("pong")

	redisMock.EXPECT().Ping(gomock.Any()).Return(res)
	redisMock.EXPECT().Close().Return(errors.New("error"))

	s, err := redisutils.New(redisMock)
	require.NoError(t, err)

	err = s.Cleanup()
	testutils.ErrorEqual(t, errors.New("error closing redis connection: error"), err)
}

func TestStoragePingError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	redisMock := redisutils.NewMockRedisClient(ctrl)

	res := redis.NewStatusCmd(context.Background())
	res.SetErr(errors.New("error"))

	redisMock.EXPECT().Ping(gomock.Any()).Return(res)

	_, err := redisutils.New(redisMock)
	require.Error(t, err)
}
