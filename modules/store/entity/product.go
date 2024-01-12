package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/dhanielsales/golang-scaffold/internal/time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  uuid.UUID `json:"category_id"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   *int64    `json:"updated_at"`
}

func NewProduct(name, description string, price float64, CategoryID uuid.UUID) (*Product, error) {
	product := &Product{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug.Make(name),
		Price:       price,
		CategoryID:  CategoryID,
		Description: &description,
		CreatedAt:   time.Now(),
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (c *Product) Update(name, description string, price float64, CategoryID uuid.UUID) {
	if name != "" {
		c.Name = name
		c.Slug = slug.Make(name)
	}

	if description != "" {
		c.Description = &description
	}

	if price != 0 {
		c.Price = price
	}

	c.CategoryID = CategoryID

	updatedAt := time.Now()
	c.UpdatedAt = &updatedAt
}

func (c *Product) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Price == 0 || c.Price < 0 {
		return errors.New("price is required")
	}

	return nil
}
