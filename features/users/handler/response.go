package handler

import "alta/temanpetani/features/users"

type AuthResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

type UserResponse struct {
	ID            uint64 `json:"id,omitempty"`
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

func NewAuthResponse(user users.UserCore, jwtToken string) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: jwtToken,
		Role:  user.Role,
	}
}

func NewUserResponse(user users.UserCore) UserResponse {
	return UserResponse{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Role:          user.Role,
		Address:       user.Address,
		Avatar:        user.Avatar,
		Bank:          user.Bank,
		AccountNumber: user.AccountNumber,
	}
}
