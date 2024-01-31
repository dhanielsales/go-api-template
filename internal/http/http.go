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

func New(port string, logger log.Logger) *HttpServer {
	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	// app.Use(fiberLogger.New(fiberLogger.Config{
	// 	TimeFormat: "2006-01-02 15:04:05",
	// 	Format:     "${time} [${ip}] ${status} - ${latency} ${method} ${path}\n",
	// }))
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
		ErrorHandler: newErrorHandler(logger),
	}
}

func (h *HttpServer) Start() {
	h.App.Listen(":" + h.port)
}

func (h *HttpServer) Cleanup() {
	h.App.ShutdownWithTimeout(time.Second * 5)
}
