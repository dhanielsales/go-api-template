package redis_repository

import (
	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/redis"
)

type CacheRepository struct {
	Redis *redis.Storage
}

func New(redisStorage *redis.Storage) models.CategoryCacheRepository {
	return &CacheRepository{
		Redis: redisStorage,
	}
}
