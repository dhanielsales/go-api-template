package category

import (
	"context"
	"time"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/redisutils"
)

func (c *CategoryRepository) SetCategoryInCache(ctx context.Context, category *models.Category, expiration time.Duration) error {
	key := redisutils.ComposeKey(CATEGORY_CACHE, category.ID.String())

	err := c.Redis.Client.Set(ctx, key, category, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
