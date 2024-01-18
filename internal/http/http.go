package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/dhanielsales/golang-scaffold/config/log"
)

type HttpServer struct {
	App          *fiber.App
	port         string
	ErrorHandler *HttpErrorHandler
}

func Bootstrap(port string, logger log.Logger) *HttpServer {
	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(fiberLogger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	return &HttpServer{
		App:          app,
		port:         port,
		ErrorHandler: NewErrorHandler(logger),
	}
}

func (h *HttpServer) Start() {
	h.App.Listen("0.0.0.0:" + h.port)
}

func (h *HttpServer) Cleanup() {
	h.App.ShutdownWithTimeout(time.Second * 5)
}
