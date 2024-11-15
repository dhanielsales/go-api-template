package category

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service/category"

	"github.com/labstack/echo/v4"
)

// POST /api/v0/category/
//
// @Summary Create one category.
// @Description creates one category.
// @Tags Category
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param Category body CreateCategoryRequest true "Category to create"
// @Success		201 {object} int64
// @Header		201,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/category/ [post]
func (t *CategoryController) CreateCategory(c echo.Context) error {
	var req CreateCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return err
	}

	affected, err := t.service.CreateCategory(c.Request().Context(), category.CreateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, affected)
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}
