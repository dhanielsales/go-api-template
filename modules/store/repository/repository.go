package repository

import (
	"github.com/dhanielsales/golang-scaffold/entity"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	fake_repository "github.com/dhanielsales/golang-scaffold/modules/store/repository/fake"
	postgres_repository "github.com/dhanielsales/golang-scaffold/modules/store/repository/postgres"
	redis_repository "github.com/dhanielsales/golang-scaffold/modules/store/repository/redis"
)

type StoreRepository struct {
	Postgres    *postgres.Storage
	Redis       *redis.Storage
	Persistence entity.CategoryProductPersistenceRepository
	Fake        entity.CategoryProductPersistenceRepository
	Cache       entity.CategoryCacheRepository
}

func New(postgresStorage *postgres.Storage, redisStorage *redis.Storage) *StoreRepository {
	fake := fake_repository.New()
	persistence := postgres_repository.New(postgresStorage)
	cache := redis_repository.New(redisStorage)

	return &StoreRepository{
		Postgres:    postgresStorage,
		Redis:       redisStorage,
		Persistence: persistence,
		Cache:       cache,
		Fake:        fake,
	}
}
