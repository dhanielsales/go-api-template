package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dhanielsales/golang-scaffold/config/log"

	appError "github.com/dhanielsales/golang-scaffold/internal/error"
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
		currErr := appError.New(nil, appError.ServerError, "Unknow error")
		h.logger.Error(log.Params{Message: currErr.Description, Error: currErr, Meta: meta})
		return c.Status(currErr.Name.Status()).JSON(currErr)
	}

	appErr, ok := appError.Is(err)
	if ok {
		h.logger.Error(log.Params{Message: appErr.Description, Error: appErr, Meta: meta})
		return c.Status(appErr.Name.Status()).JSON(appErr)
	} else {
		currErr := appError.New(err, appError.ServerError, "Internal server error")
		h.logger.Error(log.Params{Message: currErr.Description, Error: currErr, Meta: meta})
		return c.Status(currErr.Name.Status()).JSON(currErr)
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
