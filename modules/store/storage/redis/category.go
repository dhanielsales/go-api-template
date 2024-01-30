package redis

import (
	"context"
	"time"

	"github.com/google/uuid"
	r "github.com/redis/go-redis/v9"

	"github.com/dhanielsales/golang-scaffold/internal/redis"
	"github.com/dhanielsales/golang-scaffold/modules/store/entity"
)

const (
	MAX_RETRY      = 5
	CATEGORY_CACHE = "category"
)

func (c *Cache) GetCategoryInCache(ctx context.Context, categoryId uuid.UUID) *entity.Category {
	key := redis.ComposeKey(CATEGORY_CACHE, categoryId.String())

	var category entity.Category

	err := c.Redis.Client.Get(ctx, key).Scan(&category)
	if err != nil {
		return nil
	}

	return &category
}

func (c *Cache) SetCategoryInCache(ctx context.Context, category entity.Category, expiration time.Duration) error {
	key := redis.ComposeKey(CATEGORY_CACHE, category.ID.String())

	err := c.Redis.Client.Set(ctx, key, category, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) DeleteCategoryInCache(ctx context.Context, categoryId uuid.UUID) error {
	key := redis.ComposeKey(CATEGORY_CACHE, categoryId.String())

	err := c.Redis.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) DeleteAllCategoryInCache(ctx context.Context) error {
	keyPattern := redis.ComposeKey(CATEGORY_CACHE, "*")

	return redis.CallTx(ctx, c.Redis.Client, func(pipe r.Pipeliner) error {
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
