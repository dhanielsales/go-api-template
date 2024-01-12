package http

import (
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/modules/store/application"
)

func NewHttp(service *application.StoreService, httpServer *http.HttpServer, validator *http.Validator) {
	controller := newController(service, validator)

	setupCategoryRoutes(httpServer, controller)
	setupProductRoutes(httpServer, controller)
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
