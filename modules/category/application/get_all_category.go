package category_application

import (
	"golang.org/x/net/context"

	category "github.com/dhanielsales/golang-scaffold/modules/category/entity"
)

func (s *CategoryService) GetAll(ctx context.Context) (*[]category.Category, error) {
	dbResult, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []category.Category

	for _, dbCategory := range *dbResult {
		result = append(result, category.Category{
			Id:          dbCategory.Id,
			Name:        dbCategory.Name,
			Description: dbCategory.Description.String,
			CreatedAt:   dbCategory.CreatedAt,
			UpdatedAt:   dbCategory.UpdatedAt,
		})
	}

	return &result, nil
}
