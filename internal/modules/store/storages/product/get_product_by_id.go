package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"

	"github.com/google/uuid"
)

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := r.Storage.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return storages.ToProduct(&product), nil
}
