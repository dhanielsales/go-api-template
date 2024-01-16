package redis

import (
	"github.com/dhanielsales/golang-scaffold/internal/redis"
)

type Cache struct {
	Redis *redis.Storage
}

func New(redisStorage *redis.Storage) *Cache {
	return &Cache{
		Redis: redisStorage,
	}
}
