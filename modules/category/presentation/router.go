package category_http

import "github.com/dhanielsales/golang-scaffold/internal/http"

func SetupRoutes(httpServer *http.HttpServer, controller *CategoryController) {
	router := httpServer.App.Group("/category")

	// add middlewares here

	router.Post("/", controller.create)
	router.Get("/", controller.getAll)
}
