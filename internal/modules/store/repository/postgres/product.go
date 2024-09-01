package postgres_repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/dhanielsales/go-api-template/pkg/postgres"

	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/repository/postgres/gen"
)

func ToProduct(category *db.Product) *models.Product {
	res := models.Product{
		ID:         category.ID,
		Name:       category.Name,
		Slug:       category.Slug,
		Price:      category.Price,
		CategoryID: category.CategoryID,
		CreatedAt:  category.CreatedAt,
	}

	if category.UpdatedAt.Valid {
		res.UpdatedAt = &category.UpdatedAt.Int64
	}

	if category.Description.Valid {
		res.Description = &category.Description.String
	}

	return &res
}

func (r *PostgresRepository) CreateProduct(ctx context.Context, product *models.Product) (*int64, error) {
	params := db.CreateProductParams{
		ID:         product.ID,
		Name:       product.Name,
		Slug:       product.Slug,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}

	if product.Description != nil {
		params.Description = sql.NullString{String: *product.Description, Valid: true}
	}

	res, err := r.Queries.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) UpdateProduct(ctx context.Context, id uuid.UUID, product *models.Product) (*int64, error) {
	params := db.UpdateProductParams{
		ID:         id,
		Name:       product.Name,
		Slug:       product.Slug,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		UpdatedAt:  sql.NullInt64{Int64: *product.UpdatedAt, Valid: true},
	}

	if product.Description != nil {
		params.Description = sql.NullString{String: *product.Description, Valid: true}
	}

	res, err := r.Queries.UpdateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) DeleteProduct(ctx context.Context, id uuid.UUID) (*int64, error) {
	res, err := r.Queries.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affecteds, nil
}

func (r *PostgresRepository) GetProductById(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := r.Queries.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return ToProduct(&product), nil
}

func (r *PostgresRepository) GetManyProduct(ctx context.Context, params models.GetManyProductPayload) (*[]models.Product, error) {
	pagination := postgres.Pagination(params.Page, params.PerPage)
	sorting := postgres.Sorting(params.OrderBy, params.OrderDirection)

	products, err := r.Queries.GetManyProduct(ctx, db.GetManyProductParams{
		OrderBy: sorting,
		Offset:  pagination.Offset,
		Limit:   pagination.Limit,
	})
	if err != nil {
		return nil, err
	}

	res := make([]models.Product, len(products))
	for _, product := range products {
		res = append(res, *ToProduct(&product))
	}

	return &res, nil
}

func (r *PostgresRepository) GetManyProductByCategoryId(ctx context.Context, categoryID uuid.UUID) (*[]models.Product, error) {
	products, err := r.Queries.GetManyProductByCategoryId(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	res := make([]models.Product, len(products))

	for _, product := range products {
		res = append(res, *ToProduct(&product))
	}

	return &res, nil
}
