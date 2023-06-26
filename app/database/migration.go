package database

import (
	_templateModel "alta/temanpetani/features/templates/data"
	_userModel "alta/temanpetani/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&_userModel.User{}, &_templateModel.ScheduleTemplate{})
	if err != nil {
		return err
	}
	return nil
}
