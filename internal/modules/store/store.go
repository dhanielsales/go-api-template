package store

import (
	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/postgres"
	"github.com/dhanielsales/go-api-template/pkg/redis"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers"
	"github.com/dhanielsales/go-api-template/internal/modules/store/repository"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service"
)

func Bootstrap(pg *postgres.Storage, redis *redis.Storage, httpServer *httputils.HttpServer, validator *httputils.Validator) {
	repository := repository.New(pg, redis)
	service := service.New(repository)

	controllers.New(service, httpServer, validator)
}
