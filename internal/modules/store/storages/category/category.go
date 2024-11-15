package category

import (
	"database/sql"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
)

type CategoryRepository struct {
	Postgres *sqlutils.Storage
	Queries  db.QueryWrapper
	Redis    *redisutils.Storage // TODO convert to interface to allow mocks to unit tests
}

func New(sql *sqlutils.Storage, queries db.QueryWrapper, redis *redisutils.Storage) *CategoryRepository {
	return &CategoryRepository{
		Postgres: sql,
		Queries:  queries,
		Redis:    redis,
	}
}

func NewWithDefaultQueries(sql *sqlutils.Storage, redis *redisutils.Storage) *CategoryRepository {
	return New(sql, db.NewQueryWrapper(sql.Client), redis)
}

func (r *CategoryRepository) WithTx(tx *sql.Tx) models.CategoryRepository {
	return &CategoryRepository{
		Postgres: r.Postgres,
		Queries:  r.Queries.WithTx(tx),
	}
}

func (r *CategoryRepository) Client() *sql.DB {
	return r.Postgres.Client
}
