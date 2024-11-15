package product

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/product"

	"github.com/labstack/echo/v4"
)

// GET /api/v0/product/
//
// @Summary Get all categories.
// @Description fetch every product available.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Success 	200 				{object} []models.Product
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/product/ [get]
func (t *ProductController) GetManyProduct(c echo.Context) error {
	categories, err := t.service.GetManyProduct(c.Request().Context(), product.GetManyProductParams{
		Page:    c.QueryParam("page"),
		PerPage: c.QueryParam("perPage"),
		OrderBy: c.QueryParam("orderBy"),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, categories)
}
