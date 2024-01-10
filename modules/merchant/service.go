package merchant

import (
	"ambil-api/domain"
	"ambil-api/utils"
)

type Service interface {
	GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error)
	GetAllItem(input domain.MerchantItemFilterRequest) ([]domain.MerchantItemData, error)
	GetAllGallery(input domain.MerchantGalleryFilterRequest) ([]domain.MerchantGalleryData, error)
	GetAllCategory(input domain.MerchantCategoryFilterRequest) ([]domain.MerchantCategoryData, error)
	GetDetail(merchantId string) (domain.MerchantData, error)
	Store(input domain.MerchantRequest) (domain.MerchantData, error)
	StoreItem(input domain.MerchantItemRequest) (domain.MerchantItemData, error)
	StoreRating(input domain.MerchantRatingRequest) (domain.MerchantRatingData, error)
	StoreRatingImage(input domain.MerchantRatingImageRequest) (domain.MerchantRatingImageData, error)
	StoreGallery(input domain.MerchantGalleryRequest) (domain.MerchantGalleryData, error)
	StoreCategory(input domain.MerchantCategoryRequest) (domain.MerchantCategoryData, error)
	Update(merchantId string, updatedData domain.MerchantRequest) error
	UpdateItem(merchantItemId string, updatedData domain.MerchantItemRequest) error
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

func (s *service) GetAllGallery(input domain.MerchantGalleryFilterRequest) ([]domain.MerchantGalleryData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	galleries, err := s.repository.GetAllGallery(input)
	if err != nil {
		return []domain.MerchantGalleryData{}, err
	}

	return galleries, err
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

func (s *service) GetAllItem(input domain.MerchantItemFilterRequest) ([]domain.MerchantItemData, error) {

	if input.Limit == 0 {
		input.Limit = 20
	}

	if input.Page == 0 {
		input.Page = 1
	}

	items, err := s.repository.GetAllItem(input)
	if err != nil {
		return []domain.MerchantItemData{}, err
	}

	return items, err
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

func (s *service) StoreGallery(input domain.MerchantGalleryRequest) (domain.MerchantGalleryData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchantGallery, err := s.repository.StoreGallery(input)
	if err != nil {
		return domain.MerchantGalleryData{}, err
	}

	return merchantGallery, err
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

func (s *service) StoreItem(input domain.MerchantItemRequest) (domain.MerchantItemData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchantItem, err := s.repository.StoreItem(input)
	if err != nil {
		return domain.MerchantItemData{}, err
	}

	return merchantItem, err
}

func (s *service) StoreRating(input domain.MerchantRatingRequest) (domain.MerchantRatingData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchantRating, err := s.repository.StoreRating(input)
	if err != nil {
		return domain.MerchantRatingData{}, err
	}

	return merchantRating, err
}

func (s *service) StoreRatingImage(input domain.MerchantRatingImageRequest) (domain.MerchantRatingImageData, error) {

	input.CreatedAt = utils.GetCurrentDateTime()
	input.UpdatedAt = utils.GetCurrentDateTime()

	merchantRatingImage, err := s.repository.StoreRatingImage(input)
	if err != nil {
		return domain.MerchantRatingImageData{}, err
	}

	return merchantRatingImage, err
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

func (s *service) UpdateItem(merchantItemId string, input domain.MerchantItemRequest) error {

	merchantItem, err := s.repository.GetDetailItem(merchantItemId)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"updated_at": utils.GetCurrentDateTime(),
	}

	if input.Price != 0 {
		updateData["price"] = input.Price
	}

	if input.Picture != "" {
		updateData["picture"] = input.Picture
	}

	if input.Name != "" {
		updateData["name"] = input.Name
	}

	if input.IsActive >= 0 {
		updateData["is_active"] = input.IsActive
	}

	err = s.repository.UpdateItem(merchantItem, updateData)

	return err
}
