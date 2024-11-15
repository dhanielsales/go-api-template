package category

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

type GetManyCategoryParams struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}

func (s *service) GetManyCategory(ctx context.Context, params GetManyCategoryParams) ([]*models.Category, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx *sql.Tx) ([]*models.Category, error) {
		queries := s.repository.WithTx(tx)

		dbResult, err := queries.GetManyCategory(ctx, models.GetManyCategoryPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("can't process category entity")
		}

		return dbResult, nil
	})
}
