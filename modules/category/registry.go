package category

import "gorm.io/gorm"

func CategoryRegistry(db *gorm.DB) Service {
	categoryRepository := NewRepository(db)
	categoryService := NewService(categoryRepository)

	return categoryService
}
