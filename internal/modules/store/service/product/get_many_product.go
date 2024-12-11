package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

type GetManyProductParams struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}

func (s *service) GetManyProduct(ctx context.Context, params GetManyProductParams) ([]*models.Product, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx sqlutils.SQLTX) ([]*models.Product, error) {
		queries := s.repository.WithTx(tx)

		products, err := queries.GetManyProduct(ctx, models.GetManyProductPayload{
			Page:           params.Page,
			PerPage:        params.PerPage,
			OrderBy:        params.OrderBy,
			OrderDirection: params.OrderDirection,
		})
		if err != nil {
			return nil, apperror.FromError(err).WithDescription("can't process product entity")
		}

		return products, nil
	})
}
