package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func ComposeKey(keys ...string) string {
	var key string

	for _, k := range keys {
		key += k + ":"
	}

	return key
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

		if err != redis.TxFailedErr {
			return err
		}
	}

	return MAX_RETRIES_ERR
}
