package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Client *redis.Client
}

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

func NewDefault() (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return New(client)
}

func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return err
	}

	return nil
}
