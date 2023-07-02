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
			return users.UserCore{}, "", errors.New("Email dan Password Salah")
		} else {
			return users.UserCore{}, "", tx.Error
		}
	}

	checkPassword := helpers.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return users.UserCore{}, "", errors.New("Password Salah")
	}

	token, errToken := middlewares.CreateToken(userGorm.ID, userGorm.Role)
	if errToken != nil {
		return users.UserCore{}, "", errToken
	}

	dataCore := NewUserCore(userGorm)
	return dataCore, token, nil
}

func (repo *userQuery) Insert(input users.UserCore) error {
	userInputGorm := NewUserModel(input)
	tx := repo.db.Create(&userInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Menambahkan Data Pengguna")
	}

	return nil
}

func (repo *userQuery) SelectById(id uint64) (users.UserCore, error) {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		return users.UserCore{}, errors.New("Pengguna Tidak Ditemukan")
	}

	userCore := NewUserCore(userGorm)
	return userCore, nil
}

func (repo *userQuery) UpdateById(id uint64, input users.UserCore) error {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		return errors.New("Pengguna Tidak Ditemukan")
	}

	userInputGorm := NewUserModel(input)
	tx = repo.db.Model(&userGorm).Updates(userInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Memperbarui Data Pengguna")
	}

	return nil
}

func (repo *userQuery) UpdateImage(id uint64, imageUrl string) error {
	var userGorm User

	tx := repo.db.Model(&userGorm).Where("id = ?", id).Update("avatar", imageUrl)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *userQuery) DeleteById(id uint64) error {
	var userGorm User
	tx := repo.db.First(&userGorm, id)
	if tx.Error != nil {
		return errors.New("Pengguna Tidak Ditemukan")
	}

	tx = repo.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Pengguna Tidak Ditemukan")
	}

	return nil
}
