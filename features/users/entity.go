package users

type CoreUserRequest struct {
	FullName			string 	`validate:"required"`
	Email					string 	`validate:"required,email"`
	Password			string	`validate:"required,min=8"`
	Phone					string	`validate:"required"`
}

type UserDataInterface interface {
	Insert(data CoreUserRequest) (userId uint, err error)
}

type UserServiceInterface interface {
	RegisterUser(data CoreUserRequest) (userId uint, err error)
}

