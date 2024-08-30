package controllers

import (
	"net/http"

	"github.com/dhanielsales/golang-scaffold/internal/modules/store/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func setupCategoryRoutes(r fiber.Router, controller *StoreController) {
	router := r.Group("/category")

	// Setup middlewares here
	// EX: router.Use(middleware)

	// Setup routes here
	router.Post("/", controller.createCategory)
	router.Get("/", controller.getManyCategory)
	router.Get("/:id", controller.getOneCategory)
	router.Put("/:id", controller.updateCategory)
	router.Delete("/:id", controller.deleteCategory)
}

type createCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
	ImageUrl    string `json:"imageUrl" validate:"required,url"`
}

// @Summary Create one category.
// @Description creates one category.
// @Tags Category
// @Accept */*
// @Produce json
// @Param Category body createCategoryRequest true "Category to create"
// @Success 201 {object} int64
// @Router /api/v0/category [post]
func (t *StoreController) createCategory(c *fiber.Ctx) error {
	var req createCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	affected, err := t.service.CreateCategory(c.Context(), service.CreateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(affected)
}

// @Summary Get all categories.
// @Description fetch every category available.
// @Tags Category
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Category
// @Router /api/v0/category [get]
func (t *StoreController) getManyCategory(c *fiber.Ctx) error {
	categories, err := t.service.GetManyCategory(c.Context(), service.GetManyCategoryParams{
		Page:    c.Query("page"),
		PerPage: c.Query("perPage"),
		OrderBy: c.Query("orderBy"),
	})
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.JSON(categories)
}

// @Summary Get one category.
// @Description fetch one category by id.
// @Tags Category
// @Accept */*
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entity.Category
// @Router /api/v0/category/{id} [get]
func (t *StoreController) getOneCategory(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	category, err := t.service.GetCategoryById(c.Context(), id)
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.JSON(category)
}

type updateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}

// @Summary Update one category.
// @Description updates one category by id.
// @Tags Category
// @Accept */*
// @Produce json
// @Param id path string true "Category ID"
// @Param Category body updateCategoryRequest true "Category to update"
// @Success 200 {object} int64
// @Router /api/v0/category/{id} [put]
func (t *StoreController) updateCategory(c *fiber.Ctx) error {
	var req updateCategoryRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.UpdateCategory(c.Context(), id, service.UpdateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(http.StatusOK).JSON(affected)
}

// @Summary Delete one category.
// @Description deletes one category by id.
// @Tags Category
// @Accept */*
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} int64
// @Router /api/v0/category/{id} [delete]
func (t *StoreController) deleteCategory(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.DeleteCategory(c.Context(), id)
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(http.StatusOK).JSON(affected)
}
