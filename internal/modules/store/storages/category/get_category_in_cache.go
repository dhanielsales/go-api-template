package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/redisutils"

	"github.com/google/uuid"
)

func (c *CategoryRepository) GetCategoryInCache(ctx context.Context, categoryID uuid.UUID) *models.Category {
	key := redisutils.ComposeKey(CATEGORY_CACHE, categoryID.String())

	var category models.Category

	err := c.Redis.Client.Get(ctx, key).Scan(&category)
	if err != nil {
		return nil
	}

	return &category
}
