package user

import (
	"ambil-api/domain"
	"ambil-api/utils"
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input domain.AuthRequest) (string, error)
	Create(input domain.RegisterRequest) (domain.UserData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input domain.RegisterRequest) (domain.UserData, error) {

	if input.Email != input.ConfirmPassword {
		return domain.UserData{}, errors.New("password not match")
	}

	newUser, _ := hashedUser(domain.UserRequest{
		Email:    input.Email,
		Password: input.Password,
	})

	user, err := s.repository.Create(newUser)
	if err != nil {
		return domain.UserData{}, err
	}

	_, err = s.repository.CreateDescription(domain.UserDescriptionRequest{
		UserId:      user.Id,
		Fullname:    input.Fullname,
		PhoneNumber: input.PhoneNumber,
	})

	return user, err
}

func (s *service) Login(input domain.AuthRequest) (string, error) {
	userCheck, err := s.repository.GetDetail(
		domain.UserData{
			UserRequest: domain.UserRequest{
				Email: input.Email,
			},
		},
	)

	if err != nil {
		return "", err
	}

	err = verifyPassword(input.Password, userCheck.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(userCheck.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func hashedUser(u domain.UserRequest) (domain.UserRequest, error) {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = string(hashedPassword)

	return u, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
