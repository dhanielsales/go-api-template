package controllers

import (
	"net/http"

	"github.com/dhanielsales/go-api-template/internal/modules/store/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// POST /api/v0/product
//
// @Summary Create one product.
// @Description creates one product.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		Product body createProductRequest true "Product to create"
// @Success		201 {object} int64
// @Header		201,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	error.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	error.AppError	"Internal Server Error."
// @Router /api/v0/product [post]
func (t *StoreController) createProduct(c *fiber.Ctx) error {
	var req createProductRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	affected, err := t.service.CreateProduct(c.Context(), service.CreateProductPayload{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategotyID:  uuid.MustParse(req.CategoryID),
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)

	}

	return c.Status(http.StatusOK).JSON(affected)
}

// GET /api/v0/product
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
// @Failure		500					{object}	error.AppError	"Internal Server Error."
// @Router /api/v0/product [get]
func (t *StoreController) getManyProduct(c *fiber.Ctx) error {
	categories, err := t.service.GetManyProduct(c.Context(), service.GetManyProductParams{
		Page:    c.Query("page"),
		PerPage: c.Query("perPage"),
		OrderBy: c.Query("orderBy"),
	})
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)

	}

	return c.JSON(categories)
}

// GET /api/v0/product/{id}
//
// @Summary Get one product.
// @Description fetch one product by id.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		id path string true "Product ID"
// @Success 	200 				{object} 	models.Product
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	error.AppError	"Internal Server Error."
// @Router /api/v0/product/{id} [get]
func (t *StoreController) getOneProduct(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	product, err := t.service.GetProductById(c.Context(), id)
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)

	}

	return c.JSON(product)
}

type updateProductRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description,omitempty" validate:"min=1,max=300"`
}

// PUT /api/v0/product
//
// @Summary Update one product.
// @Description updates one product by id.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param Product body updateProductRequest true "Product to update"
// @Success		200 {object} int64
// @Header		200,400,500	string		X-Conversational-ID	"Unique request ID."
// @Failure		400					{object}	error.AppError	"Bad Request. Invalid request body."
// @Failure		500					{object}	error.AppError	"Internal Server Error."
// @Success 200 {object} int64
// @Router /api/v0/product/{id} [put]
func (t *StoreController) updateProduct(c *fiber.Ctx) error {
	var req updateProductRequest

	if err := t.validator.DecodeAndValidate(c, req); err != nil {
		return t.http.ErrorHandler.Response(c, err)
	}

	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.UpdateProduct(c.Context(), id, service.UpdateProductPayload{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return t.http.ErrorHandler.Response(c, err)

	}

	return c.Status(http.StatusOK).JSON(affected)
}

// DELETE /api/v0/product/{id}
//
// @Summary Delete one product.
// @Description deletes one product by id.
// @Tags Product
// @Produce		json
// @Accept */*
// @Produce json
// @Param			X-Conversational-ID		header		string					false	"Unique request ID."
// @Param			Authorization		header		string					true	"Authorization JWT"
// @Param 		id path string true "Product ID"
// @Success 	200 {object} int64
// @Header		200,500			string		X-Conversational-ID	"Unique request ID."
// @Failure		500					{object}	error.AppError	"Internal Server Error."
// @Router /api/v0/product/{id} [delete]
func (t *StoreController) deleteProduct(c *fiber.Ctx) error {
	id := uuid.MustParse(c.Params("id"))
	affected, err := t.service.DeleteProduct(c.Context(), id)
	if err != nil {
		return t.http.ErrorHandler.Response(c, err)

	}

	return c.Status(http.StatusOK).JSON(affected)
}
