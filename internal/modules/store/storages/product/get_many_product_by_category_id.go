package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"

	"github.com/google/uuid"
)

func (r *ProductRepository) GetManyProductByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]*models.Product, error) {
	products, err := r.Storage.GetManyProductByCategoryId(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return storages.ToProducts(products), nil
}
