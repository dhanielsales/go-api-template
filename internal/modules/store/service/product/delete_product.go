package product

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"

	"github.com/google/uuid"
)

func (s *service) DeleteProduct(ctx context.Context, id uuid.UUID) (int64, error) {
	affected, err := s.repository.DeleteProduct(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, apperror.FromError(err).WithDescription("product not found").WithStatusCode(http.StatusNotFound)
		}

		return 0, apperror.FromError(err).WithDescription("can't process product entity").WithStatusCode(http.StatusUnprocessableEntity)
	}

	return affected, nil
}
