package service

import (
	storerepository "github.com/dhanielsales/go-api-template/internal/modules/store/repository"
)

type StoreService struct {
	repository *storerepository.StoreRepository
}

func New(repository *storerepository.StoreRepository) *StoreService {
	return &StoreService{
		repository: repository,
	}
}
