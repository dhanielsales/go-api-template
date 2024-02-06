package fake_repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/dhanielsales/golang-scaffold/entity"
)

func (r *FakeRepository) CreateProduct(ctx context.Context, product *entity.Product) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) UpdateProduct(ctx context.Context, id uuid.UUID, product *entity.Product) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) DeleteProduct(ctx context.Context, id uuid.UUID) (*int64, error) {
	affecteds := int64(0)
	return &affecteds, nil
}

func (r *FakeRepository) GetProductById(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return &entity.Product{
		ID:         id,
		Name:       "Category",
		Slug:       "category",
		Price:      1000,
		CategoryID: uuid.New(),
		CreatedAt:  1707250478084,
	}, nil
}

func (r *FakeRepository) GetManyProduct(ctx context.Context, params entity.GetManyProductPayload) (*[]entity.Product, error) {
	return &[]entity.Product{
		{
			ID:         uuid.New(),
			Name:       "Category",
			Slug:       "category",
			Price:      1000,
			CategoryID: uuid.New(),
			CreatedAt:  1707250478084,
		},
	}, nil
}

func (r *FakeRepository) GetManyProductByCategoryId(ctx context.Context, categoryID uuid.UUID) (*[]entity.Product, error) {
	return &[]entity.Product{
		{
			ID:         uuid.New(),
			Name:       "Category",
			Slug:       "category",
			Price:      1000,
			CategoryID: categoryID,
			CreatedAt:  1707250478084,
		},
	}, nil
}
