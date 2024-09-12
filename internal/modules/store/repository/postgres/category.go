package postgres_repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/postgres"

	db "github.com/dhanielsales/go-api-template/internal/modules/store/repository/postgres/gen"
)

func (r *PostgresRepository) CreateCategory(ctx context.Context, category *models.Category) (int64, error) {
	params := db.CreateCategoryParams{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Queries.CreateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *PostgresRepository) UpdateCategory(ctx context.Context, id uuid.UUID, category *models.Category) (int64, error) {
	params := db.UpdateCategoryParams{
		ID:        id,
		Name:      category.Name,
		Slug:      category.Slug,
		UpdatedAt: sql.NullInt64{Int64: *category.UpdatedAt, Valid: true},
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Queries.UpdateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *PostgresRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	res, err := r.Queries.DeleteCategory(ctx, id)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *PostgresRepository) GetCategoryByID(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	return postgres.CallTx(ctx, r.Postgres.Client, func(tx *sql.Tx) (*models.Category, error) {
		repository := r.Queries.WithTx(tx)

		dbCategory, err := repository.GetCategoryById(ctx, id)
		if err != nil {
			return nil, err
		}

		dbProducts, err := repository.GetManyProductByCategoryId(ctx, id)
		if err != nil {
			return nil, err
		}

		category := ToCategory(&dbCategory)
		category.Products = ToProducts(dbProducts)

		return category, nil
	})
}

func (r *PostgresRepository) GetManyCategory(ctx context.Context, params models.GetManyCategoryPayload) ([]*models.Category, error) {
	pagination := postgres.Pagination(params.Page, params.PerPage)
	sorting := postgres.Sorting(params.OrderBy, params.OrderDirection)

	categories, err := r.Queries.GetManyCategory(ctx, db.GetManyCategoryParams{
		Limit:   pagination.Limit,
		Offset:  pagination.Offset,
		OrderBy: sorting,
	})
	if err != nil {
		return nil, err
	}

	return ToCategories(categories), nil
}
