package data

import (
	"alta/temanpetani/features/templates"
	"errors"

	"gorm.io/gorm"
)

type templateQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) templates.TemplateDataInterface {
	return &templateQuery{
		db: db,
	}
}

func (repo *templateQuery) InsertSchedule(input templates.ScheduleTemplateCore) error {
	templateInputGorm := NewScheduleTemplateModel(input)

	tx := repo.db.Create(&templateInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
