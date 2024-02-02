package storage

import (
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	"github.com/dhanielsales/golang-scaffold/entity"
	q "github.com/dhanielsales/golang-scaffold/modules/store/storage/postgres"
	c "github.com/dhanielsales/golang-scaffold/modules/store/storage/redis"
)

type StoreStorage struct {
	Postgres *postgres.Storage
	Redis    *redis.Storage
	Queries  *q.Queries
	Cache    *c.Cache
}

func New(postgresStorage *postgres.Storage, redisStorage *redis.Storage) *StoreStorage {
	queries := q.New(postgresStorage.Client)
	cache := c.New(redisStorage)

	return &StoreStorage{
		Postgres: postgresStorage,
		Redis:    redisStorage,
		Queries:  queries,
		Cache:    cache,
	}
}

func ToCategory(category *q.Category) *entity.Category {
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

func ToProduct(category *q.Product) *entity.Product {
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
