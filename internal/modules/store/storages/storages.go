package storages

import (
	"context"
	"database/sql"

	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/google/uuid"
)

// Storage defines the interface for interacting with storage operations for categories and products.
type Storage interface {
	CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (sql.Result, error)
	CreateProduct(ctx context.Context, arg db.CreateProductParams) (sql.Result, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (sql.Result, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) (sql.Result, error)
	GetCategoryById(ctx context.Context, id uuid.UUID) (db.Category, error)
	GetManyCategory(ctx context.Context, arg db.GetManyCategoryParams) ([]db.Category, error)
	GetManyProduct(ctx context.Context, arg db.GetManyProductParams) ([]db.Product, error)
	GetManyProductByCategoryId(ctx context.Context, categoryID uuid.UUID) ([]db.Product, error)
	GetProductById(ctx context.Context, id uuid.UUID) (db.Product, error)
	UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) (sql.Result, error)
	UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (sql.Result, error)
	WithTx(tx sqlutils.SQLTX) Storage
}

var _ Storage = (*storage)(nil)

// storage provides a concrete implementation of the Storage interface using a database client.
type storage struct {
	client  sqlutils.Querier
	queries *db.Queries
}

// WithTx returns a new storage instance that uses the given transaction for all operations.
func NewStorage(client sqlutils.Querier) *storage {
	return &storage{
		client:  client,
		queries: db.New(),
	}
}

func (s *storage) WithTx(tx sqlutils.SQLTX) Storage {
	return &storage{
		queries: s.queries,
		client:  tx,
	}
}

func (s *storage) CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (sql.Result, error) {
	return s.queries.CreateCategory(ctx, s.client, arg)
}

func (s *storage) CreateProduct(ctx context.Context, arg db.CreateProductParams) (sql.Result, error) {
	return s.queries.CreateProduct(ctx, s.client, arg)
}

func (s *storage) DeleteCategory(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return s.queries.DeleteCategory(ctx, s.client, id)
}

func (s *storage) DeleteProduct(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return s.queries.DeleteProduct(ctx, s.client, id)
}

func (s *storage) GetCategoryById(ctx context.Context, id uuid.UUID) (db.Category, error) {
	return s.queries.GetCategoryById(ctx, s.client, id)
}

func (s *storage) GetManyCategory(ctx context.Context, arg db.GetManyCategoryParams) ([]db.Category, error) {
	return s.queries.GetManyCategory(ctx, s.client, arg)
}

func (s *storage) GetManyProduct(ctx context.Context, arg db.GetManyProductParams) ([]db.Product, error) {
	return s.queries.GetManyProduct(ctx, s.client, arg)
}

func (s *storage) GetManyProductByCategoryId(ctx context.Context, categoryID uuid.UUID) ([]db.Product, error) {
	return s.queries.GetManyProductByCategoryId(ctx, s.client, categoryID)
}

func (s *storage) GetProductById(ctx context.Context, id uuid.UUID) (db.Product, error) {
	return s.queries.GetProductById(ctx, s.client, id)
}

func (s *storage) UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) (sql.Result, error) {
	return s.queries.UpdateCategory(ctx, s.client, arg)
}

func (s *storage) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (sql.Result, error) {
	return s.queries.UpdateProduct(ctx, s.client, arg)
}
