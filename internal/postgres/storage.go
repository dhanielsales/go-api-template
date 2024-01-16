package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	Client *sql.DB
}

func Bootstrap(url string) (*Storage, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: db,
	}, nil
}

func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return err
	}

	return nil
}
