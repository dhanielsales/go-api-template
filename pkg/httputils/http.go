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

type HTTPServer struct {
	App          *fiber.App
	port         string
	ErrorHandler *HTTPErrorHandler
}

func New(envValues *env.Values) *HTTPServer {
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
		AllowOrigins: envValues.HTTP_ALLOW_ORIGIN,
		AllowHeaders: conversational.CID_HEADER_KEY,
	}))

	app.Use(requestid.New(requestid.Config{
		Header:     conversational.CID_HEADER_KEY,
		ContextKey: conversational.CID_CONTEXT_KEY.String(),
		Generator:  conversational.NewCID,
	}))

	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// add docs
	if envValues.ENV != "production" {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	return &HTTPServer{
		App:          app,
		port:         envValues.HTTP_PORT,
		ErrorHandler: errorHandler,
	}
}

func (h *HTTPServer) Start() {
	_ = h.App.Listen(":" + h.port)
}

func (h *HTTPServer) Cleanup() error {
	return h.App.ShutdownWithTimeout(time.Second * 5)
}
