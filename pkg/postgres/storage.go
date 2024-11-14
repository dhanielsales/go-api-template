package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Storage encapsulates a PostgreSQL database client, providing methods for managing connection cleanup and interacting with the database.
type Storage struct {
	Client *sql.DB
}

// New initializes and returns a new Storage instance with the given PostgreSQL client.
func New(client *sql.DB) *Storage {
	return &Storage{
		Client: client,
	}
}

// Cleanup closes the PostgreSQL database connection, releasing all allocated resources.
func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return fmt.Errorf("error closing postgress connection: %w", err)
	}

	return nil
}
