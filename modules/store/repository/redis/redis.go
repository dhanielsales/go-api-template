package redis_repository

import (
	"github.com/dhanielsales/golang-scaffold/entity"
	"github.com/dhanielsales/golang-scaffold/internal/redis"
)

type CacheRepository struct {
	Redis *redis.Storage
}

func New(redisStorage *redis.Storage) entity.CategoryCacheRepository {
	return &CacheRepository{
		Redis: redisStorage,
	}
}
