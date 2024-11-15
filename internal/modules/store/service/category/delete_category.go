package category

import (
	"context"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"

	"github.com/google/uuid"
)

func (s *service) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	affected, err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("can't process category entity")
	}

	err = s.repository.DeleteCategoryInCache(ctx, id)
	if err != nil {
		return 0, apperror.FromError(err).WithDescription("can't process category entity")
	}

	return affected, nil
}
