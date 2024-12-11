package product

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
)

func (r *ProductRepository) GetManyProduct(ctx context.Context, params models.GetManyProductPayload) ([]*models.Product, error) {
	pagination := sqlutils.Pagination(params.Page, params.PerPage)
	sorting := sqlutils.Sorting(params.OrderBy, params.OrderDirection)

	products, err := r.Storage.GetManyProduct(ctx, db.GetManyProductParams{
		OrderBy: sorting,
		Offset:  pagination.Offset,
		Limit:   pagination.Limit,
	})
	if err != nil {
		return nil, err
	}

	return storages.ToProducts(products), nil
}
