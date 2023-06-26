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

func (repo *templateQuery) InsertTask(input templates.TaskTemplateCore) error {
	templateInputGorm := NewTaskTemplateModel(input)

	tx := repo.db.Create(&templateInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *templateQuery) SelectAllSchedule() ([]templates.ScheduleTemplateCore, error) {
	var templatesData []ScheduleTemplate
	tx := repo.db.Find(&templatesData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error templates not found")
	}

	var templatesCoreAll []templates.ScheduleTemplateCore
	for _, value := range templatesData {
		templateCore := NewScheduleTemplateCore(value)
		templatesCoreAll = append(templatesCoreAll, templateCore)
	}
	return templatesCoreAll, nil
}

func (repo *templateQuery) SelectAllTasks(scheduleId uint64) ([]templates.TaskTemplateCore, error) {
	var templatesData []TaskTemplate
	tx := repo.db.Where("schedule_id = ?", scheduleId).Find(&templatesData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var templatesCoreAll []templates.TaskTemplateCore
	for _, value := range templatesData {
		templateCore := NewTaskTemplateCore(value)
		templatesCoreAll = append(templatesCoreAll, templateCore)
	}
	return templatesCoreAll, nil
}

func (repo *templateQuery) SelectScheduleById(id uint64) (templates.ScheduleTemplateCore, error) {
	var templateGorm ScheduleTemplate
	tx := repo.db.First(&templateGorm, id)
	if tx.Error != nil {
		return templates.ScheduleTemplateCore{}, errors.New("error template not found")
	}

	templateCore := NewScheduleTemplateCore(templateGorm)
	return templateCore, nil
}
