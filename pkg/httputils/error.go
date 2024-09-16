package httputils

import (
	"github.com/dhanielsales/go-api-template/pkg/conversational"
	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/labstack/echo/v4"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
)

type HTTPErrorHandler struct{}

func newErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{}
}

func (h HTTPErrorHandler) Response(err error, c echo.Context) {
	meta := getMeta(c)
	cid := conversational.GetCIDFromContext(c.Request().Context())

	if err == nil {
		currErr := apperror.New("unknow error")
		logger.Error(
			currErr.Error(),
			logger.LogString("cid", cid),
			logger.LogAny("request_meta", meta),
		)
		c.Logger().Error(currErr)
		c.Response().Header().Add(conversational.CID_HEADER_KEY, cid)
		_ = c.JSON(currErr.StatusCode(), currErr)
		return
	}

	if echoErr, ok := err.(*echo.HTTPError); ok {
		logger.Error(
			echoErr.Error(),
			logger.LogString("cid", cid),
			logger.LogAny("request_meta", meta),
		)
		c.Logger().Error(echoErr)
		c.Response().Header().Add(conversational.CID_HEADER_KEY, cid)
		_ = c.JSON(echoErr.Code, echoErr)
		return
	}

	if apperr, ok := err.(*apperror.AppError); ok {
		if apperr.Level == apperror.Warn {
			logger.Warn(
				apperr.Error(),
				logger.LogString("cid", cid),
				logger.LogAny("request_meta", meta),
				logger.LogString("stack", apperr.Stack()),
			)
			c.Logger().Warn(apperr)
		} else {
			logger.Error(
				apperr.Error(),
				logger.LogString("cid", cid),
				logger.LogAny("request_meta", meta),
				logger.LogString("stack", apperr.Stack()),
			)
			c.Logger().Error(apperr)
		}

		c.Response().Header().Add(conversational.CID_HEADER_KEY, cid)
		_ = c.JSON(apperr.StatusCode(), apperr)
		return
	} else {
		currErr := apperror.FromError(err)
		logger.Error(
			currErr.Error(),
			logger.LogString("cid", cid),
			logger.LogAny("request_meta", meta),
			logger.LogString("stack", currErr.Stack()),
		)
		c.Logger().Error(currErr)
		c.Response().Header().Add(conversational.CID_HEADER_KEY, cid)
		_ = c.JSON(currErr.StatusCode(), currErr)
		return
	}
}

func getMeta(c echo.Context) map[string]any {
	return map[string]any{
		"request_uri":        c.Request().URL,
		"request_method":     c.Request().Method,
		"request_ip":         c.RealIP(),
		"request_user_agent": c.Request().Header.Get("User-Agent"),
	}
}
