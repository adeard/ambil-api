package user

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Create(input domain.UserRequest) (domain.UserData, error)
	CreateUserLevel(input domain.UserLevelRequest) (domain.UserLevelData, error)
	CreateDescription(input domain.UserDescriptionRequest) (domain.UserDescriptionData, error)
	CreateMerchantFavourite(input domain.UserMerchantFavouriteRequest) (domain.UserMerchantFavouriteData, error)
	GetDetail(input domain.UserData) (domain.UserData, error)
	GetDetailDescription(input domain.UserDescriptionData) (domain.UserDescriptionData, error)
	GetUserLevel(input domain.UserLevelRequest) ([]domain.UserLevelData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(input domain.UserRequest) (domain.UserData, error) {

	user := domain.UserData{UserRequest: input}

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CreateUserLevel(input domain.UserLevelRequest) (domain.UserLevelData, error) {

	userLevel := domain.UserLevelData{UserLevelRequest: input}

	err := r.db.Create(&userLevel).Error

	return userLevel, err
}

func (r *repository) CreateDescription(input domain.UserDescriptionRequest) (domain.UserDescriptionData, error) {

	user := domain.UserDescriptionData{UserDescriptionRequest: input}

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CreateMerchantFavourite(input domain.UserMerchantFavouriteRequest) (domain.UserMerchantFavouriteData, error) {

	user := domain.UserMerchantFavouriteData{UserMerchantFavouriteRequest: input}

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) GetDetail(input domain.UserData) (domain.UserData, error) {
	var user domain.UserData

	q := r.db.Debug()

	if input.Email != "" {
		q = q.Where("email = ?", input.Email)
	}

	if input.Id != 0 {
		q = q.Where("id = ?", input.Id)
	}

	err := q.First(&user).Error

	return user, err
}

func (r *repository) GetDetailDescription(input domain.UserDescriptionData) (domain.UserDescriptionData, error) {
	var user domain.UserDescriptionData

	q := r.db.Debug()

	if input.Fullname != "" {
		q = q.Where("fullname = ?", input.Fullname)
	}

	if input.UserId != 0 {
		q = q.Where("user_id = ?", input.UserId)
	}

	if input.Id != 0 {
		q = q.Where("id = ?", input.Id)
	}

	if input.PhoneNumber != "" {
		q = q.Where("phone_number = ?", input.PhoneNumber)
	}

	err := q.First(&user).Error

	return user, err
}

func (r *repository) GetUserLevel(input domain.UserLevelRequest) ([]domain.UserLevelData, error) {
	var userLevels []domain.UserLevelData

	q := r.db.Debug()

	if input.LevelName != "" {
		q = q.Where("level_name = ?", input.LevelName)
	}

	err := q.Find(&userLevels).Error

	return userLevels, err
}
