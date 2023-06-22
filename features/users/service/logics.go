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

func (service *userService) Create(input users.UserCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
