package httputils

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/contextutils"
	"github.com/labstack/echo/v4"
)

// contextMiddleware wraps an Echo handler function to provide additional context handling.
func contextMiddleware(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(contextValue{ctx})
	}
}

// contextValue extends Echo's Context to retrieve values from both the Echo context and the standard context.
type contextValue struct {
	echo.Context
}

// Get retrieves a value associated with the provided key from the Echo context.
// If the key does not exist in the Echo context, it attempts to retrieve it from the standard request context.
func (ctx contextValue) Get(key string) any {
	val := ctx.Context.Get(key)
	if val != nil {
		return val
	}
	return ctx.Request().Context().Value(key)
}

// Set stores a value in the request context, associating it with the provided key.
func (ctx contextValue) Set(key string, val any) {
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), contextutils.ContextKey(key), val)))
}
