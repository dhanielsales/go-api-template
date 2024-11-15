package product

import (
	"database/sql"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/storages/gen"
)

type ProductRepository struct {
	Postgres *sqlutils.Storage
	Queries  db.QueryWrapper
}

func New(sql *sqlutils.Storage, queries db.QueryWrapper) *ProductRepository {
	return &ProductRepository{
		Postgres: sql,
		Queries:  queries,
	}
}

func NewWithDefaultQueries(sql *sqlutils.Storage) *ProductRepository {
	return New(sql, db.NewQueryWrapper(sql.Client))
}

func (r *ProductRepository) WithTx(tx *sql.Tx) models.ProductRepository { // TODO any
	return &ProductRepository{
		Postgres: r.Postgres,
		Queries:  r.Queries.WithTx(tx),
	}
}

func (r *ProductRepository) Client() *sql.DB {
	return r.Postgres.Client
}
