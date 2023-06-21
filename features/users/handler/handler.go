package handler

import "alta/temanpetani/features/users"

type UserHandler struct {
	service users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}