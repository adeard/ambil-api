package merchant

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.MerchantRequest) (domain.MerchantData, error)
	Update(input domain.MerchantData, updateData map[string]interface{}) error
	GetAll(input domain.MerchantFilterRequest) ([]domain.MerchantData, error)
	GetDetail(merchantId string) (domain.MerchantData, error)
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

func (r *repository) Store(input domain.MerchantRequest) (domain.MerchantData, error) {

	merchant := domain.MerchantData{MerchantRequest: input}

	err := r.db.Create(&merchant).Error

	return merchant, err
}

func (r *repository) GetDetail(merchantId string) (domain.MerchantData, error) {
	var driver domain.MerchantData
	err := r.db.Where("id = ? ", merchantId).First(&driver).Error

	return driver, err
}

func (r *repository) Update(merchant domain.MerchantData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&merchant).
		Where("id = ?", merchant.Id).
		Updates(updateData).Error

	return err
}
