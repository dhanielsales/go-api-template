package fake_repository

import (
	"database/sql"

	"github.com/dhanielsales/golang-scaffold/entity"
)

type FakeRepository struct {
}

func New() *FakeRepository {
	return &FakeRepository{}
}

func (r *FakeRepository) WithTx(tx *sql.Tx) entity.CategoryProductPersistenceRepository {
	return &FakeRepository{}
}
