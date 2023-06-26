package database

import (
	_userModel "alta/temanpetani/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	db.Exec("DROP TABLE users")
	err := db.AutoMigrate(&_userModel.User{})
	if err != nil {
		return err
	}
	return nil
}
