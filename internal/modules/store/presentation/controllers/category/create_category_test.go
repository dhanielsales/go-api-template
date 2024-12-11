package category_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/category"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"go.uber.org/mock/gomock"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		controller *category.CategoryController
	}

	type args struct {
		c *testutils.EchoContext
	}

	type expected struct {
		err     error
		payload []byte
	}

	ctx := context.Background()

	tests := []struct {
		name     string
		fields   *fields
		args     *args
		expected *expected
	}{
		{
			name: "Error invalid payload",
			fields: &fields{
				controller: newCategoryController(t, func(mocks *mocks) {
					mocks.validator.EXPECT().DecodeAndValidate(gomock.Any(), gomock.Any()).Return(errors.New("error"))
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil),
			},
			expected: &expected{
				err: errors.New("error"),
			},
		},
		{
			name: "Error creating category",
			fields: &fields{
				controller: newCategoryController(t, func(mocks *mocks) {
					mocks.validator.EXPECT().DecodeAndValidate(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryService.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error"))
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, []byte(`{ "name": "name", "description": "description" }`)),
			},
			expected: &expected{
				err: errors.New("error"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				controller: newCategoryController(t, func(mocks *mocks) {
					mocks.validator.EXPECT().DecodeAndValidate(gomock.Any(), gomock.Any()).Return(nil)
					mocks.categoryService.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, []byte(`{ "name": "name", "description": "description" }`)),
			},
			expected: &expected{
				err:     nil,
				payload: testutils.Int64ToByte(int64(1)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.fields.controller.CreateCategory(tt.args.c)

			testutils.ErrorEqual(t, tt.expected.err, err)
			testutils.BytesEqual(t, tt.expected.payload, tt.args.c.Rec.Body.Bytes())
			testutils.BytesEqual(t, tt.expected.payload, tt.args.c.Rec.Body.Bytes())
		})
	}
}
