package category_test

import (
	reflect "reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/storages"
	"github.com/dhanielsales/go-api-template/internal/modules/store/storages/category"

	"go.uber.org/mock/gomock"
)

// newCategoryRepository is a test helper function that initializes
// a new CategoryRepository with mock implementations of its dependency interfaces.
func newCategoryRepository(t *testing.T, expect func(mocks *mocks)) *category.CategoryRepository {
	t.Helper()

	ctrl := gomock.NewController(t)
	storage := storages.NewMockStorage(ctrl)
	storage.EXPECT().WithTx(gomock.Any()).Return(storage).AnyTimes()
	if expect != nil {
		expect(&mocks{storage, ctrl})
	}

	return category.New(nil, storage, nil)
}

type mocks struct {
	storage *storages.MockStorage
	ctrl    *gomock.Controller
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		storage storages.Storage
	}

	type want struct {
		repository *category.CategoryRepository
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
				repository: category.New(nil, storage, nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repository := category.New(nil, tt.args.storage, nil)

			if !reflect.DeepEqual(repository, tt.want.repository) {
				t.Errorf("repository.New(): %+v; want: %+v", repository, tt.want.repository)
			}
		})
	}
}
