package category

import (
	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"
	"github.com/dhanielsales/go-api-template/pkg/httputils"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	validator *httputils.Validator
	service   category.CategoryService
}

func New(service category.CategoryService, validator *httputils.Validator) *CategoryController {
	return &CategoryController{
		validator: validator,
		service:   service,
	}
}

func (ctrl *CategoryController) SetupRoutes(app *echo.Group) {
	router := app.Group("/category")

	// Setup middlewares here
	// EX: router.Use(middleware)

	router.POST("/", ctrl.CreateCategory)
	router.GET("/", ctrl.GetManyCategory)
	router.GET("/:id", ctrl.GetOneCategory)
	router.PUT("/:id", ctrl.UpdateCategory)
	router.DELETE("/:id", ctrl.DeleteCategory)
}
