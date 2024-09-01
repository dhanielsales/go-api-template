package httputils

import (
	"time"

	"github.com/dhanielsales/go-api-template/internal/config/env"
	"github.com/dhanielsales/go-api-template/pkg/conversational"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

type HttpServer struct {
	App          *fiber.App
	port         string
	ErrorHandler *HttpErrorHandler
}

func New(env *env.EnvVars) *HttpServer {
	errorHandler := newErrorHandler()

	// create the fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.Response,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// add middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: env.HTTP_ALLOW_ORIGIN,
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
	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	if env.ENV != "production" {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	return &HttpServer{
		App:          app,
		port:         env.HTTP_PORT,
		ErrorHandler: errorHandler,
	}
}

func (h *HttpServer) Start() {
	h.App.Listen(":" + h.port)
}

func (h *HttpServer) Cleanup() {
	h.App.ShutdownWithTimeout(time.Second * 5)
}
