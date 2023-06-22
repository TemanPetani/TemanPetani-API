package data

import (
	"alta/temanpetani/features/users"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) Login(email string, password string) (users.UserCore, string, error) {
	var userGorm User
	tx := repo.db.Where("email = ?", email).First(&userGorm)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return users.UserCore{}, "", errors.New("login failed, wrong email and password")
		} else {
			return users.UserCore{}, "", tx.Error
		}
	}

	checkPassword := helpers.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return users.UserCore{}, "", errors.New("login failed, wrong password")
	}

	token, errToken := middlewares.CreateToken(userGorm.ID, userGorm.Role)
	if errToken != nil {
		return users.UserCore{}, "", errToken
	}

	dataCore := NewUserCore(userGorm)
	return dataCore, token, nil
}

func (repo *userQuery) Insert(input users.UserCore) error {
	hashedPassword, errHash := helpers.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}

	userInputGorm := NewUserModel(input)
	userInputGorm.Password = hashedPassword

	tx := repo.db.Create(&userInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *userQuery) SelectById(id uint64) (users.UserCore, error) {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		return users.UserCore{}, errors.New("error user not found")
	}

	userCore := NewUserCore(userGorm)
	return userCore, nil
}
