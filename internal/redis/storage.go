package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Client *redis.Client
}

func Bootstrap(url string) (*Storage, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	ctx := context.Background()
	ping := client.Ping(ctx)
	_, err = ping.Result()

	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: client,
	}, nil
}

func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return err
	}

	return nil
}
