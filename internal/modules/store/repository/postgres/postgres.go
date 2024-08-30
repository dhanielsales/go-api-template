package postgres_repository

import (
	"database/sql"

	"github.com/dhanielsales/golang-scaffold/internal/models"
	"github.com/dhanielsales/golang-scaffold/pkg/postgres"

	"github.com/dhanielsales/golang-scaffold/internal/modules/store/repository/postgres/db"
)

type PostgresRepository struct {
	Postgres *postgres.Storage
	Queries  *db.Queries
}

func New(postgres *postgres.Storage) *PostgresRepository {
	return &PostgresRepository{
		Postgres: postgres,
		Queries:  db.New(postgres.Client),
	}
}

func (r *PostgresRepository) WithTx(tx *sql.Tx) models.CategoryProductPersistenceRepository {
	return &PostgresRepository{
		Postgres: r.Postgres,
		Queries:  r.Queries.WithTx(tx),
	}
}
