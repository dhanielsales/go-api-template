package store

import (
	"github.com/dhanielsales/golang-scaffold/pkg/httputils"
	"github.com/dhanielsales/golang-scaffold/pkg/postgres"
	"github.com/dhanielsales/golang-scaffold/pkg/redis"

	"github.com/dhanielsales/golang-scaffold/internal/modules/store/presentation/controllers"
	"github.com/dhanielsales/golang-scaffold/internal/modules/store/repository"
	"github.com/dhanielsales/golang-scaffold/internal/modules/store/service"
)

func Bootstrap(pg *postgres.Storage, redis *redis.Storage, httpServer *httputils.HttpServer, validator *httputils.Validator) {
	repository := repository.New(pg, redis)
	service := service.New(repository)

	controllers.New(service, httpServer, validator)
}
