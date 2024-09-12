package redis_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	r "github.com/redis/go-redis/v9"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/redis"
)

const (
	MAX_RETRY      = 5
	CATEGORY_CACHE = "category"
)

func (c *CacheRepository) GetCategoryInCache(ctx context.Context, categoryID uuid.UUID) *models.Category {
	key := redis.ComposeKey(CATEGORY_CACHE, categoryID.String())

	var category models.Category

	err := c.Redis.Client.Get(ctx, key).Scan(&category)
	if err != nil {
		return nil
	}

	return &category
}

func (c *CacheRepository) SetCategoryInCache(ctx context.Context, category *models.Category, expiration time.Duration) error {
	key := redis.ComposeKey(CATEGORY_CACHE, category.ID.String())

	err := c.Redis.Client.Set(ctx, key, category, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRepository) DeleteCategoryInCache(ctx context.Context, categoryID uuid.UUID) error {
	key := redis.ComposeKey(CATEGORY_CACHE, categoryID.String())

	err := c.Redis.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRepository) DeleteAllCategoryInCache(ctx context.Context) error {
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
