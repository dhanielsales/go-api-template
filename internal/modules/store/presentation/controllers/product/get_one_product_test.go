package product_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/product"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestGetOneCategory(t *testing.T) {
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
	id := uuid.New()

	tests := []struct {
		name     string
		fields   *fields
		args     *args
		expected *expected
	}{
		{
			name: "Error invalid uuid",
			fields: &fields{
				controller: newProductController(t, nil),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil).WithParam("id", "invalid"),
			},
			expected: &expected{
				err: errors.New("error Internal Server Error: err invalid UUID length: 7 - invalid parameter 'id'"),
			},
		},
		{
			name: "Error getting one category",
			fields: &fields{
				controller: newProductController(t, func(mocks *mocks) {
					mocks.productService.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil).WithParam("id", "c7fe4f9f-f3f0-49ef-861b-66072d72c197"),
			},
			expected: &expected{
				err: errors.New("error"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				controller: newProductController(t, func(mocks *mocks) {
					mocks.productService.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{
						ID:    id,
						Name:  "Name",
						Slug:  "name",
						Price: 123,
					}, nil)
				}),
			},
			args: &args{
				c: testutils.NewEchoContext(ctx, nil).WithParam("id", "c7fe4f9f-f3f0-49ef-861b-66072d72c197"),
			},
			expected: &expected{
				payload: testutils.ToByte(&models.Product{
					ID:    id,
					Name:  "Name",
					Slug:  "name",
					Price: 123,
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.fields.controller.GetOneProduct(tt.args.c)

			testutils.ErrorEqual(t, tt.expected.err, err)
			testutils.BytesEqual(t, tt.expected.payload, tt.args.c.Rec.Body.Bytes())
		})
	}
}
