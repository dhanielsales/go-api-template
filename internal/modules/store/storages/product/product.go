package product

import (
	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"

	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

type ProductRepository struct {
	Postgres *sqlutils.Storage
	Storage  storages.Storage
}

func New(sql *sqlutils.Storage, storage storages.Storage) *ProductRepository {
	return &ProductRepository{
		Postgres: sql,
		Storage:  storage,
	}
}

func NewWithDefaultStorage(sql *sqlutils.Storage) *ProductRepository {
	return New(sql, storages.NewStorage(sql.Client))
}

func (r *ProductRepository) WithTx(tx sqlutils.SQLTX) models.ProductRepository {
	return &ProductRepository{
		Postgres: r.Postgres,
		Storage:  storages.NewStorage(tx),
	}
}

func (r *ProductRepository) Client() sqlutils.SQLDB {
	return r.Postgres.Client
}
