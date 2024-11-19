package sqlutils

import (
	"context"
	"database/sql"
	"fmt"
)

//go:generate mockgen -source ./$GOFILE -destination ./mock_$GOFILE -package $GOPACKAGE

// Querier defines an interface for executing SQL queries and statements.
// It abstracts the basic database operations to support flexibility for various implementations.
type Querier interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// SQLTX represents an SQL transaction, extending the Querier interface with methods to manage transaction lifecycle.
type SQLTX interface {
	Querier
	Commit() error
	Rollback() error
}

// SQLDB represents an SQL database client, extending the Querier interface with methods for transaction management and connection lifecycle.
type SQLDB interface {
	Querier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTX, error)
	Close() error
	Ping() error
}

// Storage encapsulates a SQL database client, providing methods for managing connection cleanup and interacting with the database.
type Storage struct {
	Client SQLDB
}

// New initializes and returns a new Storage instance with the given SQL client.
func New(client SQLDB) *Storage {
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
