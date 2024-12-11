package product_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/product"
	"github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"go.uber.org/mock/gomock"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	type fields struct {
		controller *product.ProductController
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
				controller: newProductController(t, func(mocks *mocks) {}),
			},
			args: &args{
				c: testutils.
					NewEchoContext(ctx, []byte(`{`)),
			},
			expected: &expected{
				err: apperror.FromError(io.ErrUnexpectedEOF).WithDescription("invalid validation").WithDetails(io.ErrUnexpectedEOF),
			},
		},
		{
			name: "Error creating product",
			fields: &fields{
				controller: newProductController(t, func(mocks *mocks) {
					mocks.productService.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error"))
				}),
			},
			args: &args{
				c: testutils.
					NewEchoContext(ctx, []byte(`{ "name": "name", "description": "description", "price": 123, "category_id": "e80472ca-6167-4eca-8b51-eb69c04d8c50" }`)),
			},
			expected: &expected{
				err: errors.New("error"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				controller: newProductController(t, func(mocks *mocks) {
					mocks.productService.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				}),
			},
			args: &args{
				c: testutils.
					NewEchoContext(ctx, []byte(`{ "name": "name", "description": "description", "price": 123, "category_id": "e80472ca-6167-4eca-8b51-eb69c04d8c50" }`)),
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
			err := tt.fields.controller.CreateProduct(tt.args.c)

			testutils.ErrorEqual(t, tt.expected.err, err)
			testutils.BytesEqual(t, tt.expected.payload, tt.args.c.Rec.Body.Bytes())
		})
	}
}
