package users

type UserCore struct {
	ID            uint64
	FullName      string `validate:"required"`
	Email         string `validate:"required,email"`
	Phone         string `validate:"required"`
	Password      string `validate:"required,min=8"`
	Role          string
	Address       string
	Avatar        string
	Bank          string
	AccountNumber string
}

type UserDataInterface interface {
	Login(email string, password string) (UserCore, string, error)
	Insert(input UserCore) error
}

type UserServiceInterface interface {
	Login(email string, password string) (UserCore, string, error)
	Create(input UserCore) error
}
