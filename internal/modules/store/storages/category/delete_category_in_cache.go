package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/google/uuid"
)

func (c *CategoryRepository) DeleteCategoryInCache(ctx context.Context, categoryID uuid.UUID) error {
	key := redisutils.ComposeKey(CATEGORY_CACHE, categoryID.String())

	err := c.Redis.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
