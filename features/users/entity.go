package users

import "mime/multipart"

type UserCore struct {
	ID            uint64
	FullName      string `validate:"required"`
	Email         string `validate:"required,email"`
	Phone         string `validate:"required"`
	Password      string `validate:"required"`
	Role          string
	Address       string
	Avatar        string
	Bank          string
	AccountNumber string
}

type UserDataInterface interface {
	Login(email string, password string) (UserCore, string, error)
	Insert(input UserCore) error
	SelectById(id uint64) (UserCore, error)
	UpdateById(id uint64, input UserCore) error
	UpdateImage(id uint64, imageUrl string) error
	DeleteById(id uint64) error
}

type UserServiceInterface interface {
	Login(email string, password string) (UserCore, string, error)
	Create(input UserCore) error
	GetById(id uint64) (UserCore, error)
	UpdateById(id uint64, input UserCore) error
	UpdateImage(id uint64, image *multipart.FileHeader) error
	DeleteById(id uint64) error
}
