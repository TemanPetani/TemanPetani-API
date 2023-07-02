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
		return errors.New("Gagal Menambahkan Data Template")
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
		return errors.New("Gagal Menambahkan Data Template")
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
		return nil, errors.New("Data Template Tidak Ditemukan")
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
		return templates.ScheduleTemplateCore{}, errors.New("Data Template Tidak Ditemukan")
	}

	templateCore := NewScheduleTemplateCore(templateGorm)
	return templateCore, nil
}

func (repo *templateQuery) UpdateScheduleById(id uint64, input templates.ScheduleTemplateCore) error {
	var templateGorm ScheduleTemplate
	tx := repo.db.First(&templateGorm, id)
	if tx.Error != nil {
		return errors.New("Data Template Tidak Ditemukan")
	}

	templateInputGorm := NewScheduleTemplateModel(input)
	tx = repo.db.Model(&templateGorm).Updates(templateInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Memperbarui Data Template")
	}

	return nil
}

func (repo *templateQuery) DeleteScheduleById(id uint64) error {
	var templateGorm ScheduleTemplate
	tx := repo.db.First(&templateGorm, id)
	if tx.Error != nil {
		return errors.New("Data Template Tidak Ditemukan")
	}

	tx = repo.db.Delete(&templateGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Menghapus Data Template")
	}

	return nil
}

func (repo *templateQuery) UpdateTaskById(id uint64, input templates.TaskTemplateCore) error {
	var templateGorm TaskTemplate
	tx := repo.db.First(&templateGorm, id)
	if tx.Error != nil {
		return errors.New("Data Template Tidak Ditemukan")
	}

	templateInputGorm := NewTaskTemplateModel(input)
	tx = repo.db.Model(&templateGorm).Updates(templateInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Memperbarui Data Template")
	}

	return nil
}

func (repo *templateQuery) DeleteTaskById(id uint64) error {
	var templateGorm TaskTemplate
	tx := repo.db.First(&templateGorm, id)
	if tx.Error != nil {
		return errors.New("Data Template Tidak Ditemukan")
	}

	tx = repo.db.Delete(&templateGorm, id)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Menghapus Data Template")
	}

	return nil
}
