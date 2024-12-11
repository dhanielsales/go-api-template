package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

func (r *CategoryRepository) GetManyCategory(ctx context.Context, params models.GetManyCategoryPayload) ([]*models.Category, error) {
	pagination := sqlutils.Pagination(params.Page, params.PerPage)
	sorting := sqlutils.Sorting(params.OrderBy, params.OrderDirection)

	categories, err := r.Storage.GetManyCategory(ctx, db.GetManyCategoryParams{
		Limit:   pagination.Limit,
		Offset:  pagination.Offset,
		OrderBy: sorting,
	})
	if err != nil {
		return nil, err
	}

	return storages.ToCategories(categories), nil
}
