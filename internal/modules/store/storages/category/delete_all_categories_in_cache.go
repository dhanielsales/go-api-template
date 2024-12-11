package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"

	"github.com/redis/go-redis/v9"
)

func (c *CategoryRepository) DeleteAllCategoriesInCache(ctx context.Context) error {
	keyPattern := redisutils.ComposeKey(CATEGORY_CACHE, "*")

	return redisutils.WithTx(ctx, c.Redis.Client, func(pipe redis.Pipeliner) error {
		var cursor uint64
		for {
			var err error
			var keys []string
			keys, cursor, err = c.Redis.Client.Scan(ctx, cursor, keyPattern, 10).Result()
			if err != nil {
				return err
			}

			if len(keys) > 0 {
				if err := pipe.Del(ctx, keys...).Err(); err != nil {
					return err
				}
			}

			if cursor == 0 {
				break
			}
		}

		return nil
	})
}
