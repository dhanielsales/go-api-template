package product

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

func (s *service) GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx sqlutils.SQLTX) (*models.Product, error) {
		queries := s.repository.WithTx(tx)

		product, err := queries.GetProductByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, apperror.FromError(err).WithDescription("product not found").WithStatusCode(http.StatusNotFound)
			}

			return nil, apperror.FromError(err).WithDescription("can't process product entity")
		}

		return product, nil
	})
}
