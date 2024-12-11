package category_test

import (
	reflect "reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/models"
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	"github.com/dhanielsales/go-api-template/pkg/sqlutils"

	"go.uber.org/mock/gomock"
)

// newCategoryService is a test helper function that initializes
// a new CategoryService with mock implementations of its dependency interfaces.
func newCategoryService(t *testing.T, expect func(mocks *mocks)) category.CategoryService {
	t.Helper()

	ctrl := gomock.NewController(t)
	categoryRepository := models.NewMockCategoryRepository(ctrl)
	categoryRepository.EXPECT().Client().Return(sqlutils.MockSQLDBHelper(ctrl)).AnyTimes()
	categoryRepository.EXPECT().WithTx(gomock.Any()).Return(categoryRepository).AnyTimes()
	if expect != nil {
		expect(&mocks{categoryRepository})
	}

	return category.New(categoryRepository)
}

type mocks struct {
	repository *models.MockCategoryRepository
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		repository models.CategoryRepository
	}

	type want struct {
		service category.CategoryService
	}

	ctrl := gomock.NewController(t)
	categoryRepository := models.NewMockCategoryRepository(ctrl)

	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "New returns a correct Service",
			args: &args{categoryRepository},
			want: &want{
				service: category.New(categoryRepository),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := category.New(tt.args.repository)

			if !reflect.DeepEqual(service, tt.want.service) {
				t.Errorf("service.New(): %+v; want: %+v", service, tt.want.service)
			}
		})
	}
}
