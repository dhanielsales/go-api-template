package httputils

import (
	"bytes"

	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorDetail struct {
		Error       bool   `json:"-"`
		FailedField string `json:"failed_field"`
		Tag         string `json:"criteria"`
		Value       any    `json:"value"`
		Message     string `json:"message"`
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

type Validator struct {
	transcrib transcriber.Transcriber
}

func NewValidator(transcrib transcriber.Transcriber) *Validator {
	return &Validator{
		transcrib: transcrib,
	}
}

func (v Validator) DecodeAndValidate(c *fiber.Ctx, target any) error {
	if v.transcrib == nil {
		return apperror.New("transcrib is nil")
	}

	ctx := c.Context()
	source := bytes.NewBuffer(c.Body())

	if err := v.transcrib.DecodeAndValidate(ctx, source, target); err != nil {
		return apperror.FromError(err).WithDescription("invalid validation").WithDetails(err)
	}

	return nil
}
