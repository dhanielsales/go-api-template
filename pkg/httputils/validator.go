package httputils

import (
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"

	"github.com/labstack/echo/v4"
)

// ErrorDetail represents detailed information about an error that occurred during validation.
type ErrorDetail struct {
	Error       bool   `json:"-"`
	FailedField string `json:"failed_field"`
	Tag         string `json:"criteria"`
	Value       any    `json:"value"`
	Message     string `json:"message"`
}

// GlobalErrorHandlerResp represents the response structure for a global error handler.
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Validator handles the decoding and validation of incoming request data.
type Validator struct {
	transcrib transcriber.Transcriber
}

// NewValidator creates a new instance of Validator with the provided transcriber.
func NewValidator(transcrib transcriber.Transcriber) *Validator {
	return &Validator{
		transcrib: transcrib,
	}
}

// DecodeAndValidate decodes the request body and validates the target object.
func (v Validator) DecodeAndValidate(c echo.Context, target any) error {
	if v.transcrib == nil {
		return apperror.New("transcrib is nil")
	}

	ctx := c.Request().Context()
	if err := v.transcrib.DecodeAndValidate(ctx, c.Request().Body, target); err != nil {
		return apperror.FromError(err).WithDescription("invalid validation").WithDetails(err)
	}

	return nil
}
