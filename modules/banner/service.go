package banner

import (
	"ambil-api/domain"
	"ambil-api/utils"
)

type Service interface {
	GetAll(input domain.BannerFilterRequest) ([]domain.BannerData, error)
	GetAllBannerType(input domain.BannerTypeFilterRequest) ([]domain.BannerTypeData, error)
	GetDetail(bannerId string) (domain.BannerData, error)
	Store(input domain.BannerRequest) (domain.BannerData, error)
	StoreBannerType(input domain.BannerTypeRequest) (domain.BannerTypeData, error)
	Update(bannerId string, updatedData domain.BannerRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(input domain.BannerFilterRequest) ([]domain.BannerData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	categories, err := s.repository.GetAll(input)
	if err != nil {
		return []domain.BannerData{}, err
	}

	return categories, err
}

func (s *service) GetAllBannerType(input domain.BannerTypeFilterRequest) ([]domain.BannerTypeData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	bannerType, err := s.repository.GetAllBannerType(input)
	if err != nil {
		return []domain.BannerTypeData{}, err
	}

	return bannerType, err
}

func (s *service) GetDetail(bannerId string) (domain.BannerData, error) {
	category, err := s.repository.GetDetail(bannerId)

	return category, err
}

func (s *service) Store(input domain.BannerRequest) (domain.BannerData, error) {

	input.IsActive = 1
	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	category, err := s.repository.Store(input)
	if err != nil {
		return domain.BannerData{}, err
	}

	return category, err
}

func (s *service) StoreBannerType(input domain.BannerTypeRequest) (domain.BannerTypeData, error) {

	input.IsActive = 1
	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	bannerType, err := s.repository.StoreBannerType(input)
	if err != nil {
		return domain.BannerTypeData{}, err
	}

	return bannerType, err
}

func (s *service) Update(bannerId string, input domain.BannerRequest) error {

	category, err := s.repository.GetDetail(bannerId)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"updated_at": utils.GetCurrentDateTime(),
	}

	if input.BannerTypeId != 0 {
		updateData["banner_type_id"] = input.BannerTypeId
	}

	if input.IndexNo != 0 {
		updateData["index_no"] = input.IndexNo
	}

	if input.IsActive >= 0 {
		updateData["is_active"] = input.IsActive
	}

	if input.StartDate != "" {
		updateData["start_date"] = input.StartDate
	}

	if input.EndDate != "" {
		updateData["end_date"] = input.EndDate
	}

	err = s.repository.Update(category, updateData)

	return err
}
