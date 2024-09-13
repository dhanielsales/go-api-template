package transcriber

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type target struct {
		Name string `json:"name" validate:"required"`
	}

	type args struct {
		val any
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Success nil",
			args: args{val: nil},
			want: want{err: nil},
		},
		{
			name: "Success",
			args: args{val: &struct {
				Name string `json:"name" validate:"required"`
			}{Name: "John"}},
			want: want{err: nil},
		},
		{
			name: "Success with '-' json tag",
			args: args{val: &struct {
				Name string `json:"name" validate:"required"`
				Age  int    `json:"-"`
			}{Name: "John"}},
			want: want{err: nil},
		},
		{
			name: "Error required",
			args: args{val: &struct {
				Name string `json:"name" validate:"required"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "name",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldRequired, "name"),
					Criteria: "required",
				},
			}},
		},
		{
			name: "Error validator.InvalidValidationError",
			args: args{val: 123},
			want: want{err: &validator.InvalidValidationError{Type: reflect.TypeOf(123)}},
		},
		{
			name: "Error required non-anonymous struct",
			args: args{val: &target{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "name",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldRequired, "name"),
					Criteria: "required",
				},
			}},
		},
		{
			name: "Error criteria with params",
			args: args{val: &struct {
				Name string `json:"name" validate:"len=5"`
			}{Name: "1234"}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "name",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "name", "len=5", "1234"),
					Criteria: "len=5",
				},
			}},
		},
		{
			name: "Error invalid params",
			args: args{val: &struct {
				Name string `json:"name" validate:"invalid"`
			}{Name: "1234"}},
			want: want{err: fmt.Errorf(ErrUnexpectedPanic, "Undefined validation function 'invalid' on field 'Name'")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			solver := newValidatorSolver(validator.New(validator.WithRequiredStructEnabled()))
			err := solver.Validate(ctx, tt.args.val)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func Test_formatValue(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type args struct {
		val any
	}

	type nested struct {
		Name string `json:"name"`
	}

	type target struct {
		Val   *nested `json:"val"   validate:"required_without=other"`
		Other string  `json:"other"`
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "pointer nil",
			args: args{val: struct {
				Name *string `json:"name" validate:"len=2"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "name",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "name", "len=2", "null"),
					Criteria: "len=2",
				},
			}},
		},
		{
			name: "pointer",
			args: args{val: struct {
				Name *string `json:"name" validate:"len=2"`
			}{
				Name: utils.ToPtr("a"),
			}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "name",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "name", "len=2", "a"),
					Criteria: "len=2",
				},
			}},
		},
		{
			name: "struct nested",
			args: args{val: target{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "struct required",
			args: args{val: struct {
				Val   struct{} `json:"val"   validate:"required_without=other"`
				Other string   `json:"other"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "<object>"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "struct nil",
			args: args{val: struct {
				Val   *struct{} `json:"val"   validate:"required_without=other"`
				Other string    `json:"other"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "slice empty",
			args: args{val: struct {
				Val   []string `json:"val"   validate:"min=1"`
				Other string   `json:"other"`
			}{
				Val: []string{},
			}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "min=1", "<array>"),
					Criteria: "min=1",
				},
			}},
		},
		{
			name: "slice nil",
			args: args{val: struct {
				Val   []string `json:"val"   validate:"required_without=other"`
				Other string   `json:"other"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "map empty",
			args: args{val: struct {
				Val   map[string]string `json:"val"   validate:"required_without=other"`
				Other string            `json:"other"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "map pointer",
			args: args{val: struct {
				Val   *map[string]string `json:"val"   validate:"required_without=other,len=1"`
				Other string             `json:"other"`
			}{
				Val: utils.ToPtr(map[string]string{}),
			}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "len=1", "<object>"),
					Criteria: "len=1",
				},
			}},
		},
		{
			name: "any",
			args: args{val: struct {
				Val   any    `json:"val"   validate:"required_without=other"`
				Other string `json:"other"`
			}{}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
		{
			name: "any pointer",
			args: args{val: struct {
				Val   *any   `json:"val"   validate:"required_without=other"`
				Other string `json:"other"`
			}{
				Val: new(any),
			}},
			want: want{err: InvalidFieldsErrors{
				{
					Field:    "val",
					Message:  fmt.Sprintf(ErrMessageInvalidFieldCriteria, "val", "required_without=other", "null"),
					Criteria: "required_without=other",
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			solver := newValidatorSolver(validator.New(validator.WithRequiredStructEnabled()))
			err := solver.Validate(ctx, tt.args.val)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
