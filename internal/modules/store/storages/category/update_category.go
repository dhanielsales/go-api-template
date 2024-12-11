package category

import (
	"context"
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"

	"github.com/google/uuid"
)

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id uuid.UUID, category *models.Category) (int64, error) {
	params := db.UpdateCategoryParams{
		ID:        id,
		Name:      category.Name,
		Slug:      category.Slug,
		UpdatedAt: sql.NullInt64{Int64: *category.UpdatedAt, Valid: true},
	}

	if category.Description != nil {
		params.Description = sql.NullString{String: *category.Description, Valid: true}
	}

	res, err := r.Storage.UpdateCategory(ctx, params)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}
