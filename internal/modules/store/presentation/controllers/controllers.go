package controllers

import (
	"github.com/dhanielsales/go-api-template/pkg/httputils"

	storeservice "github.com/dhanielsales/go-api-template/internal/modules/store/service"
)

func New(service *storeservice.StoreService, httpServer *httputils.HTTPServer, validator *httputils.Validator) {
	controller := newController(service, httpServer, validator)

	router := httpServer.App.Group("/api/v0")
	// Setup middlewares here
	// EX: router.Use(middleware)

	setupCategoryRoutes(router, controller)
	setupProductRoutes(router, controller)
}

type StoreController struct {
	validator *httputils.Validator
	service   *storeservice.StoreService
	http      *httputils.HTTPServer
}

func newController(service *storeservice.StoreService, http *httputils.HTTPServer, validator *httputils.Validator) *StoreController {
	return &StoreController{
		validator: validator,
		service:   service,
		http:      http,
	}
}
