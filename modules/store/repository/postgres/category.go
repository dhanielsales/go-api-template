package postgres_repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/dhanielsales/golang-scaffold/entity"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/modules/store/repository/postgres/db"
)

func ToCategory(category *db.Category) *entity.Category {
	res := entity.Category{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.UpdatedAt.Valid {
		res.UpdatedAt = &category.UpdatedAt.Int64
	}

	if category.Description.Valid {
		res.Description = &category.Description.String
	}

	return &res
}

func (r *PostgresRepository) CreateCategory(ctx context.Context, category *entity.Category) (*int64, error) {
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
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) UpdateCategory(ctx context.Context, id uuid.UUID, category *entity.Category) (*int64, error) {
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
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (*int64, error) {
	res, err := r.Queries.DeleteCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) GetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return postgres.CallTx(ctx, r.Postgres.Client, func(tx *sql.Tx) (*entity.Category, error) {
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

		products := make([]entity.Product, len(dbProducts))

		for i, dbProduct := range dbProducts {
			products[i] = *ToProduct(&dbProduct)
		}

		category.Products = &products

		return category, nil
	})
}

func (r *PostgresRepository) GetManyCategory(ctx context.Context, params entity.GetManyCategoryPayload) (*[]entity.Category, error) {
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

	res := make([]entity.Category, len(categories))
	for i, category := range categories {
		res[i] = *ToCategory(&category)
	}

	return &res, nil
}
