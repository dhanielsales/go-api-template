package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	_ "github.com/lib/pq"
)

// NewPostgresDB initializes a new connection to a PostgreSQL database using the provided data source name.
// It returns a postgresDB instance or an error if the connection fails.
func NewPostgresDB(dataSourceName string) (*postgresDB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres connection: %w", err)
	}

	return &postgresDB{db}, nil
}

// Ensure postgresDB implements the sqlutils.SQLDB interface.
var _ sqlutils.SQLDB = (*postgresDB)(nil)

// postgresDB wraps an *sql.DB instance to implement the sqlutils.SQLDB interface.
type postgresDB struct {
	*sql.DB
}

// BeginTx starts a new database transaction with the provided context and transaction options.
// It returns an instance of sqlutils.SQLTX or an error if the transaction fails.
func (pdb postgresDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (sqlutils.SQLTX, error) {
	tx, err := pdb.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}
	return tx, nil
}
