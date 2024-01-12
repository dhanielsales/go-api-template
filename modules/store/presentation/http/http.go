package http

import (
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/modules/store/application"
)

func NewHttp(service *application.StoreService, httpServer *http.HttpServer, validator *http.Validator) {
	controller := newController(service, validator)

	router := httpServer.App.Group("/api/v0/")
	// Setup middlewares here
	// EX: router.Use(middleware)

	setupCategoryRoutes(router, controller)
	setupProductRoutes(router, controller)
}

type StoreController struct {
	service   *application.StoreService
	validator *http.Validator
}

func newController(service *application.StoreService, validator *http.Validator) *StoreController {
	return &StoreController{
		service:   service,
		validator: validator,
	}
}
