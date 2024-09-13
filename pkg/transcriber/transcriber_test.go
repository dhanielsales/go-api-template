package transcriber_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	io "io"
	"reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/pkg/transcriber"
	"github.com/stretchr/testify/assert"

	gomock "go.uber.org/mock/gomock"
)

// newTranscriberTest is a test helper function that initializes
// a new transcriber with mock implementations of its dependency interfaces.
func newTranscriberTest(t *testing.T, expect func(mocks *mocks)) transcriber.Transcriber {
	t.Helper()
	ctrl := gomock.NewController(t)
	solver := transcriber.NewMockSolver(ctrl)

	if expect != nil {
		expect(&mocks{
			solver: solver,
		})
	}

	return transcriber.NewTranscriber(solver)
}

type mocks struct {
	solver *transcriber.MockSolver
}

func TestNewTranscriber(t *testing.T) {
	t.Parallel()

	type args struct {
		solver transcriber.Solver
	}

	type want struct {
		transcrib transcriber.Transcriber
	}

	solver := transcriber.NewMockSolver(gomock.NewController(t))

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Success",
			args: args{solver: solver},
			want: want{transcrib: transcriber.NewTranscriber(solver)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			transcrib := transcriber.NewTranscriber(tt.args.solver)
			if !reflect.DeepEqual(transcrib, tt.want.transcrib) {
				t.Errorf("transcriber.NewTranscriber(): %+v; want: %+v", transcrib, tt.want.transcrib)
			}
		})
	}
}

func TestDecodeAndValidate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type args struct {
		source io.Reader
		target any
	}

	type fields struct {
		transcrib transcriber.Transcriber
	}

	type want struct {
		err error
	}

	tests := []struct {
		name   string
		fields *fields
		args   *args
		want   want
	}{
		{
			name: "Error: target is nil",
			fields: &fields{
				transcrib: newTranscriberTest(t, nil),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{"key": "value"}`)),
				target: nil,
			},
			want: want{
				err: transcriber.ErrTargetIsNil,
			},
		},
		{
			name: "Error: target is not a pointer",
			fields: &fields{
				transcrib: newTranscriberTest(t, nil),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{"key": "value"}`)),
				target: struct {
					Key string `json:"key"`
				}{},
			},
			want: want{
				err: transcriber.ErrTargetIsNotPointer,
			},
		},
		{
			name: "Error: target is not an struct",
			fields: &fields{
				transcrib: newTranscriberTest(t, nil),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{"key": "value"}`)),
				target: new(string),
			},
			want: want{
				err: transcriber.ErrTargetIsNotStruct,
			},
		},
		{
			name: "Error: decoder.Decode() returns an json.UnmarshalTypeError",
			fields: &fields{
				transcrib: newTranscriberTest(t, nil),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{ "key": 123 }`)),
				target: &struct {
					Key string `json:"key"`
				}{},
			},
			want: want{
				err: transcriber.InvalidFieldsErrors{
					transcriber.InvalidFieldError{
						Field:   "key",
						Message: fmt.Sprintf(transcriber.ErrMessageInvalidFieldType, "key", "string", "number"),
					},
				},
			},
		},
		{
			name: "Error: decoder.Decode() returns an error",
			fields: &fields{
				transcrib: newTranscriberTest(t, nil),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{ "key": "value" `)),
				target: &struct {
					Key string `json:"key"`
				}{},
			},
			want: want{
				err: io.ErrUnexpectedEOF,
			},
		},
		{
			name: "Error: validation",
			fields: &fields{
				transcrib: newTranscriberTest(t, func(mocks *mocks) {
					mocks.solver.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(errors.New("validation error"))
				}),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{"key": "value"}`)),
				target: &struct {
					Key string `json:"key" validate:"number"`
				}{},
			},
			want: want{
				err: errors.New("validation error"),
			},
		},
		{
			name: "Success with io.EOF",
			fields: &fields{
				transcrib: newTranscriberTest(t, func(mocks *mocks) {
					mocks.solver.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
				}),
			},
			args: &args{
				source: bytes.NewReader(nil),
				target: &struct {
					Key string `json:"key"`
				}{},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "Success",
			fields: &fields{
				transcrib: newTranscriberTest(t, func(mocks *mocks) {
					mocks.solver.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
				}),
			},
			args: &args{
				source: bytes.NewReader([]byte(`{"key": "value"}`)),
				target: &struct {
					Key string `json:"key" validate:"required"`
				}{},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.fields.transcrib.DecodeAndValidate(ctx, tt.args.source, tt.args.target)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestDefaultTranscriber(t *testing.T) {
	t.Parallel()
	transcrib := transcriber.DefaultTranscriber()
	ctx := context.Background()

	source := bytes.NewReader([]byte(`{"key": "value"}`))
	target := &struct {
		Key string `json:"key" validate:"required"`
	}{}

	err := transcrib.DecodeAndValidate(ctx, source, target)
	if err != nil {
		t.Errorf("transcrib.DecodeAndValidate(): %v; want: nil", err)
	}
}
