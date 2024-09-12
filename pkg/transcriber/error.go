package transcriber

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrTargetIsNotPointer = errors.New("validator: target must be a pointer")
	ErrTargetIsNotStruct  = errors.New("validator: target must be a pointer to a struct")
	ErrTargetIsNil        = errors.New("validator: target cannot be nil")
	ErrPayloadMaxSize     = errors.New("validator: the source's size exceeds the limit")
)

const (
	ErrMessageInvalidField         = "invalid field error: %s"
	ErrMessageInvalidFieldType     = "invalid field type on '%v', expected type '%v', received value '%v'"
	ErrMessageInvalidFieldCriteria = "invalid field criteria on '%v', expected criteria '%v', received value '%v'"
	ErrMessageInvalidFieldRequired = "'%v' field is required"
)

// InvalidFieldError is a struct that holds the field name and the error message.
type InvalidFieldError struct {
	Field    string `json:"field"`
	Message  string `json:"message"`
	Criteria string `json:"-"`
}

func (e InvalidFieldError) Error() string {
	return fmt.Sprintf(ErrMessageInvalidField, e.Message)
}

type InvalidFieldsErrors []InvalidFieldError

func (errs InvalidFieldsErrors) Error() string {
	buff := bytes.NewBufferString("")

	for i := range errs {
		buff.WriteString(errs[i].Error())
		if i < len(errs)-1 {
			buff.WriteString(", ")
		}
	}

	return strings.TrimSpace(buff.String())
}
