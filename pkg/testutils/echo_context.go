package testutils

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	echo.Context
	Rec *httptest.ResponseRecorder
}

func NewEchoContext(ctx context.Context, body []byte) *EchoContext {
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/uri", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	return &EchoContext{echo.New().NewContext(req, rec), rec}
}

func (c *EchoContext) WithParam(key, value string) *EchoContext {
	c.SetParamNames(key)
	c.SetParamValues(value)

	return c
}
