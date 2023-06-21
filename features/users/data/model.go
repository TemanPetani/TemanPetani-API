package data

import (
	"alta/temanpetani/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint64 `gorm:"primarykey"`
	FullName      string `gorm:"notNull"`
	Email         string `gorm:"unique;notNull"`
	Phone         string `gorm:"unique:notNull"`
	Password      string `gorm:"notNull"`
	Role          string `gorm:"type:enum('admin','user');default:'user'"`
	Address       string `gorm:"type:text"`
	Avatar        string
	Bank          string
	AccountNumber string `gorm:"unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func NewUserCore(userData User) users.UserCore {
	return users.UserCore{
		ID:            userData.ID,
		FullName:      userData.FullName,
		Email:         userData.Email,
		Phone:         userData.Phone,
		Password:      userData.Password,
		Role:          userData.Role,
		Address:       userData.Address,
		Avatar:        userData.Avatar,
		Bank:          userData.Bank,
		AccountNumber: userData.AccountNumber,
	}
}

func NewUserModel(dataCore users.UserCore) User {
	return User{
		ID:            dataCore.ID,
		FullName:      dataCore.FullName,
		Email:         dataCore.Email,
		Phone:         dataCore.Phone,
		Password:      dataCore.Password,
		Role:          dataCore.Role,
		Address:       dataCore.Address,
		Avatar:        dataCore.Avatar,
		Bank:          dataCore.Bank,
		AccountNumber: dataCore.AccountNumber,
	}
}
