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

// @Summary Get many fake data.
// @Description get many fake data.
// @Tags NoDb
// @Accept */*
// @Produce json
// @Success 200 {object} string
// @Router /api/v0/no-db [get]
func (t *StoreController) getManyNoDb(c *fiber.Ctx) error {
	data, err := t.service.GetManyNoDb(c.Context())

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.JSON(data)
}
