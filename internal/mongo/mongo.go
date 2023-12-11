package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
	dbName string
}

func Bootstrap(uri, dbName string, timeout time.Duration) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: client,
		dbName: dbName,
	}, nil
}

func (s *Storage) Cleanup() error {
	return s.Client.Disconnect(context.Background())
}

func (s *Storage) Database() *mongo.Database {
	return s.Client.Database(s.dbName)
}
