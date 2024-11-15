package product

import (
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	"github.com/dhanielsales/go-api-template/pkg/httputils"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	validator *httputils.Validator
	service   product.ProductService
}

func New(service product.ProductService, validator *httputils.Validator) *ProductController {
	return &ProductController{
		validator: validator,
		service:   service,
	}
}

func (ctrl *ProductController) SetupRoutes(app *echo.Group) {
	router := app.Group("/product")

	// Setup middlewares here
	// EX: router.Use(middleware)

	router.POST("/", ctrl.CreateProduct)
	router.GET("/", ctrl.GetManyProduct)
	router.GET("/:id", ctrl.GetOneProduct)
	router.PUT("/:id", ctrl.UpdateProduct)
	router.DELETE("/:id", ctrl.DeleteProduct)
}
