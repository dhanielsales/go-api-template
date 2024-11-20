package redisutils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source ./$GOFILE -destination ./mock_$GOFILE -package $GOPACKAGE

// RedisClient defines an interface for interacting with a Redis database.
// It abstracts common Redis operations to allow flexibility in implementation.
type RedisClient interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error
	Ping(ctx context.Context) *redis.StatusCmd
	TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Close() error
}

// Storage encapsulates a Redis client, providing methods for initialization and cleanup.
type Storage struct {
	Client  RedisClient
	Client2 *redis.Client
}

// New initializes a new Storage instance with the given Redis client and verifies the connection.
func New(client RedisClient) (*Storage, error) {
	ctx := context.Background()
	ping := client.Ping(ctx)
	_, err := ping.Result()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: client,
	}, nil
}

// Cleanup closes the Redis client connection, releasing any allocated resources.
func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return fmt.Errorf("error closing redis connection: %w", err)
	}

	return nil
}
