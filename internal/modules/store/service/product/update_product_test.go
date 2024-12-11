package product_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"
	"github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateProduct(t *testing.T) {
	t.Parallel()

	type fields struct {
		service product.ProductService
	}

	type args struct {
		id     uuid.UUID
		params product.UpdateProductPayload
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
			name: "Error update product - get product to update",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error getting product"))
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error getting product")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Error update product - not found product",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(sql.ErrNoRows).WithDescription("product not found").WithStatusCode(http.StatusNotFound),
			},
		},
		{
			name: "Error update product - update product unique violation",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{ID: id}, nil)
					mocks.repository.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), &pq.Error{Message: "duplicate key value violates unique constraint 'slug'"})
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(&pq.Error{Message: "duplicate key value violates unique constraint 'slug'"}).WithDescription("product with 'slug' already exists").WithStatusCode(http.StatusUnprocessableEntity),
			},
		},
		{
			name: "Error update product - update product general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{ID: id}, nil)
					mocks.repository.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error updating product"))
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error updating product")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Error update product - update product general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{ID: id}, nil)
					mocks.repository.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error updating product"))
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error updating product")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{ID: id}, nil)
					mocks.repository.EXPECT().UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(1), nil)
				}),
			},
			args: &args{
				params: product.UpdateProductPayload{},
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
			affects, err := tt.fields.service.UpdateProduct(ctx, tt.args.id, tt.args.params)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
