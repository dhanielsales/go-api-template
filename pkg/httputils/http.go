package httputils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/config/env"
	"github.com/dhanielsales/go-api-template/pkg/conversational"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// HTTPServer represents the HTTP server configuration and lifecycle management
// for the application, built using the Echo framework.
type HTTPServer struct {
	App          *echo.Echo
	port         string
	ErrorHandler *HTTPErrorHandler
}

// New creates and configures a new HTTPServer instance with the provided environment values.
func New(envValues *env.Values) *HTTPServer {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	errorHandler := newErrorHandler()
	app.HTTPErrorHandler = errorHandler.Response

	app.Pre(middleware.AddTrailingSlash()) // Needs to be Pre()
	app.Use(contextMiddleware)
	app.Use(middleware.Recover())
	app.Use(middleware.BodyLimit("2M"))
	app.Use(middleware.Logger())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: envValues.HTTP_ALLOW_ORIGIN,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, conversational.CID_HEADER_KEY},
	}))

	app.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		TargetHeader: conversational.CID_HEADER_KEY,
		Generator:    conversational.NewCID,
		RequestIDHandler: func(c echo.Context, cid string) {
			c.Set(conversational.CID_CONTEXT_KEY.String(), cid)
		},
	}))

	// add health check
	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	// add docs
	if envValues.ENV != "production" {
		app.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return &HTTPServer{
		App:          app,
		port:         envValues.HTTP_PORT,
		ErrorHandler: errorHandler,
	}
}

// Start begins the HTTP server and listens on the configured port.
func (h *HTTPServer) Start() error {
	if err := h.App.Start(":" + h.port); err != nil {
		return fmt.Errorf("error starting http server: %w", err)
	}

	return nil
}

// Cleanup shuts down the HTTP server gracefully.
func (h *HTTPServer) Cleanup(ctx context.Context) error {
	if err := h.App.Shutdown(ctx); err != nil {
		return fmt.Errorf("error closing http server: %w", err)
	}

	return nil
}
