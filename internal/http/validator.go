package http

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
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
	validator *validator.Validate
}

func NewValidator(validate *validator.Validate) *Validator {
	return &Validator{
		validator: validate,
	}
}

func (v Validator) Validate(data any) []ErrorDetail {
	validationErrors := []ErrorDetail{}

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorDetail

			if field, ok := reflect.TypeOf(data).FieldByName(err.Field()); ok {
				elem.FailedField = strings.Split(field.Tag.Get("json"), ",")[0]
			} else {
				elem.FailedField = err.Field()
			}

			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			elem.Message = fmt.Sprintf("field '%s' needs to implement '%s'", elem.FailedField, elem.Tag)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (v Validator) ValidateField(field any, fieldName, tag string) *ErrorDetail {
	err := v.validator.Var(field, tag)
	if err != nil {
		return &ErrorDetail{
			FailedField: fieldName,
			Error:       true,
			Tag:         tag,
			Value:       field,
			Message:     err.Error(),
		}
	}

	return nil
}

func setupCustomValidator(validate *validator.Validate) {
	validate.RegisterValidation("sorting", func(fl validator.FieldLevel) bool {
		orderString := fl.Field().String()

		if orderString == "" {
			return true
		}

		parts := strings.Split(orderString, ":")
		if len(parts) != 2 {
			return false
		}

		fieldName := strings.TrimSpace(parts[0])
		direction := strings.TrimSpace(parts[1])

		if direction != "asc" && direction != "desc" {
			return false
		}

		if fieldName == "" {
			return false
		}

		return true
	})
}
