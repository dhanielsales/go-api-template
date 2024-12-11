package product_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/google/uuid"
	"github.com/lib/pq"
	gomock "go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	type fields struct {
		service product.ProductService
	}

	type args struct {
		data product.CreateProductPayload
	}

	type expected struct {
		data int64
		err  error
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
			name: "Error invalid payload",
			fields: &fields{
				service: newProductService(t, nil),
			},
			args: &args{
				data: product.CreateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("name is required")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Error creating product - unique violation",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(0), &pq.Error{Message: "duplicate key value violates unique constraint 'slug'"})
				}),
			},
			args: &args{
				data: product.CreateProductPayload{
					Name:        "Dame",
					Description: "Description",
					Price:       1,
					CategoryID:  id,
				},
			},
			expected: &expected{
				err: apperror.FromError(&pq.Error{Message: "duplicate key value violates unique constraint 'slug'"}).WithDescription("product with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity),
			},
		},
		{
			name: "Error creating product - general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error creating product"))
				}),
			},
			args: &args{
				data: product.CreateProductPayload{
					Name:        "Dame",
					Description: "Description",
					Price:       1,
					CategoryID:  id,
				},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error creating product")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				}),
			},
			args: &args{
				data: product.CreateProductPayload{
					Name:        "Dame",
					Description: "Description",
					Price:       1,
					CategoryID:  id,
				},
			},
			expected: &expected{
				err:  nil,
				data: int64(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.CreateProduct(ctx, tt.args.data)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
