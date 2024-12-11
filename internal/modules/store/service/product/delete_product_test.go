package product_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	type fields struct {
		service product.ProductService
	}

	type args struct {
		id uuid.UUID
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
			name: "Error deleting product - general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("error deleting product"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error deleting product")).WithDescription("can't process product entity").WithStatusCode(http.StatusUnprocessableEntity),
			},
		},
		{
			name: "Error deleting product - not found",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(int64(0), sql.ErrNoRows)
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(sql.ErrNoRows).WithDescription("product not found").WithStatusCode(http.StatusNotFound),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				}),
			},
			args: &args{id},
			expected: &expected{
				err:  nil,
				data: int64(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.DeleteProduct(ctx, tt.args.id)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
