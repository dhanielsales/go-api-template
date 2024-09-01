package controllers

import (
	"github.com/dhanielsales/go-api-template/pkg/httputils"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service"
)

func New(service *service.StoreService, httpServer *httputils.HttpServer, validator *httputils.Validator) {
	controller := newController(service, httpServer, validator)

	router := httpServer.App.Group("/api/v0/")
	// Setup middlewares here
	// EX: router.Use(middleware)

	setupCategoryRoutes(router, controller)
	setupProductRoutes(router, controller)
}

type StoreController struct {
	validator *httputils.Validator
	service   *service.StoreService
	http      *httputils.HttpServer
}

func newController(service *service.StoreService, http *httputils.HttpServer, validator *httputils.Validator) *StoreController {
	return &StoreController{
		validator: validator,
		service:   service,
		http:      http,
	}
}
