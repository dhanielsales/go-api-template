package product

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"
	"github.com/dhanielsales/go-api-template/pkg/apperror"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// PUT /api/v0/product/
//
// @Summary Update one product.
// @Description updates one product by id.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param Product body UpdateProductRequest true "Product to update"
// @Success		200 {object} int64
// @Header		200,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Success 200 {object} int64
// @Router /api/v0/product/{id}/ [put]
func (t *ProductController) UpdateProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return apperror.FromError(err).WithDescription("invalid parameter 'id'")
	}

	var req UpdateProductRequest
	if err := t.validator.DecodeAndValidate(c, &req); err != nil {
		return err
	}

	affected, err := t.service.UpdateProduct(c.Request().Context(), id, product.UpdateProductPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, affected)
}

type UpdateProductRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}
