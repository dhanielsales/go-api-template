package transcriber

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/dhanielsales/go-api-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var (
	_ Solver = &validatorSolver{}

	once     sync.Once
	instance *validator.Validate
)

const (
	fieldTag      = "json"
	validationTag = "validate"

	nullValue        = "null"
	emptyObjectValue = "{}"
	objectValue      = "<object>"
	arrayValue       = "<array>"

	ErrUnexpectedPanic = "validator: unexpected panic occurred: %v"
)

// defaultSolver returns a new instance of validatorSolver with the default validator.
// Using a single instance of Validate, it will caches struct info.
func defaultSolver() *validatorSolver {
	if instance == nil {
		once.Do(func() {
			instance = validator.New(validator.WithRequiredStructEnabled())
		})
	}

	return newValidatorSolver(instance)
}

// newValidatorSolver returns a new instance of validatorSolver with the given validator.
func newValidatorSolver(validate *validator.Validate) *validatorSolver {
	validate.SetTagName(validationTag)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		const n = 2
		name := strings.SplitN(fld.Tag.Get(fieldTag), ",", n)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &validatorSolver{
		v:             validate,
		fieldTag:      fieldTag,
		validationTag: validationTag,
	}
}

type validatorSolver struct {
	v             *validator.Validate
	fieldTag      string
	validationTag string
}

func (v *validatorSolver) Validate(ctx context.Context, val any) (err error) {
	if val == nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Error("panic recoved from validator", logger.LogAny("panic", r))
			err = fmt.Errorf(ErrUnexpectedPanic, r)
		}
	}()

	if err = v.v.Struct(val); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			invalidFields := make(InvalidFieldsErrors, 0, len(validationErrs))

			for _, validationErr := range validationErrs {
				fieldKey := v.getFieldKey(validationErr)
				expectation := v.formatExpectation(validationErr)
				var message string
				if expectation == "required" {
					message = fmt.Sprintf(ErrMessageInvalidFieldRequired, fieldKey)
				} else {
					value := v.formatValue(validationErr)
					message = fmt.Sprintf(ErrMessageInvalidFieldCriteria, fieldKey, expectation, value)
				}

				invalidFields = append(invalidFields, InvalidFieldError{
					Field:    fieldKey,
					Message:  message,
					Criteria: expectation,
				})
			}

			return invalidFields
		}

		return err
	}

	return err
}

func (v *validatorSolver) getFieldKey(fieldErr validator.FieldError) string {
	namespaceSlice := strings.Split(fieldErr.Namespace(), ".")
	if len(namespaceSlice) == 1 {
		return namespaceSlice[0]
	}

	return strings.Join(namespaceSlice[1:], ".")
}

func (v *validatorSolver) formatExpectation(fieldErr validator.FieldError) string {
	if fieldErr.Param() != "" {
		return fieldErr.Tag() + "=" + fieldErr.Param()
	}

	return fieldErr.Tag()
}

func (v *validatorSolver) formatValue(fieldErr validator.FieldError) string {
	//nolint:exhaustive // No need to check every possible kind
	switch fieldErr.Kind() {
	case reflect.Ptr:
		if reflect.ValueOf(fieldErr.Value()).IsNil() {
			return nullValue
		}
	case reflect.Struct:
		return objectValue
	case reflect.Slice, reflect.Array:
		if reflect.ValueOf(fieldErr.Value()).IsNil() {
			return nullValue
		}

		return arrayValue
	case reflect.Map:
		if reflect.ValueOf(fieldErr.Value()).IsNil() {
			return nullValue
		}

		return objectValue
	case reflect.Interface:
		if !reflect.ValueOf(fieldErr.Value()).IsValid() {
			return nullValue
		}
	}

	return fmt.Sprintf("%v", fieldErr.Value())
}
