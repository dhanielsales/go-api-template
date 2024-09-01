package postgres_repository

import (
	"database/sql"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/pkg/postgres"

	db "github.com/dhanielsales/go-api-template/internal/modules/store/repository/postgres/gen"
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
