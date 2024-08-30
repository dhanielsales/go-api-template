package redis_repository

import (
	"github.com/dhanielsales/golang-scaffold/internal/models"
	"github.com/dhanielsales/golang-scaffold/pkg/redis"
)

type CacheRepository struct {
	Redis *redis.Storage
}

func New(redisStorage *redis.Storage) models.CategoryCacheRepository {
	return &CacheRepository{
		Redis: redisStorage,
	}
}
