package fake_repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/dhanielsales/golang-scaffold/entity"
)

func (r *FakeRepository) CreateCategory(ctx context.Context, category *entity.Category) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) UpdateCategory(ctx context.Context, id uuid.UUID, category *entity.Category) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) GetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return &entity.Category{
		ID:        id,
		Name:      "Category",
		Slug:      "category",
		Products:  &[]entity.Product{},
		CreatedAt: 1707250478084,
	}, nil

}

func (r *FakeRepository) GetManyCategory(ctx context.Context, params entity.GetManyCategoryPayload) (*[]entity.Category, error) {
	return &[]entity.Category{
		{
			ID:        uuid.New(),
			Name:      "Category",
			Slug:      "category",
			Products:  &[]entity.Product{},
			CreatedAt: 1707250478084,
		},
	}, nil
}
