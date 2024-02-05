package postgres_repository

import (
	"database/sql"

	"github.com/dhanielsales/golang-scaffold/entity"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"

	"github.com/dhanielsales/golang-scaffold/modules/store/repository/postgres/db"
)

type PostgresRepository struct {
	Postgres *postgres.Storage
	Queries  *db.Queries
}

func New(postgres *postgres.Storage) *PostgresRepository {
	return &PostgresRepository{
		Postgres: postgres,
	}
}

func (r *PostgresRepository) WithTx(tx *sql.Tx) entity.CategoryProductPersistenceRepository {
	return &PostgresRepository{
		Postgres: r.Postgres,
		Queries:  r.Queries.WithTx(tx),
	}
}
