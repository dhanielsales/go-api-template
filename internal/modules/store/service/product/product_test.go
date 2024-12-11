package product_test

import (
	reflect "reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"go.uber.org/mock/gomock"
)

// newProductService is a test helper function that initializes
// a new ProductService with mock implementations of its dependency interfaces.
func newProductService(t *testing.T, expect func(mocks *mocks)) product.ProductService {
	t.Helper()

	ctrl := gomock.NewController(t)
	repository := models.NewMockProductRepository(ctrl)
	sqldbmock := sqlutils.MockSQLDBHelper(ctrl)
	repository.EXPECT().Client().Return(sqldbmock).AnyTimes()
	repository.EXPECT().WithTx(gomock.Any()).Return(repository).AnyTimes()
	if expect != nil {
		expect(&mocks{sqldbmock, repository})
	}

	return product.New(repository)
}

type mocks struct {
	sqldbmock  *sqlutils.MockSQLDB
	repository *models.MockProductRepository
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		repository models.ProductRepository
	}

	type want struct {
		service product.ProductService
	}

	ctrl := gomock.NewController(t)
	repository := models.NewMockProductRepository(ctrl)

	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "New returns a correct Service",
			args: &args{repository},
			want: &want{
				service: product.New(repository),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := product.New(tt.args.repository)

			if !reflect.DeepEqual(service, tt.want.service) {
				t.Errorf("service.New(): %+v; want: %+v", service, tt.want.service)
			}
		})
	}
}
