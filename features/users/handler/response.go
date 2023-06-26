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
	FullName      string `json:"fullname,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Password      string `json:"password,omitempty"`
	Role          string `json:"role,omitempty"`
	Address       string `json:"address,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Bank          string `json:"bank,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
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
