package category

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GET /api/v0/category/{id}/
//
// @Summary Get one category.
// @Description fetch one category by id.
// @Tags Category
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		id path string true "Category ID"
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Success 	200 				{object} 	models.Category
// @Router /api/v0/category/{id}/ [get]
func (t *CategoryController) GetOneCategory(c echo.Context) error {
	id := uuid.MustParse(c.Param("id")) // TODO check the error
	category, err := t.service.GetCategoryByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, category)
}
