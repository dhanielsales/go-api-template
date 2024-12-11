package category

import (
	"context"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

func (s *service) DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error) {
	return sqlutils.WithTx(ctx, s.repository.Client(), func(tx sqlutils.SQLTX) (int64, error) {
		queries := s.repository.WithTx(tx)

		affected, err := queries.DeleteCategory(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		err = s.repository.DeleteCategoryInCache(ctx, id)
		if err != nil {
			return 0, apperror.FromError(err).WithDescription("can't process category entity")
		}

		return affected, nil
	})
}
