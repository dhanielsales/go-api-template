package category

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

type CreateCategoryPayload struct {
	Name        string
	Description string
}

func (s *service) CreateCategory(ctx context.Context, data CreateCategoryPayload) (int64, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx *sql.Tx) (int64, error) {
		queries := s.repository.WithTx(tx)

		category, err := models.NewCategory(data.Name, data.Description)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		affecteds, err := queries.CreateCategory(ctx, category)
		if err != nil {
			if sqlutils.IsUniqueViolationByField(err, "slug") {
				return 0, apperror.FromError(err).WithDescription("category with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity)
			}

			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		err = s.repository.DeleteAllCategoryInCache(ctx)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		return affecteds, nil
	})
}
