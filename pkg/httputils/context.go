package httputils

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/contextutils"
	"github.com/labstack/echo/v4"
)

func contextMiddleware(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(contextValue{ctx})
	}
}

type contextValue struct {
	echo.Context
}

func (ctx contextValue) Get(key string) any {
	val := ctx.Context.Get(key)
	if val != nil {
		return val
	}
	return ctx.Request().Context().Value(key)
}

func (ctx contextValue) Set(key string, val any) {
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), contextutils.ContextKey(key), val)))
}
