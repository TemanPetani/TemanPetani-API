package handler

import (
	"alta/temanpetani/features/users"
)

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	FullName      string `json:"fullname" form:"fullname"`
	Email         string `json:"email" form:"email"`
	Phone         string `json:"phone" form:"phone"`
	Password      string `json:"password" form:"password"`
	Role          string `json:"role" form:"role"`
	Address       string `json:"address" form:"address"`
	Bank          string `json:"bank" form:"bank"`
	AccountNumber string `json:"accountNumber" form:"accountNumber"`
}

func UserRequestToCore(userRequest UserRequest) users.UserCore {
	return users.UserCore{
		FullName:      userRequest.FullName,
		Email:         userRequest.Email,
		Password:      userRequest.Password,
		Phone:         userRequest.Phone,
		Role:          userRequest.Role,
		Address:       userRequest.Address,
		Bank:          userRequest.Bank,
		AccountNumber: userRequest.AccountNumber,
	}
}
