package category

import (
	"context"

	"github.com/google/uuid"
)

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	res, err := r.Storage.DeleteCategory(ctx, id)
	if err != nil {
		return 0, err
	}

	affecteds, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affecteds, nil
}
