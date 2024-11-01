package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Client *sql.DB
}

func New(client *sql.DB) *Storage {
	return &Storage{
		Client: client,
	}
}

func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return fmt.Errorf("error closing postgress connection: %w", err)
	}

	return nil
}
