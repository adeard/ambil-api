package merchant

import (
	"ambil-api/domain"
	"ambil-api/utils"
)

type Service interface {
	GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error)
	GetAllCategory(input domain.MerchantCategoryFilterRequest) ([]domain.MerchantCategoryData, error)
	GetDetail(merchantId string) (domain.MerchantData, error)
	Store(input domain.MerchantRequest) (domain.MerchantData, error)
	StoreCategory(input domain.MerchantCategoryRequest) (domain.MerchantCategoryData, error)
	Update(driverId string, updatedData domain.MerchantRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	merchants, err := s.repository.GetAll(input)
	if err != nil {
		return []domain.MerchantData{}, err
	}

	return merchants, err
}

func (s *service) GetAllCategory(input domain.MerchantCategoryFilterRequest) ([]domain.MerchantCategoryData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	categories, err := s.repository.GetAllCategory(input)
	if err != nil {
		return []domain.MerchantCategoryData{}, err
	}

	return categories, err
}

func (s *service) GetDetail(merchantId string) (domain.MerchantData, error) {
	merchant, err := s.repository.GetDetail(merchantId)

	return merchant, err
}

func (s *service) Store(input domain.MerchantRequest) (domain.MerchantData, error) {

	input.IsActive = 1
	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchant, err := s.repository.Store(input)
	if err != nil {
		return domain.MerchantData{}, err
	}

	return merchant, err
}

func (s *service) StoreCategory(input domain.MerchantCategoryRequest) (domain.MerchantCategoryData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchantCategory, err := s.repository.StoreCategory(input)
	if err != nil {
		return domain.MerchantCategoryData{}, err
	}

	return merchantCategory, err
}

func (s *service) Update(merchantId string, input domain.MerchantRequest) error {

	merchant, err := s.repository.GetDetail(merchantId)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"updated_at": utils.GetCurrentDateTime(),
	}

	if input.UserId != 0 {
		updateData["user_id"] = input.UserId
	}

	if input.Address != "" {
		updateData["address"] = input.Address
	}

	if input.Coordinate != "" {
		updateData["coordinate"] = input.Coordinate
	}

	if input.Name != "" {
		updateData["name"] = input.Name
	}

	if input.PhoneNumber != "" {
		updateData["phone_number"] = input.PhoneNumber
	}

	if input.Picture != "" {
		updateData["picture"] = input.Picture
	}

	if input.IsActive >= 0 {
		updateData["is_active"] = input.IsActive
	}

	if input.Score != 0 {
		updateData["score"] = input.Score
	}

	err = s.repository.Update(merchant, updateData)

	return err
}
