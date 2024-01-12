package application

import store_storage "github.com/dhanielsales/golang-scaffold/modules/store/storage"

type StoreService struct {
	storage *store_storage.StoreStorage
}

func NewService(storage *store_storage.StoreStorage) *StoreService {
	return &StoreService{
		storage: storage,
	}
}
