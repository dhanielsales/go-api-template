package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	appError "github.com/dhanielsales/golang-scaffold/internal/error"
	"github.com/dhanielsales/golang-scaffold/internal/http"
	"github.com/dhanielsales/golang-scaffold/modules/store/application"
)

func setupProductRoutes(r fiber.Router, controller *StoreController) {
	router := r.Group("/product")

	// Setup middlewares here
	// EX: router.Use(middleware)

	// Setup routes here
	router.Post("/", controller.createProduct)
	router.Get("/", controller.getManyProduct)
	router.Get("/:id", controller.getOneProduct)
	router.Put("/:id", controller.updateProduct)
	router.Delete("/:id", controller.deleteProduct)
}

type createProductRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=50"`
	Description string  `json:"description,omitempty" validate:"min=1,max=300"`
	Price       float64 `json:"price" validate:"required,min=0"`
	CategoryID  string  `json:"category_id" validate:"required,uuid4"`
}

// @Summary Create one product.
// @Description creates one product.
// @Tags Product
// @Accept */*
// @Produce json
// @Param Product body createProductRequest true "Product to create"
// @Success 201 {object} int64
// @Router /product [post]
func (t *StoreController) createProduct(c *fiber.Ctx) error {
	var req createProductRequest

	if err := c.BodyParser(&req); err != nil {
		return http.ResponseError(c, appError.New(err, appError.BadRequestError, "Malformed request body"))
	}

	if errs := t.validator.Validate(req); len(errs) > 0 && errs[0].Error {
		err := appError.New(nil, appError.BadRequestError, errs[0].Message)
		err.AddDetail("failed_fields", errs)
		return http.ResponseError(c, err)
	}

	affected, err := t.service.CreateProduct(c.Context(), application.CreateProductPayload{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategotyID:  uuid.MustParse(req.CategoryID),
	})

	if err != nil {
		return http.ResponseError(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}

// @Summary Get all categories.
// @Description fetch every product available.
// @Tags Product
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Product
// @Router /product [get]
func (t *StoreController) getManyProduct(c *fiber.Ctx) error {
	categories, err := t.service.GetManyProduct(c.Context(), application.GetManyProductParams{
		Page:    c.Query("page"),
		PerPage: c.Query("perPage"),
		OrderBy: c.Query("orderBy"),
	})
	if err != nil {
		return http.ResponseError(c, err)
	}

	return c.JSON(categories)
}

// @Summary Get one product.
// @Description fetch one product by id.
// @Tags Product
// @Accept */*
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Router /product/{id} [get]
func (t *StoreController) getOneProduct(c *fiber.Ctx) error {
	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return http.ResponseError(c, currErr)
	}

	id := uuid.MustParse(c.Params("id"))
	product, err := t.service.GetProductById(c.Context(), id)
	if err != nil {
		return http.ResponseError(c, err)
	}

	return c.JSON(product)
}

type updateProductRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}

// @Summary Update one product.
// @Description updates one product by id.
// @Tags Product
// @Accept */*
// @Produce json
// @Param id path string true "Product ID"
// @Param Product body updateProductRequest true "Product to update"
// @Success 200 {object} int64
// @Router /product/{id} [put]
func (t *StoreController) updateProduct(c *fiber.Ctx) error {
	var req updateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return http.ResponseError(c, appError.New(err, appError.BadRequestError, "Malformed request body"))
	}

	if errs := t.validator.Validate(req); len(errs) > 0 && errs[0].Error {
		err := appError.New(nil, appError.BadRequestError, errs[0].Message)
		err.AddDetail("failed_fields", errs)
		return http.ResponseError(c, err)
	}

	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return http.ResponseError(c, currErr)
	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.UpdateProduct(c.Context(), id, application.UpdateProductPayload{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return http.ResponseError(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}

// @Summary Delete one product.
// @Description deletes one product by id.
// @Tags Product
// @Accept */*
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} int64
// @Router /product/{id} [delete]
func (t *StoreController) deleteProduct(c *fiber.Ctx) error {
	if err := t.validator.ValidateField(c.Params("id"), "id", "required,uuid4"); err != nil {
		currErr := appError.New(nil, appError.BadRequestError, err.Message)
		currErr.AddDetail("failed_fields", err)
		return http.ResponseError(c, currErr)
	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.DeleteProduct(c.Context(), id)
	if err != nil {
		return http.ResponseError(c, err)
	}

	return c.Status(fiber.StatusOK).Send(http.Int64ToByte(*affected))
}
