package user

import (
	"ambil-api/domain"
	"ambil-api/utils"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input domain.AuthRequest) (string, error)
	Create(input domain.UserRequest) (domain.UserData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input domain.UserRequest) (domain.UserData, error) {

	newUser, _ := hashedUser(domain.UserRequest{
		Email:    input.Email,
		Password: input.Password,
	})

	user, err := s.repository.Create(newUser)

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
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return u, nil

}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
