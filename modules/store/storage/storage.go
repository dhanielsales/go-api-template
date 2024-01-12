package storage

import (
	"github.com/dhanielsales/golang-scaffold/internal/postgres"

	"github.com/dhanielsales/golang-scaffold/modules/store/entity"
	postgresStorage "github.com/dhanielsales/golang-scaffold/modules/store/storage/postgres"
)

type StoreStorage struct {
	Db      *postgres.Storage
	Queries *postgresStorage.Queries
}

func NewStorage(db *postgres.Storage) *StoreStorage {
	queries := postgresStorage.New(db.Client)

	return &StoreStorage{
		Db:      db,
		Queries: queries,
	}
}

func ToCategory(category *postgresStorage.Category) *entity.Category {
	res := entity.Category{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.UpdatedAt.Valid {
		res.UpdatedAt = &category.UpdatedAt.Int64
	}

	if category.Description.Valid {
		res.Description = &category.Description.String
	}

	return &res
}

func ToProduct(category *postgresStorage.Product) *entity.Product {
	res := entity.Product{
		ID:         category.ID,
		Name:       category.Name,
		Slug:       category.Slug,
		Price:      category.Price,
		CategoryID: category.CategoryID,
		CreatedAt:  category.CreatedAt,
	}

	if category.UpdatedAt.Valid {
		res.UpdatedAt = &category.UpdatedAt.Int64
	}

	if category.Description.Valid {
		res.Description = &category.Description.String
	}

	return &res
}
