package redis

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

var MAX_RETRIES_ERR = errors.New("Max retries reached")

// CallTx executes a Redis transaction with retry logic, allowing for concurrent-safe operations.
func CallTx(ctx context.Context, client *redis.Client, fn func(pipe redis.Pipeliner) error) error {
	for retries := MAX_RETRIES; retries > 0; retries-- {
		err := client.Watch(ctx, func(tx *redis.Tx) error {
			_, err := tx.TxPipelined(ctx, fn)
			return err
		})

		if err == nil {
			return nil
		}

		if err == redis.TxFailedErr {
			return err
		}
	}

	return MAX_RETRIES_ERR
}
