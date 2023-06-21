package migration

import (
	_userModel "alta/temanpetani/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&_userModel.Users{})
	if err != nil {
		return err
	}
	return nil
}