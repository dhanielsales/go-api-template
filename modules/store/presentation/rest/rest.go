package rest

import (
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/modules/store/application"
)

func New(service *application.StoreService, httpServer *http.HttpServer, validator *http.Validator) {
	controller := newController(service, httpServer, validator)

	router := httpServer.App.Group("/api/v0/")
	// Setup middlewares here
	// EX: router.Use(middleware)

	setupCategoryRoutes(router, controller)
	setupProductRoutes(router, controller)
}

type StoreController struct {
	validator *http.Validator
	service   *application.StoreService
	http      *http.HttpServer
}

func newController(service *application.StoreService, http *http.HttpServer, validator *http.Validator) *StoreController {
	return &StoreController{
		validator: validator,
		service:   service,
		http:      http,
	}
}
