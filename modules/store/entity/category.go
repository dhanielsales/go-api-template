package entity

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/dhanielsales/golang-scaffold/internal/time"
)

type Category struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description *string    `json:"description"`
	Products    *[]Product `json:"products,omitempty"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   *int64     `json:"updated_at"`
}

func NewCategory(name, description string) *Category {
	return &Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug.Make(name),
		Description: &description,
		CreatedAt:   time.Now(),
	}
}

func (c *Category) Update(name, description string) {
	if name != "" {
		c.Name = name
		c.Slug = slug.Make(name)
	}

	if description != "" {
		c.Description = &description
	}

	updatedAt := time.Now()
	c.UpdatedAt = &updatedAt
}
