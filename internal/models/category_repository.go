package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *Category) (int64, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, category *Category) (int64, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (int64, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*Category, error)
	GetManyCategory(ctx context.Context, data GetManyCategoryPayload) ([]*Category, error)
	DeleteAllCategoryInCache(ctx context.Context) error
	DeleteCategoryInCache(ctx context.Context, categoryID uuid.UUID) error
	GetCategoryInCache(ctx context.Context, categoryID uuid.UUID) *Category
	SetCategoryInCache(ctx context.Context, category *Category, expiration time.Duration) error
	WithTx(tx *sql.Tx) CategoryRepository
	Client() *sql.DB
}

type GetManyCategoryPayload struct {
	Page           string
	PerPage        string
	OrderBy        string
	OrderDirection string
}
