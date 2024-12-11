package category

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

type UpdateCategoryPayload struct {
	Name        string
	Description string
}

func (s *service) UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateCategoryPayload) (int64, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx sqlutils.SQLTX) (int64, error) {
		queries := s.repository.WithTx(tx)

		category, err := queries.GetCategoryByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return 0, apperror.FromError(err).WithDescription("category not found").WithStatusCode(http.StatusNotFound)
			}

			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		category.Update(data.Name, data.Description)

		affected, err := queries.UpdateCategory(ctx, id, category)
		if err != nil {
			if sqlutils.IsUniqueViolationByField(err, "slug") {
				return 0, apperror.FromError(err).WithDescription("category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity)
			}

			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		err = s.repository.DeleteCategoryInCache(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		return affected, nil
	})
}
