package service

import (
	"alta/temanpetani/features/users"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

func (service *userService) Login(email string, password string) (users.UserCore, string, error) {
	if email == "" || password == "" {
		return users.UserCore{}, "", errors.New("error validation: email or password cannot be empty")
	}
	dataLogin, token, err := service.userData.Login(email, password)
	return dataLogin, token, err
}
