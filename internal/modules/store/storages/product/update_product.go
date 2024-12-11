package product

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"

	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"

	"github.com/google/uuid"
)

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

	res, err := r.Storage.UpdateProduct(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}
