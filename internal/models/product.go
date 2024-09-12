package models

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/dhanielsales/go-api-template/pkg/utils"
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

func (c *Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func NewProduct(name, description string, price float64, categoryID uuid.UUID) (*Product, error) {
	product := &Product{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug.Make(name),
		Price:       price,
		CategoryID:  categoryID,
		Description: &description,
		CreatedAt:   utils.TimeNow(),
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (c *Product) Update(name, description string, price float64, categoryID uuid.UUID) {
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

	c.CategoryID = categoryID

	updatedAt := utils.TimeNow()
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
