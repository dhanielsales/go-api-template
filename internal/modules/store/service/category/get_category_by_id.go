package category

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
	"github.com/google/uuid"
)

func (s *service) GetCategoryByID(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx *sql.Tx) (*models.Category, error) {
		categoryInCache := s.repository.GetCategoryInCache(ctx, id)

		if categoryInCache != nil {
			return categoryInCache, nil
		}

		queries := s.repository.WithTx(tx)

		category, err := queries.GetCategoryByID(ctx, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, apperror.FromError(err).WithDescription("category not found").WithStatusCode(http.StatusNotFound)
			}

			return nil, apperror.FromError(err).WithDescription("can't process category entity")
		}

		err = s.repository.SetCategoryInCache(ctx, category, time.Hour*24)
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("can't process category entity")
		}

		return category, nil
	})
}
