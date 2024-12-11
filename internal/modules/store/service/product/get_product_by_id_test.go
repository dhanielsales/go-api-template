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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetProductByID(t *testing.T) {
	t.Parallel()

	type fields struct {
		service product.ProductService
	}

	type args struct {
		id uuid.UUID
	}

	type expected struct {
		data *models.Product
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
			name: "Error getting product - general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error deleting product"))
				}),
			},
			args: &args{id},
			expected: &expected{
				err: apperror.FromError(errors.New("error deleting product")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Error getting product - not found",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
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
					mocks.repository.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(&models.Product{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}, nil)
				}),
			},
			args: &args{id},
			expected: &expected{
				err: nil,
				data: &models.Product{
					ID:   id,
					Name: "Name",
					Slug: "name",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.GetProductByID(ctx, tt.args.id)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
