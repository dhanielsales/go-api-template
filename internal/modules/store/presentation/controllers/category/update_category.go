package category

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// PUT /api/v0/category/
//
// @Summary Update one category.
// @Description updates one category by id.
// @Tags Category
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param Category body UpdateCategoryRequest true "Category to update"
// @Success		200 {object} int64
// @Header		200,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Success 200 {object} int64
// @Router /api/v0/category/{id}/ [put]
func (t *CategoryController) UpdateCategory(c echo.Context) error {
	var req UpdateCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return err
	}

	id := uuid.MustParse(c.Param("id"))
	affected, err := t.service.UpdateCategory(c.Request().Context(), id, category.UpdateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, affected)
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}
