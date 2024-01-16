package store

import (
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	"github.com/dhanielsales/golang-scaffold/internal/redis"

	"github.com/dhanielsales/golang-scaffold/modules/store/application"
	"github.com/dhanielsales/golang-scaffold/modules/store/storage"

	store_http "github.com/dhanielsales/golang-scaffold/modules/store/presentation/http"
)

func Bootstrap(postgres *postgres.Storage, redis *redis.Storage, httpServer *http.HttpServer, validator *http.Validator) {
	storage := storage.NewStorage(postgres, redis)
	service := application.NewService(storage)

	store_http.NewHttp(service, httpServer, validator)
}
