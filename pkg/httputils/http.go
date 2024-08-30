package httputils

import (
	"time"

	"github.com/dhanielsales/golang-scaffold/pkg/conversational"
	"github.com/dhanielsales/golang-scaffold/pkg/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

type HttpServer struct {
	App          *fiber.App
	port         string
	ErrorHandler *HttpErrorHandler
}

func New(port string, logger log.Logger, swaggerOn bool) *HttpServer {
	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: conversational.CID_HEADER_KEY,
	}))

	app.Use(requestid.New(requestid.Config{
		Header:     conversational.CID_HEADER_KEY,
		ContextKey: conversational.CID_CONTEXT_KEY,
		Generator:  conversational.NewCID,
	}))

	// app.Use(fiberLogger.New(fiberLogger.Config{
	// 	TimeFormat: "2006-01-02 15:04:05",
	// 	Format:     "${time} [${ip}] ${status} - ${latency} ${method} ${path}\n",
	// }))
	app.Use(fiberlogger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	if swaggerOn {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

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
