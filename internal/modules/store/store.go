package store

import (
	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/postgres"
	"github.com/dhanielsales/go-api-template/pkg/redis"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers"
	storerepository "github.com/dhanielsales/go-api-template/internal/modules/store/repository"
	storeservice "github.com/dhanielsales/go-api-template/internal/modules/store/service"
)

func Bootstrap(postgresStorage *postgres.Storage, redisStorage *redis.Storage, httpServer *httputils.HTTPServer, validator *httputils.Validator) {
	repository := storerepository.New(postgresStorage, redisStorage)
	service := storeservice.New(repository)

	controllers.New(service, httpServer, validator)
}
