package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Storage encapsulates a Redis client, providing methods for initialization and cleanup.
type Storage struct {
	Client *redis.Client
}

// New initializes a new Storage instance with the given Redis client and verifies the connection.
func New(client *redis.Client) (*Storage, error) {
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
