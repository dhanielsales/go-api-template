package category

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) (int64, error) {
	params := db.CreateCategoryParams{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Storage.CreateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id uuid.UUID, category *models.Category) (int64, error) {
	params := db.UpdateCategoryParams{
		ID:        id,
		Name:      category.Name,
		Slug:      category.Slug,
		UpdatedAt: sql.NullInt64{Int64: *category.UpdatedAt, Valid: true},
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Storage.UpdateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	res, err := r.Storage.DeleteCategory(ctx, id)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

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

func (r *CategoryRepository) GetManyCategory(ctx context.Context, params models.GetManyCategoryPayload) ([]*models.Category, error) {
	pagination := sqlutils.Pagination(params.Page, params.PerPage)
	sorting := sqlutils.Sorting(params.OrderBy, params.OrderDirection)

	categories, err := r.Storage.GetManyCategory(ctx, db.GetManyCategoryParams{
		Limit:   pagination.Limit,
		Offset:  pagination.Offset,
		OrderBy: sorting,
	})
	if err != nil {
		return nil, err
	}

	return storages.ToCategories(categories), nil
}
