package database

import (
	_plantModel "alta/temanpetani/features/plants/data"
	_productModel "alta/temanpetani/features/products/data"
	_templateModel "alta/temanpetani/features/templates/data"
	_userModel "alta/temanpetani/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&_userModel.User{}, &_templateModel.ScheduleTemplate{}, &_templateModel.TaskTemplate{}, &_productModel.Products{}, &_plantModel.Schedule{}, &_plantModel.Task{})
	if err != nil {
		return err
	}
	return nil
}
