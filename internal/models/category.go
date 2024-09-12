package models

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/dhanielsales/go-api-template/pkg/utils"
)

type Category struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description *string    `json:"description"`
	Products    []*Product `json:"products,omitempty"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   *int64     `json:"updated_at"`
}

func (c *Category) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Category) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func NewCategory(name, description string) (*Category, error) {
	category := &Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug.Make(name),
		Description: &description,
		CreatedAt:   utils.TimeNow(),
	}

	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) Update(name, description string) {
	if name != "" {
		c.Name = name
		c.Slug = slug.Make(name)
	}

	if description != "" {
		c.Description = &description
	}

	updatedAt := utils.TimeNow()
	c.UpdatedAt = &updatedAt
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
