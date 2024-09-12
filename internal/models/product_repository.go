package models

import (
	"context"

	"github.com/google/uuid"
)

type ProductPersistenceRepository interface {
	CreateProduct(ctx context.Context, data *Product) (int64, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, data *Product) (int64, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) (int64, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*Product, error)
	GetManyProduct(ctx context.Context, data GetManyProductPayload) ([]*Product, error)
	GetManyProductByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]*Product, error)
}

type GetManyProductPayload struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}
