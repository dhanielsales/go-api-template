package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type HttpServer struct {
	App  *fiber.App
	port string
}

func Bootstrap(port string) *HttpServer {
	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	return &HttpServer{
		App:  app,
		port: port,
	}
}

func (h *HttpServer) Start() {
	h.App.Listen("0.0.0.0:" + h.port)
}

func (h *HttpServer) Cleanup() {
	h.App.Shutdown()
}
