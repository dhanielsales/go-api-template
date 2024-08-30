package repository

import (
	"github.com/dhanielsales/golang-scaffold/internal/models"

	"github.com/dhanielsales/golang-scaffold/pkg/postgres"
	"github.com/dhanielsales/golang-scaffold/pkg/redis"

	postgresrepository "github.com/dhanielsales/golang-scaffold/internal/modules/store/repository/postgres"
	redisrepository "github.com/dhanielsales/golang-scaffold/internal/modules/store/repository/redis"
)

type StoreRepository struct {
	Postgres    *postgres.Storage
	Redis       *redis.Storage
	Persistence models.CategoryProductPersistenceRepository
	Cache       models.CategoryCacheRepository
}

func New(postgresStorage *postgres.Storage, redisStorage *redis.Storage) *StoreRepository {
	persistence := postgresrepository.New(postgresStorage)
	cache := redisrepository.New(redisStorage)

	return &StoreRepository{
		Postgres:    postgresStorage,
		Redis:       redisStorage,
		Persistence: persistence,
		Cache:       cache,
	}
}
