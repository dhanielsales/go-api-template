package category_application

import category_storage "github.com/dhanielsales/golang-scaffold/modules/category/storage"

type CategoryService struct {
	storage *category_storage.CategoryStorage
}

func NewCategoryService(storage *category_storage.CategoryStorage) *CategoryService {
	return &CategoryService{
		storage: storage,
	}
}
