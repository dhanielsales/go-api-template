package category

import (
	"context"

	"github.com/dhanielsales/go-api-template/internal/models"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, data CreateCategoryPayload) (int64, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*models.Category, error)
	GetManyCategory(ctx context.Context, params GetManyCategoryParams) ([]*models.Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateCategoryPayload) (int64, error)
}

type service struct {
	repository models.CategoryRepository
}

var _ CategoryService = (*service)(nil)

func New(repository models.CategoryRepository) *service {
	return &service{
		repository: repository,
	}
}
