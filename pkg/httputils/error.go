package httputils

import (
	"github.com/dhanielsales/go-api-template/pkg/conversational"
	"github.com/dhanielsales/go-api-template/pkg/logger"

	apperror "github.com/dhanielsales/go-api-template/pkg/error"

	"github.com/gofiber/fiber/v2"
)

type HttpErrorHandler struct{}

func newErrorHandler() *HttpErrorHandler {
	return &HttpErrorHandler{}
}

func (h HttpErrorHandler) Response(c *fiber.Ctx, err error) error {
	meta := getMeta(c)
	cid := conversational.GetCIDFromContext(c.Context())

	if err == nil {
		currErr := apperror.New("unknow error")
		logger.Error(
			currErr.Error(),
			logger.LogField("cid", cid),
			logger.LogField("request_meta", meta),
		)
		c.Response().Header.Add(conversational.CID_HEADER_KEY, cid)
		return c.Status(currErr.StatusCode()).JSON(currErr)
	}

	if apperr, ok := err.(*apperror.AppError); ok {
		if apperr.Level == apperror.Warn {
			logger.Warn(
				apperr.Error(),
				logger.LogField("cid", cid),
				logger.LogField("request_meta", meta),
				logger.LogField("stack", apperr.Stack()),
			)
		} else {
			logger.Error(
				apperr.Error(),
				logger.LogField("cid", cid),
				logger.LogField("request_meta", meta),
				logger.LogField("stack", apperr.Stack()),
			)
		}

		c.Response().Header.Add(conversational.CID_HEADER_KEY, cid)
		return c.Status(apperr.StatusCode()).JSON(apperr)
	} else {
		currErr := apperror.FromError(err)
		logger.Error(
			apperr.Error(),
			logger.LogField("cid", cid),
			logger.LogField("request_meta", meta),
			logger.LogField("stack", apperr.Stack()),
		)

		c.Response().Header.Add(conversational.CID_HEADER_KEY, cid)
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