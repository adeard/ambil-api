package banner

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Store(input domain.BannerRequest) (domain.BannerData, error)
	StoreBannerType(input domain.BannerTypeRequest) (domain.BannerTypeData, error)
	Update(input domain.BannerData, updateData map[string]interface{}) error
	GetAll(input domain.BannerFilterRequest) ([]domain.BannerData, error)
	GetAllBannerType(input domain.BannerTypeFilterRequest) ([]domain.BannerTypeData, error)
	GetDetail(bannerId string) (domain.BannerData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(input domain.BannerFilterRequest) ([]domain.BannerData, error) {
	var banners []domain.BannerData
	q := r.db.Debug().Table("banners")

	if input.BannerTypeId != 0 {
		q = q.Where("banner_type_id = ?", input.BannerTypeId)
	}

	if input.IndexNo != 0 {
		q = q.Where("index_no = ?", input.IndexNo)
	}

	if input.IsActive >= 0 {
		q = q.Where("is_active = ?", input.IsActive)
	}

	if input.StartDate != "" {
		q = q.Where("start_date = ?", input.StartDate)
	}

	if input.EndDate != "" {
		q = q.Where("end_date = ?", input.EndDate)
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
		Find(&banners).
		Error

	return banners, err
}

func (r *repository) GetAllBannerType(input domain.BannerTypeFilterRequest) ([]domain.BannerTypeData, error) {
	var bannerTypes []domain.BannerTypeData
	q := r.db.Debug().Table("banner_types")

	if input.Name != "" {
		q = q.Where("name = ?", input.Name)
	}

	if input.IsActive >= 0 {
		q = q.Where("is_active = ?", input.IsActive)
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
		Find(&bannerTypes).
		Error

	return bannerTypes, err
}

func (r *repository) Store(input domain.BannerRequest) (domain.BannerData, error) {

	category := domain.BannerData{BannerRequest: input}

	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) StoreBannerType(input domain.BannerTypeRequest) (domain.BannerTypeData, error) {

	bannerType := domain.BannerTypeData{BannerTypeRequest: input}

	err := r.db.Create(&bannerType).Error

	return bannerType, err
}

func (r *repository) GetDetail(bannerId string) (domain.BannerData, error) {
	var banner domain.BannerData
	err := r.db.Where("id = ? ", bannerId).First(&banner).Error

	return banner, err
}

func (r *repository) Update(banner domain.BannerData, updateData map[string]interface{}) error {
	err := r.db.Debug().
		Model(&banner).
		Where("id = ?", banner.Id).
		Updates(updateData).Error

	return err
}
