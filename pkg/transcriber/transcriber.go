package transcriber

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// Solver is an interface that defines the necessary signature for a validator solver.
type Solver interface {
	Validate(ctx context.Context, val any) error
}

// Transcriber is an interface that defines the necessary signature for a transcriber.
type Transcriber interface {
	DecodeAndValidate(ctx context.Context, source io.Reader, target any) error
}

// transcriber is a struct that holds a solver.
type transcriber struct {
	solver Solver
}

var _ Transcriber = &transcriber{}

// NewTranscriber returns a new instance of Transcriber with the given solver.
func NewTranscriber(solver Solver) *transcriber {
	return &transcriber{
		solver: solver,
	}
}

// DefaultTranscriber returns a new instance of Transcriber with the default solver.
// Use the default solver if you don't need to customize the transcription process.
func DefaultTranscriber() *transcriber {
	return NewTranscriber(defaultSolver())
}

// DecodeAndValidate decodes the source into the target struct pointer and validates it.
func (v *transcriber) DecodeAndValidate(ctx context.Context, source io.Reader, target any) error {
	if target == nil {
		return ErrTargetIsNil
	}

	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return ErrTargetIsNotPointer
	}

	if reflect.TypeOf(target).Elem().Kind() != reflect.Struct {
		return ErrTargetIsNotStruct
	}

	decoder := json.NewDecoder(source)
	if err := decoder.Decode(target); err != nil && err != io.EOF {
		return v.formatDecodeError(err)
	}

	err := v.solver.Validate(ctx, target)
	if err != nil {
		return err
	}

	return nil
}

func (v *transcriber) formatDecodeError(err error) error {
	if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
		return InvalidFieldsErrors{{
			Field:   jsonErr.Field,
			Message: fmt.Sprintf(ErrMessageInvalidFieldType, jsonErr.Field, jsonErr.Type, jsonErr.Value),
		}}
	}

	return err
}
