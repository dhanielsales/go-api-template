package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	return sqlutils.WithTx(ctx, r.Postgres.Client, func(tx sqlutils.SQLTX) (*models.Category, error) {
		repo := r.Storage.WithTx(tx)

		dbCategory, err := repo.GetCategoryById(ctx, id)
		if err != nil {
			return nil, err
		}

		dbProducts, err := repo.GetManyProductByCategoryId(ctx, id)
		if err != nil {
			return nil, err
		}

		category := storages.ToCategory(&dbCategory)
		category.Products = storages.ToProducts(dbProducts)

		return category, nil
	})
}
