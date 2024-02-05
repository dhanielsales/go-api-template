package application

import (
	"github.com/dhanielsales/golang-scaffold/modules/store/external"
	"github.com/dhanielsales/golang-scaffold/modules/store/repository"
)

type StoreService struct {
	repository *repository.StoreRepository
	external   *external.External
}

func New(repository *repository.StoreRepository, external *external.External) *StoreService {
	return &StoreService{
		repository: repository,
		external:   external,
	}
}
