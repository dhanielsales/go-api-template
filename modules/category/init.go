package category

import (
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/internal/postgres"
	category_application "github.com/dhanielsales/golang-scaffold/modules/category/application"
	category_http "github.com/dhanielsales/golang-scaffold/modules/category/presentation"
	category_storage "github.com/dhanielsales/golang-scaffold/modules/category/storage"
)

func Bootstrap(postgresDb *postgres.Storage, httpServer *http.HttpServer) {
	storage := category_storage.NewCategoryStorage(postgresDb)
	service := category_application.NewCategoryService(storage)
	controller := category_http.NewCategoryController(service)
	category_http.SetupRoutes(httpServer, controller)
}
