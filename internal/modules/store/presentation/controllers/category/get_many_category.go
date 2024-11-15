package category

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"

	"github.com/labstack/echo/v4"
)

// GET /api/v0/category/
//
// @Summary Get all categories.
// @Description fetch every category available
// @Tags Category
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Success 	200 				{object} []models.Category
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/category/ [get]
func (t *CategoryController) GetManyCategory(c echo.Context) error {
	categories, err := t.service.GetManyCategory(c.Request().Context(), category.GetManyCategoryParams{
		Page:    c.QueryParam("page"),
		PerPage: c.QueryParam("perPage"),
		OrderBy: c.QueryParam("orderBy"),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, categories)
}
