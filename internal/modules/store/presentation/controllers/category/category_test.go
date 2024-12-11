package category_test

import (
	"reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/category"
	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"
	"github.com/labstack/echo/v4"

	servicecategory "github.com/dhanielsales/go-api-template/internal/modules/store/service/category"

	"go.uber.org/mock/gomock"
)

// newCategoryController is a test helper function that initializes
// a new CategoryController with mock implementations of its dependency interfaces.
func newCategoryController(t *testing.T, expect func(mocks *mocks)) *category.CategoryController {
	t.Helper()

	ctrl := gomock.NewController(t)
	categoryService := servicecategory.NewMockCategoryService(ctrl)
	validator := httputils.NewMockValidator[echo.Context](ctrl)

	if expect != nil {
		expect(&mocks{categoryService, validator})
	}

	return category.New(categoryService, validator)
}

type mocks struct {
	categoryService *servicecategory.MockCategoryService
	validator       *httputils.MockValidator[echo.Context]
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		service   servicecategory.CategoryService
		validator httputils.Validator[echo.Context]
	}

	type want struct {
		controller *category.CategoryController
	}

	ctrl := gomock.NewController(t)
	categoryService := servicecategory.NewMockCategoryService(ctrl)
	transcrib := transcriber.DefaultTranscriber()
	validator := httputils.NewValidator(transcrib)

	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "New returns a correct Controller",
			args: &args{
				service:   categoryService,
				validator: validator,
			},
			want: &want{
				controller: category.New(categoryService, validator),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			controller := category.New(tt.args.service, tt.args.validator)

			if !reflect.DeepEqual(controller, tt.want.controller) {
				t.Errorf("controller.New(): %+v; want: %+v", controller, tt.want.controller)
			}
		})
	}
}
