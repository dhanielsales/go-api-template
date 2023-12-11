package postgres

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	Client *sqlx.DB
}

func Bootstrap(uri string) (*Storage, error) {
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Client: db,
	}, nil
}

func (s *Storage) Cleanup() error {
	return s.Client.Close()
}
