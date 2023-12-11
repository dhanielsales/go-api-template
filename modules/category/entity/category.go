package category_entity

import (
	"github.com/dhanielsales/golang-scaffold/internal/time"
	"github.com/gofrs/uuid"
)

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func NewCategory(name, description string) (*Category, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Category{
		Id:          uuid.String(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
