package postgres_repository

import (
	"github.com/dhanielsales/go-api-template/internal/models"
	db "github.com/dhanielsales/go-api-template/internal/modules/store/repository/postgres/gen"
)

func ToProduct(product *db.Product) *models.Product {
	res := models.Product{
		ID:         product.ID,
		Name:       product.Name,
		Slug:       product.Slug,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}

	if product.UpdatedAt.Valid {
		res.UpdatedAt = &product.UpdatedAt.Int64
	}

	if product.Description.Valid {
		res.Description = &product.Description.String
	}

	return &res
}

func ToProducts(products []db.Product) []*models.Product {
	res := make([]*models.Product, len(products))

	for i, product := range products {
		res[i] = ToProduct(&product)
	}

	return res
}

func ToCategory(category *db.Category) *models.Category {
	res := models.Category{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}

	if category.UpdatedAt.Valid {
		res.UpdatedAt = &category.UpdatedAt.Int64
	}

	if category.Description.Valid {
		res.Description = &category.Description.String
	}

	return &res
}

func ToCategories(categories []db.Category) []*models.Category {
	res := make([]*models.Category, len(categories))

	for i, category := range categories {
		res[i] = ToCategory(&category)
	}

	return res
}
