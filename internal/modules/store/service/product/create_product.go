package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"

	"github.com/google/uuid"
)

type CreateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategotyID  uuid.UUID
}

func (s *service) CreateProduct(ctx context.Context, data CreateProductPayload) (int64, error) {
	product, err := models.NewProduct(data.Name, data.Description, data.Price, data.CategotyID)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("can't process product entity")
	}

	affected, err := s.repository.CreateProduct(ctx, product)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("can't process product entity")
	}

	return affected, nil
}
