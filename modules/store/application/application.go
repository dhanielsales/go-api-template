package application

import (
	"github.com/dhanielsales/golang-scaffold/modules/store/external"
	store_storage "github.com/dhanielsales/golang-scaffold/modules/store/storage"
)

type StoreService struct {
	storage  *store_storage.StoreStorage
	external *external.External
}

func New(storage *store_storage.StoreStorage, external *external.External) *StoreService {
	return &StoreService{
		storage:  storage,
		external: external,
	}
}
