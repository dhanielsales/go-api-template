package http

import (
	"github.com/gofiber/fiber/v2"
)

func setupNoDbRoutes(r fiber.Router, controller *StoreController) {
	router := r.Group("/no-db")

	// Setup middlewares here
	// EX: router.Use(middleware)

	// Setup routes here
	router.Get("/", controller.getManyNoDb)
}

func (t *StoreController) getManyNoDb(c *fiber.Ctx) error {
	data, err := t.service.GetManyNoDb(c.Context())

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.JSON(data)
}
