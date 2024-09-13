package redis_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/redis"

	"github.com/go-redis/redismock/v9"
	goredis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposeKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		keys   []string
		result string
	}{
		{"EmptyKeys", []string{}, ""},
		{"SingleKey", []string{"key1"}, "key1"},
		{"MultipleKeys", []string{"key1", "key2", "key3"}, "key1:key2:key3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := redis.ComposeKey(tt.keys...)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestCallTxTxSucceeded(t *testing.T) {
	t.Parallel()

	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	redisStorage, err := redis.New(db)
	require.NoError(t, err)

	mock.ExpectTxPipeline()
	mock.ExpectPing().SetVal("PONG")
	mock.ExpectTxPipelineExec()

	err = redis.CallTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
		pipe.Ping(context.Background())
		return nil
	})
	assert.NoError(t, err)
}

func TestCallTxTxFailed(t *testing.T) {
	t.Parallel()

	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	redisStorage, err := redis.New(db)
	require.NoError(t, err)

	mock.ExpectTxPipeline()
	mock.ExpectPing().SetErr(goredis.TxFailedErr)
	mock.ExpectTxPipelineExec()

	err = redis.CallTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
		pipe.Ping(context.Background())
		return nil
	})
	assert.Error(t, err)
}

func TestCallTxMaxRetriesExceeded(t *testing.T) {
	t.Parallel()
	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	redisStorage, err := redis.New(db)
	require.NoError(t, err)

	mock.ExpectTxPipeline()
	mock.ExpectPing().SetErr(errors.New("ping error"))
	mock.ExpectTxPipelineExec()
	mock.ExpectPing().SetErr(errors.New("ping error"))
	mock.ExpectTxPipelineExec()
	mock.ExpectPing().SetErr(errors.New("ping error"))
	mock.ExpectTxPipelineExec()
	mock.ExpectPing().SetErr(errors.New("ping error"))
	mock.ExpectTxPipelineExec()
	mock.ExpectPing().SetErr(errors.New("ping error"))
	mock.ExpectTxPipelineExec()

	err = redis.CallTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
		pipe.Ping(context.Background())
		return nil
	})

	assert.Equal(t, redis.MAX_RETRIES_ERR, err)
}
