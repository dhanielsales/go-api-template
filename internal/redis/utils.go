package redis

import (
	"context"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ComposeKey(keys ...string) string {
	return strings.Join(keys, ":")
}

var (
	MAX_RETRIES     = 5
	MAX_RETRIES_ERR = errors.New("Max retries reached")
)

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
