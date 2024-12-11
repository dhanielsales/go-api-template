package category

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
)

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) (int64, error) {
	params := db.CreateCategoryParams{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Storage.CreateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}
