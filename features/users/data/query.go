package data

import (
	"alta/temanpetani/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Insert implements users.UserDataInterface
func (*UserData) Insert(data users.CoreUserRequest) (userId uint, err error) {
	panic("unimplemented")
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
