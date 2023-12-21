package user

import (
	"ambil-api/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Create(input domain.UserRequest) (domain.UserData, error)
	GetDetail(input domain.UserData) (domain.UserData, error)
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
