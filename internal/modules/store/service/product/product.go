package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"

	"github.com/google/uuid"
)

//go:generate mockgen -source ./$GOFILE -destination ./mock_$GOFILE -package $GOPACKAGE

// TODO
type ProductService interface {
	CreateProduct(ctx context.Context, data CreateProductPayload) (int64, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) (int64, error)
	GetManyProduct(ctx context.Context, params GetManyProductParams) ([]*models.Product, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, data UpdateProductPayload) (int64, error)
}

// TODO
type service struct {
	repository models.ProductRepository
}

var _ ProductService = (*service)(nil)

// TODO
func New(repository models.ProductRepository) *service {
	return &service{
		repository: repository,
	}
}
