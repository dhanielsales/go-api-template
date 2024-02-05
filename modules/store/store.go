package store

import (
	"github.com/dhanielsales/golang-scaffold/internal/gql"
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	"github.com/dhanielsales/golang-scaffold/modules/store/application"
	"github.com/dhanielsales/golang-scaffold/modules/store/external"
	"github.com/dhanielsales/golang-scaffold/modules/store/repository"

	"github.com/dhanielsales/golang-scaffold/modules/store/presentation/rest"
)

func Bootstrap(postgres *postgres.Storage, redis *redis.Storage, exampleClient *gql.Client, httpServer *http.HttpServer, validator *http.Validator) {
	repository := repository.New(postgres, redis)
	external := external.New(exampleClient)
	service := application.New(repository, external)

	rest.New(service, httpServer, validator)
}
