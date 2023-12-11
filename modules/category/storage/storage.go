package category_storage

import (
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
)

type CategoryStorage struct {
	postgresDb *postgres.Storage
}

func NewCategoryStorage(db *postgres.Storage) *CategoryStorage {
	return &CategoryStorage{
		postgresDb: db,
	}
}
