package service

import (
	"alta/temanpetani/features/users"
)

type UserService struct {
	userData users.UserDataInterface
}

// RegisterUser implements users.UserServiceInterface
func (*UserService) RegisterUser(data users.CoreUserRequest) (userId uint, err error) {
	panic("unimplemented")
}

func New(userData users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: userData,
	}
}
