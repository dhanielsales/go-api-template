package entity

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/dhanielsales/golang-scaffold/internal/utils"
)

type Category struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description *string    `json:"description"`
	Products    *[]Product `json:"products,omitempty"`
	ImageUrl    *string    `json:"image_url"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   *int64     `json:"updated_at"`
}

func (c Category) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Category) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func NewCategory(name, description string) *Category {
	return &Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug.Make(name),
		Description: &description,
		CreatedAt:   utils.TimeNow(),
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

	updatedAt := utils.TimeNow()
	c.UpdatedAt = &updatedAt
}
