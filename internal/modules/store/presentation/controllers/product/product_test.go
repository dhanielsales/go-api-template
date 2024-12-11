package product_test

import (
	"reflect"
	"testing"

	"github.com/dhanielsales/go-api-template/internal/modules/store/presentation/controllers/product"
	"github.com/dhanielsales/go-api-template/pkg/httputils"
	"github.com/dhanielsales/go-api-template/pkg/transcriber"

	serviceproduct "github.com/dhanielsales/go-api-template/internal/modules/store/service/product"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

// newProductController is a test helper function that initializes
// a new ProductController with mock implementations of its dependency interfaces.
func newProductController(t *testing.T, expect func(mocks *mocks)) *product.ProductController {
	t.Helper()

	ctrl := gomock.NewController(t)
	productService := serviceproduct.NewMockProductService(ctrl)
	transcrib := transcriber.DefaultTranscriber()
	validator := httputils.NewValidator(transcrib)

	if expect != nil {
		expect(&mocks{productService})
	}

	return product.New(productService, validator)
}

type mocks struct {
	productService *serviceproduct.MockProductService
}

func TestNewController(t *testing.T) {
	t.Parallel()

	type args struct {
		service   serviceproduct.ProductService
		validator httputils.Validator[echo.Context]
	}

	type want struct {
		controller *product.ProductController
	}

	ctrl := gomock.NewController(t)
	productService := serviceproduct.NewMockProductService(ctrl)
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
				service:   productService,
				validator: validator,
			},
			want: &want{
				controller: product.New(productService, validator),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			controller := product.New(tt.args.service, tt.args.validator)

			if !reflect.DeepEqual(controller, tt.want.controller) {
				t.Errorf("controller.New(): %+v; want: %+v", controller, tt.want.controller)
			}
		})
	}
}
