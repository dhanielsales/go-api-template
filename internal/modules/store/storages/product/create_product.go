package product

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"

	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
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

	res, err := r.Storage.CreateProduct(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}
