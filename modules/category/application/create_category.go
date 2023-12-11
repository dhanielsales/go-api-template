package category_application

import (
	"golang.org/x/net/context"

	category "github.com/dhanielsales/golang-scaffold/modules/category/entity"
)

type CreateCategoryPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *CategoryService) Create(ctx context.Context, data CreateCategoryPayload) error {
	category, err := category.NewCategory(data.Name, data.Description)
	if err != nil {
		return err
	}

	if err := s.storage.Create(ctx, category); err != nil {
		return err
	}

	return nil
}
