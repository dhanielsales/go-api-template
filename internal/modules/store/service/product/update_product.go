package product

import (
	"context"
	"database/sql"
	"net/http"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

type UpdateProductPayload struct {
	Name        string
	Description string
	Price       float64
	CategoryID  uuid.UUID
}

func (s *service) UpdateProduct(ctx context.Context, id uuid.UUID, data UpdateProductPayload) (int64, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx sqlutils.SQLTX) (int64, error) {
		queries := s.repository.WithTx(tx)

		product, err := queries.GetProductByID(ctx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return 0, apperror.FromError(err).WithDescription("product not found").WithStatusCode(http.StatusNotFound)
			}

			return 0, apperror.FromError(err).WithDescription("can't process product entity")
		}

		product.Update(data.Name, data.Description, data.Price, data.CategoryID)

		affected, err := queries.UpdateProduct(ctx, id, product)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process product entity")
		}

		return affected, nil
	})
}
