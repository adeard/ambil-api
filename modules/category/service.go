package category

import (
	"ambil-api/domain"
	"ambil-api/utils"
)

type Service interface {
	GetAll(input domain.CategoryFilterRequest) ([]domain.CategoryData, error)
	GetDetail(categoryId string) (domain.CategoryData, error)
	Store(input domain.CategoryRequest) (domain.CategoryData, error)
	Update(categoryId string, updatedData domain.CategoryRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.CategoryFilterRequest) ([]domain.CategoryData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	categories, err := s.repository.GetAll(input)
	if err != nil {
		return []domain.CategoryData{}, err
	}

	return categories, err
}

func (s *service) GetDetail(categoryId string) (domain.CategoryData, error) {
	category, err := s.repository.GetDetail(categoryId)

	return category, err
}

func (s *service) Store(input domain.CategoryRequest) (domain.CategoryData, error) {

	input.IsActive = 1
	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	category, err := s.repository.Store(input)
	if err != nil {
		return domain.CategoryData{}, err
	}

	return category, err
}

func (s *service) Update(categoryId string, input domain.CategoryRequest) error {

	category, err := s.repository.GetDetail(categoryId)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"updated_at": utils.GetCurrentDateTime(),
	}

	if input.Name != "" {
		updateData["name"] = input.Name
	}

	err = s.repository.Update(category, updateData)

	return err
}
