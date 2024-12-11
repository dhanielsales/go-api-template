package category

import (
	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"

	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"
)

type CategoryRepository struct {
	Postgres *sqlutils.Storage
	Storage  storages.Storage
	Redis    *redisutils.Storage
}

const CATEGORY_CACHE = "category"

func New(sql *sqlutils.Storage, storage storages.Storage, redis *redisutils.Storage) *CategoryRepository {
	return &CategoryRepository{
		Postgres: sql,
		Storage:  storage,
		Redis:    redis,
	}
}

func NewWithDefaultStorage(sql *sqlutils.Storage, redis *redisutils.Storage) *CategoryRepository {
	return New(sql, storages.NewStorage(sql.Client), redis)
}

func (r *CategoryRepository) WithTx(tx sqlutils.SQLTX) models.CategoryRepository {
	return &CategoryRepository{
		Postgres: r.Postgres,
		Redis:    r.Redis,
		Storage:  storages.NewStorage(tx),
	}
}

func (r *CategoryRepository) Client() sqlutils.SQLDB {
	return r.Postgres.Client
}
