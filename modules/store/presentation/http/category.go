package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	appError "github.com/dhanielsales/golang-scaffold/internal/error"
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/modules/store/application"
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
}

// @Summary Create one category.
// @Description creates one category.
// @Tags Category
// @Accept */*
// @Produce json
// @Param Category body createCategoryRequest true "Category to create"
// @Success 201 {object} int64
// @Router /category [post]
func (t *StoreController) createCategory(c *fiber.Ctx) error {
	var req createCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return t.http.ErrorHandler.Response(c, appError.New(err, appError.BadRequestError, "Malformed request body"))
	}

	if errs := t.validator.Validate(req); len(errs) > 0 && errs[0].Error {
		err := appError.New(nil, appError.BadRequestError, errs[0].Message)
		err.AddDetail("failed_fields", errs)
		return t.http.ErrorHandler.Response(c, err)
	}

	affected, err := t.service.CreateCategory(c.Context(), application.CreateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}

// @Summary Get all categories.
// @Description fetch every category available.
// @Tags Category
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Category
// @Router /category [get]
func (t *StoreController) getManyCategory(c *fiber.Ctx) error {
	categories, err := t.service.GetManyCategory(c.Context(), application.GetManyCategoryParams{
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
// @Router /category/{id} [get]
func (t *StoreController) getOneCategory(c *fiber.Ctx) error {
	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return t.http.ErrorHandler.Response(c, currErr)
	}

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
// @Router /category/{id} [put]
func (t *StoreController) updateCategory(c *fiber.Ctx) error {
	var req updateCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return t.http.ErrorHandler.Response(c, appError.New(err, appError.BadRequestError, "Malformed request body"))
	}

	if errs := t.validator.Validate(req); len(errs) > 0 && errs[0].Error {
		err := appError.New(nil, appError.BadRequestError, errs[0].Message)
		err.AddDetail("failed_fields", errs)
		return t.http.ErrorHandler.Response(c, err)
	}

	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return t.http.ErrorHandler.Response(c, currErr)

	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.UpdateCategory(c.Context(), id, application.UpdateCategoryPayload{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}

// @Summary Delete one category.
// @Description deletes one category by id.
// @Tags Category
// @Accept */*
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} int64
// @Router /category/{id} [delete]
func (t *StoreController) deleteCategory(c *fiber.Ctx) error {
	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return t.http.ErrorHandler.Response(c, currErr)

	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.DeleteCategory(c.Context(), id)
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}
