package httputils

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dhanielsales/golang-scaffold/pkg/log"

	apperror "github.com/dhanielsales/golang-scaffold/pkg/error"
)

type HttpErrorHandler struct {
	logger log.Logger
}

func newErrorHandler(logger log.Logger) *HttpErrorHandler {
	return &HttpErrorHandler{
		logger: logger,
	}
}

func (h HttpErrorHandler) Response(c *fiber.Ctx, err error) error {
	meta := getMeta(c)

	if err == nil {
		currErr := apperror.New("Unknow error")
		h.logger.Error(log.Params{Message: currErr.Description, Error: currErr, Meta: meta}) // Add CID here
		return c.Status(currErr.StatusCode()).JSON(currErr)
	}

	if appErr, ok := err.(*apperror.AppError); ok {
		h.logger.Error(log.Params{Message: appErr.Description, Error: appErr, Meta: meta}) // Add CID here
		return c.Status(appErr.StatusCode()).JSON(appErr)
	} else {
		currErr := apperror.FromError(err)
		h.logger.Error(log.Params{Message: currErr.Description, Error: currErr, Meta: meta}) // Add CID here
		return c.Status(currErr.StatusCode()).JSON(currErr)
	}
}

func getMeta(c *fiber.Ctx) map[string]any {
	return map[string]any{
		"request_uri":        c.OriginalURL(),
		"request_method":     c.Method(),
		"request_ip":         c.IP(),
		"request_user_agent": c.Get("User-Agent"),
	}
}
