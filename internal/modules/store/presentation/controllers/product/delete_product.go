package product

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/pkg/apperror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// DELETE /api/v0/product/{id}/
//
// @Summary Delete one product.
// @Description deletes one product by id.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		id path string true "Product ID"
// @Success 	200 {object} int64
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/product/{id}/ [delete]
func (t *ProductController) DeleteProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return apperror.FromError(err).WithDescription("invalid parameter 'id'")
	}

	affected, err := t.service.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, affected)
}
