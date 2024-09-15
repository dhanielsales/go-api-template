package httputils

import (
	"context"
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/config/env"
	"github.com/dhanielsales/go-api-template/pkg/conversational"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	App          *echo.Echo
	port         string
	ErrorHandler *HTTPErrorHandler
}

func New(envValues *env.Values) *HTTPServer {
	app := echo.New()

	errorHandler := newErrorHandler()
	app.HTTPErrorHandler = errorHandler.Response

	app.Pre(middleware.AddTrailingSlash()) // Needs to be Pre
	app.Use(contextMiddleware)
	app.Use(middleware.Recover())
	app.Use(middleware.BodyLimit("2M"))
	app.Use(middleware.Logger())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{}, // envValues.HTTP_ALLOW_ORIGIN
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

func (h *HTTPServer) Start() {
	_ = h.App.Start(":" + h.port)
}

func (h *HTTPServer) Cleanup(ctx context.Context) error {
	return h.App.Shutdown(ctx)
}
