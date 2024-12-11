package product

import (
	"context"
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

type CreateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func (s *service) CreateProduct(ctx context.Context, data CreateProductPayload) (int64, error) {
	product, err := models.NewProduct(data.Name, data.Description, data.Price, data.CategoryID)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("can't process product entity")
	}

	affected, err := s.repository.CreateProduct(ctx, product)
	if err != nil {
		if sqlutils.IsUniqueViolationByField(err, "slug") {
			return 0, apperror.FromError(err).WithDescription("product with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity)
		}

		return 0, apperror.FromError(err).WithDescription("can't process product entity")
	}

	return affected, nil
}
