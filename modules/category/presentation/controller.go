package category_http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	category_application "github.com/dhanielsales/golang-scaffold/modules/category/application"
)

type CategoryController struct {
	service *category_application.CategoryService
}

func NewCategoryController(service *category_application.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

type createCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"max=300"`
}

// @Summary Create one category.
// @Description creates one category.
// @Tags Category
// @Accept */*
// @Produce json
// @Param Category body createCategoryRequest true "Category to create"
// @Success 201
// @Router /category [post]
func (t *CategoryController) create(c *fiber.Ctx) error {
	var req createCategoryRequest
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := t.service.Create(c.Context(), category_application.CreateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

// @Summary Get all categories.
// @Description fetch every category available.
// @Tags Category
// @Accept */*
// @Produce json
// @Success 200 {object} []category_entity.Category
// @Router /category [get]
func (t *CategoryController) getAll(c *fiber.Ctx) error {
	categories, err := t.service.GetAll(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get categories",
		})
	}

	return c.JSON(categories)
}
