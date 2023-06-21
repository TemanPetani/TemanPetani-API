package handler

import "alta/temanpetani/features/users"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	FullName      string `json:"full_name" form:"full_name"`
	Email         string `json:"email" form:"email"`
	Phone         string `json:"phone" form:"phone"`
	Password      string `json:"password" form:"password"`
	Role          string `json:"role" form:"role"`
	Address       string `json:"address" form:"address"`
	Avatar        string `json:"avatar" form:"avatar"`
	Bank          string `json:"bank" form:"bank"`
	AccountNumber string `json:"account_number" form:"account_number"`
}

func UserRequestToCore(userRequest UserRequest) users.UserCore {
	return users.UserCore{
		FullName:      userRequest.FullName,
		Email:         userRequest.Email,
		Password:      userRequest.Password,
		Phone:         userRequest.Phone,
		Role:          userRequest.Role,
		Address:       userRequest.Address,
		Avatar:        userRequest.Avatar,
		Bank:          userRequest.Bank,
		AccountNumber: userRequest.AccountNumber,
	}
}
