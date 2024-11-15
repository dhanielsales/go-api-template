package sqlutils

import (
	"database/sql"
	"fmt"
)

// Storage encapsulates a SQL database client, providing methods for managing connection cleanup and interacting with the database.
type Storage struct {
	Client *sql.DB
}

// New initializes and returns a new Storage instance with the given SQL client.
func New(client *sql.DB) *Storage {
	return &Storage{
		Client: client,
	}
}

// Cleanup closes the SQL database connection, releasing all allocated resources.
func (s *Storage) Cleanup() error {
	err := s.Client.Close()
	if err != nil {
		return fmt.Errorf("error closing postgress connection: %w", err)
	}

	return nil
}
