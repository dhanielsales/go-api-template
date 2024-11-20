package redisutils_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/go-redis/redismock/v9"
	redis "github.com/redis/go-redis/v9"
	gomock "go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
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
			result := redisutils.ComposeKey(tt.keys...)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestWithTx(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		txHandler func(pipe redis.Pipeliner) error
		prepare   func() redisutils.RedisClient
		expected  error
	}{
		{
			name:      "Success",
			txHandler: func(pipe redis.Pipeliner) error { return nil },
			prepare: func() redisutils.RedisClient {
				ctrl := gomock.NewController(t)
				redisMock := redisutils.NewMockRedisClient(ctrl)

				redisMock.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(nil)

				return redisMock
			},
			expected: nil,
		},
		{
			name:      "Error on handler",
			txHandler: func(pipe redis.Pipeliner) error { return errors.New("error on handler") },
			prepare: func() redisutils.RedisClient {
				ctrl := gomock.NewController(t)
				redisMock := redisutils.NewMockRedisClient(ctrl)

				redisMock.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(errors.New("error on handler")).MaxTimes(5)

				return redisMock
			},
			expected: redisutils.ErrMaxRetryReachedOut,
		},
		{
			name:      "Error tx failed",
			txHandler: func(pipe redis.Pipeliner) error { return errors.New("error on handler") },
			prepare: func() redisutils.RedisClient {
				ctrl := gomock.NewController(t)
				redisMock := redisutils.NewMockRedisClient(ctrl)

				redisMock.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(redis.TxFailedErr)

				return redisMock
			},
			expected: redis.TxFailedErr,
		},
		{
			name:      "Error tx pipelined",
			txHandler: func(pipe redis.Pipeliner) error { return errors.New("error on handler") },
			prepare: func() redisutils.RedisClient {
				db, mock := redismock.NewClientMock()

				mock.ExpectTxPipeline()
				mock.ExpectTxPipelineExec()

				return db
			},
			expected: redisutils.ErrMaxRetryReachedOut,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			redisdb := tt.prepare()
			err := redisutils.WithTx(context.Background(), redisdb, tt.txHandler)

			testutils.ErrorEqual(t, tt.expected, err)
		})
	}
}

// func TestCallTxTxSucceeded(t *testing.T) {
// 	t.Parallel()

// 	ctrl := gomock.NewController(t)
// 	redisMock := redisutils.NewMockRedisClient(ctrl)

// 	res := redis.NewStatusCmd(context.Background())
// 	res.SetVal("pong")

// 	redisMock.EXPECT().Ping(gomock.Any()).Return(res)

// 	redisStorage, err := redisutils.New(redisMock)
// 	require.NoError(t, err)

// 	err = redisutils.WithTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
// 		pipe.Ping(context.Background())
// 		return nil
// 	})
// 	assert.NoError(t, err)
// }

// func TestCallTxTxFailed(t *testing.T) {
// 	t.Parallel()

// 	db, mock := redismock.NewClientMock()

// 	mock.ExpectPing().SetVal("PONG")

// 	redisStorage, err := redisutils.New(db)
// 	require.NoError(t, err)

// 	mock.ExpectTxPipeline()
// 	mock.ExpectPing().SetErr(goredis.TxFailedErr)
// 	mock.ExpectTxPipelineExec()

// 	err = redisutils.WithTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
// 		pipe.Ping(context.Background())
// 		return nil
// 	})
// 	assert.Error(t, err)
// }

// func TestCallTxMaxRetriesExceeded(t *testing.T) {
// 	t.Parallel()
// 	db, mock := redismock.NewClientMock()

// 	mock.ExpectPing().SetVal("PONG")

// 	redisStorage, err := redisutils.New(db)
// 	require.NoError(t, err)

// 	mock.ExpectTxPipeline()
// 	mock.ExpectPing().SetErr(errors.New("ping error"))
// 	mock.ExpectTxPipelineExec()
// 	mock.ExpectPing().SetErr(errors.New("ping error"))
// 	mock.ExpectTxPipelineExec()
// 	mock.ExpectPing().SetErr(errors.New("ping error"))
// 	mock.ExpectTxPipelineExec()
// 	mock.ExpectPing().SetErr(errors.New("ping error"))
// 	mock.ExpectTxPipelineExec()
// 	mock.ExpectPing().SetErr(errors.New("ping error"))
// 	mock.ExpectTxPipelineExec()

// 	err = redisutils.WithTx(context.Background(), redisStorage.Client, func(pipe goredis.Pipeliner) error {
// 		pipe.Ping(context.Background())
// 		return nil
// 	})

// 	assert.Equal(t, redisutils.MAX_RETRIES_ERR, err)
// }
