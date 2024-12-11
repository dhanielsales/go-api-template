package product_test

import (
	reflect "reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages/product"

	"go.uber.org/mock/gomock"
)

// newProductRepository is a test helper function that initializes
// a new ProductRepository with mock implementations of its dependency interfaces.
func newProductRepository(t *testing.T, expect func(mocks *mocks)) *product.ProductRepository {
	t.Helper()

	ctrl := gomock.NewController(t)
	storage := storages.NewMockStorage(ctrl)
	storage.EXPECT().WithTx(gomock.Any()).Return(storage).AnyTimes()
	if expect != nil {
		expect(&mocks{storage})
	}

	return product.New(nil, storage)
}

type mocks struct {
	storage *storages.MockStorage
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		storage storages.Storage
	}

	type want struct {
		repository *product.ProductRepository
	}

	ctrl := gomock.NewController(t)
	storage := storages.NewMockStorage(ctrl)

	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "New returns a correct repository",
			args: &args{storage},
			want: &want{
				repository: product.New(nil, storage),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repository := product.New(nil, tt.args.storage)

			if !reflect.DeepEqual(repository, tt.want.repository) {
				t.Errorf("repository.New(): %+v; want: %+v", repository, tt.want.repository)
			}
		})
	}
}
