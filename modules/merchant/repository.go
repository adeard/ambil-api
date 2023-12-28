package merchant

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error)
	GetAllItem(input domain.MerchantItemFilterRequest) ([]domain.MerchantItemData, error)
	GetAllCategory(input domain.MerchantCategoryFilterRequest) ([]domain.MerchantCategoryData, error)
	GetDetail(merchantId string) (domain.MerchantData, error)
	GetDetailItem(merchantItemId string) (domain.MerchantItemData, error)
	Store(input domain.MerchantRequest) (domain.MerchantData, error)
	StoreItem(input domain.MerchantItemRequest) (domain.MerchantItemData, error)
	StoreCategory(input domain.MerchantCategoryRequest) (domain.MerchantCategoryData, error)
	Update(input domain.MerchantData, updateData map[string]interface{}) error
	UpdateItem(input domain.MerchantItemData, updateData map[string]interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error) {
	var drivers []domain.MerchantData
	q := r.db.Debug().Table("merchants")

	if input.Address != "" {
		q = q.Where("address = ?", input.Address)
	}

	if input.Coordinate != "" {
		q = q.Where("coordinate = ?", input.Coordinate)
	}

	if input.Name != "" {
		q = q.Where("name = ?", input.Name)
	}

	if input.PhoneNumber != "" {
		q = q.Where("phone_number = ?", input.PhoneNumber)
	}

	if input.Picture != "" {
		q = q.Where("picture = ?", input.Picture)
	}

	if input.UserId != 0 {
		q = q.Where("user_id = ?", input.UserId)
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

func (r *repository) GetAllCategory(input domain.MerchantCategoryFilterRequest) ([]domain.MerchantCategoryData, error) {
	var drivers []domain.MerchantCategoryData
	q := r.db.Debug().Table("merchant_categories")

	if input.MerchantId != 0 {
		q = q.Where("merchant_id = ?", input.MerchantId)
	}

	if input.CategoryId != 0 {
		q = q.Where("category_id = ?", input.CategoryId)
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

func (r *repository) GetAllItem(input domain.MerchantItemFilterRequest) ([]domain.MerchantItemData, error) {
	var drivers []domain.MerchantItemData
	q := r.db.Debug().Table("merchant_items")

	if input.IsActive >= 0 {
		q = q.Where("is_active = ?", input.IsActive)
	}

	if input.Price != 0 {
		q = q.Where("price = ?", input.Price)
	}

	if input.TotalSold != 0 {
		q = q.Where("total_sold = ?", input.TotalSold)
	}

	if input.MerchantId != 0 {
		q = q.Where("merchant_id = ?", input.MerchantId)
	}

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

func (r *repository) Store(input domain.MerchantRequest) (domain.MerchantData, error) {

	merchant := domain.MerchantData{MerchantRequest: input}

	err := r.db.Create(&merchant).Error

	return merchant, err
}

func (r *repository) StoreItem(input domain.MerchantItemRequest) (domain.MerchantItemData, error) {

	merchantItem := domain.MerchantItemData{MerchantItemRequest: input}

	err := r.db.Create(&merchantItem).Error

	return merchantItem, err
}

func (r *repository) StoreCategory(input domain.MerchantCategoryRequest) (domain.MerchantCategoryData, error) {

	merchantCategory := domain.MerchantCategoryData{MerchantCategoryRequest: input}

	err := r.db.Create(&merchantCategory).Error

	return merchantCategory, err
}

func (r *repository) GetDetail(merchantId string) (domain.MerchantData, error) {
	var merchant domain.MerchantData
	err := r.db.Where("id = ? ", merchantId).First(&merchant).Error

	return merchant, err
}

func (r *repository) GetDetailItem(merchantItemId string) (domain.MerchantItemData, error) {
	var merchantItem domain.MerchantItemData
	err := r.db.Where("id = ? ", merchantItemId).First(&merchantItem).Error

	return merchantItem, err
}

func (r *repository) Update(merchant domain.MerchantData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&merchant).
		Where("id = ?", merchant.Id).
		Updates(updateData).Error

	return err
}

func (r *repository) UpdateItem(merchantItem domain.MerchantItemData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&merchantItem).
		Where("id = ?", merchantItem.Id).
		Updates(updateData).Error

	return err
}
