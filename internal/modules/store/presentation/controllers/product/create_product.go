package product

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// POST /api/v0/product/
//
// @Summary Create one product.
// @Description creates one product.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		Product body CreateProductRequest true "Product to create"
// @Success		201 {object} int64
// @Header		201,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/product/ [post]
func (ctrl *ProductController) CreateProduct(c echo.Context) error {
	var req CreateProductRequest

	if err := ctrl.validator.DecodeAndValidate(c, req); err != nil {
		return err
	}

	affected, err := ctrl.service.CreateProduct(c.Request().Context(), product.CreateProductPayload{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategotyID:  uuid.MustParse(req.CategoryID),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, affected)
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=50"`
	Description string  `json:"description,omitempty" validate:"min=1,max=300"`
	Price       float64 `json:"price" validate:"required,min=0"`
	CategoryID  string  `json:"category_id" validate:"required,uuid4"`
}
