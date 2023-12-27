package category

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.CategoryRequest) (domain.CategoryData, error)
	Update(input domain.CategoryData, updateData map[string]interface{}) error
	GetAll(input domain.CategoryFilterRequest) ([]domain.CategoryData, error)
	GetDetail(categoryId string) (domain.CategoryData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(input domain.CategoryFilterRequest) ([]domain.CategoryData, error) {
	var drivers []domain.CategoryData
	q := r.db.Debug().Table("categories")

	if input.Name != "" {
		q = q.Where("name = ?", input.Name)
	}

	if input.OrderBy != "" {
		sort := "asc"
		order := input.OrderBy

		if input.SortBy != "" {
			sort = input.SortBy
		}

		q = q.Order(order + " " + sort)
	}

	offset := (input.Limit * (input.Page - 1))

	err := q.
		Limit(input.Limit).
		Offset(offset).
		Find(&drivers).
		Error

	return drivers, err
}

func (r *repository) Store(input domain.CategoryRequest) (domain.CategoryData, error) {

	category := domain.CategoryData{CategoryRequest: input}

	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) GetDetail(categoryId string) (domain.CategoryData, error) {
	var category domain.CategoryData
	err := r.db.Where("id = ? ", categoryId).First(&category).Error

	return category, err
}

func (r *repository) Update(category domain.CategoryData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&category).
		Where("id = ?", category.Id).
		Updates(updateData).Error

	return err
}
