package service

import (
	"github.com/dhanielsales/golang-scaffold/internal/modules/store/repository"
)

type StoreService struct {
	repository *repository.StoreRepository
}

func New(repository *repository.StoreRepository) *StoreService {
	return &StoreService{
		repository: repository,
	}
}
