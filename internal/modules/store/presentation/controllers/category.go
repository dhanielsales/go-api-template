package controllers

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func setupCategoryRoutes(r *echo.Group, controller *StoreController) {
	router := r.Group("/category")

	// Setup middlewares here
	// EX: router.Use(middleware)

	// Setup routes here
	router.POST("/", controller.createCategory)
	router.GET("/", controller.getManyCategory)
	router.GET("/:id", controller.getOneCategory)
	router.PUT("/:id", controller.updateCategory)
	router.DELETE("/:id", controller.deleteCategory)
}

type createCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}

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
// @Param Category body createCategoryRequest true "Category to create"
// @Success		201 {object} int64
// @Header		201,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/category/ [post]
func (t *StoreController) createCategory(c echo.Context) error {
	var req createCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return err
	}

	affected, err := t.service.CreateCategory(c.Request().Context(), service.CreateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, affected)
}

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
func (t *StoreController) getManyCategory(c echo.Context) error {
	categories, err := t.service.GetManyCategory(c.Request().Context(), service.GetManyCategoryParams{
		Page:    c.QueryParam("page"),
		PerPage: c.QueryParam("perPage"),
		OrderBy: c.QueryParam("orderBy"),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, categories)
}

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
func (t *StoreController) getOneCategory(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	category, err := t.service.GetCategoryByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, category)
}

type updateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}

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
// @Param Category body updateCategoryRequest true "Category to update"
// @Success		200 {object} int64
// @Header		200,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	apperror.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Success 200 {object} int64
// @Router /api/v0/category/{id}/ [put]
func (t *StoreController) updateCategory(c echo.Context) error {
	var req updateCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return err
	}

	id := uuid.MustParse(c.Param("id"))
	affected, err := t.service.UpdateCategory(c.Request().Context(), id, service.UpdateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, affected)
}

// DELETE /api/v0/category/{id}/
//
// @Summary Delete one category.
// @Description deletes one category by id.
// @Tags Category
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		id path string true "Category ID"
// @Success 	200 {object} int64
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	apperror.AppError	"Internal Server Error."
// @Router /api/v0/category/{id}/ [delete]
func (t *StoreController) deleteCategory(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	affected, err := t.service.DeleteCategory(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, affected)
}
