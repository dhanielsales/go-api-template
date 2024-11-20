package redisutils

import (
	"context"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
)

// ComposeKey concatenates multiple strings into a single Redis key, using ":" as a separator.
func ComposeKey(keys ...string) string {
	return strings.Join(keys, ":")
}

// Maximum number of retries for a transaction.
const MAX_RETRIES = 5

var ErrMaxRetryReachedOut = errors.New("max retries reached out")

// WithTx executes a Redis transaction with retry logic, allowing for concurrent-safe operations.
func WithTx(ctx context.Context, client RedisClient, fn func(pipe redis.Pipeliner) error) error {
	for retries := MAX_RETRIES; retries > 0; retries-- {
		err := client.Watch(ctx, func(tx *redis.Tx) error {
			_, err := tx.TxPipelined(ctx, fn)
			return err
		})

		if err == nil {
			return nil
		}

		if errors.Is(err, redis.TxFailedErr) {
			return err
		}
	}

	return ErrMaxRetryReachedOut
}
