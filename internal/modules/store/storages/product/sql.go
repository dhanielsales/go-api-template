package product

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"

	"github.com/google/uuid"
)

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) (int64, error) {
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
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id uuid.UUID, product *models.Product) (int64, error) {
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
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) (int64, error) {
	res, err := r.Queries.DeleteProduct(ctx, id)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := r.Queries.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return storages.ToProduct(&product), nil
}

func (r *ProductRepository) GetManyProduct(ctx context.Context, params models.GetManyProductPayload) ([]*models.Product, error) {
	pagination := sqlutils.Pagination(params.Page, params.PerPage)
	sorting := sqlutils.Sorting(params.OrderBy, params.OrderDirection)

	products, err := r.Queries.GetManyProduct(ctx, db.GetManyProductParams{
		OrderBy: sorting,
		Offset:  pagination.Offset,
		Limit:   pagination.Limit,
	})
	if err != nil {
		return nil, err
	}

	return storages.ToProducts(products), nil
}

func (r *ProductRepository) GetManyProductByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]*models.Product, error) {
	products, err := r.Queries.GetManyProductByCategoryId(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return storages.ToProducts(products), nil
}
