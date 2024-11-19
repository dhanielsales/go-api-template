package store

import (
	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/redisutils"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	controllerscategory "github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/category"
	controllersproduct "github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/product"

	storagescategory "github.com/dhanielsales/go-api-template/internal/modules/store/storages/category"
	storagesproduct "github.com/dhanielsales/go-api-template/internal/modules/store/storages/product"

	servicecategory "github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	serviceproduct "github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
)

func Bootstrap(sqlStorage *sqlutils.Storage, redisStorage *redisutils.Storage, httpServer *httputils.HTTPServer, validator *httputils.Validator) {
	categoryRepo := storagescategory.NewWithDefaultStorage(sqlStorage, redisStorage)
	productRepo := storagesproduct.NewWithDefaultStorage(sqlStorage)

	categoryService := servicecategory.New(categoryRepo)
	productService := serviceproduct.New(productRepo)

	categoryCtrl := controllerscategory.New(categoryService, validator)
	productCtrl := controllersproduct.New(productService, validator)

	router := httpServer.App.Group("/api/v0")

	categoryCtrl.SetupRoutes(router)
	productCtrl.SetupRoutes(router)
}
