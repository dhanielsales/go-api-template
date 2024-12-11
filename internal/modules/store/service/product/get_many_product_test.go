package product_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	apperror "github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/dhanielsales/go-api-template/pkg/testutils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetManyProduct(t *testing.T) {
	t.Parallel()

	type fields struct {
		service product.ProductService
	}

	type args struct {
		params product.GetManyProductParams
	}

	type expected struct {
		data []*models.Product
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
			name: "Error getting categories - general error",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetManyProduct(gomock.Any(), gomock.Any()).Return(nil, errors.New("error getting categories"))
				}),
			},
			args: &args{
				params: product.GetManyProductParams{},
			},
			expected: &expected{
				err: apperror.FromError(errors.New("error getting categories")).WithDescription("can't process product entity"),
			},
		},
		{
			name: "Success",
			fields: &fields{
				service: newProductService(t, func(mocks *mocks) {
					mocks.repository.EXPECT().GetManyProduct(gomock.Any(), gomock.Any()).Return([]*models.Product{{
						ID:   id,
						Name: "Name",
						Slug: "name",
					}}, nil)
				}),
			},
			args: &args{
				params: product.GetManyProductParams{},
			},
			expected: &expected{
				err: nil,
				data: []*models.Product{{
					ID:   id,
					Name: "Name",
					Slug: "name",
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			affects, err := tt.fields.service.GetManyProduct(ctx, tt.args.params)

			testutils.ErrorEqual(t, tt.expected.err, err)
			assert.Equal(t, tt.expected.data, affects)
		})
	}
}
