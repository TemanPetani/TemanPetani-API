package data

import (
	"alta/temanpetani/features/plants"
	"errors"

	"gorm.io/gorm"
)

type plantQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) plants.PlantDataInterface {
	return &plantQuery{
		db: db,
	}
}

func (repo *plantQuery) InsertSchedule(input plants.ScheduleCore) (plants.ScheduleCore, error) {
	plantInputGorm := NewScheduleModel(input)

	tx := repo.db.Create(&plantInputGorm)
	if tx.Error != nil {
		return plants.ScheduleCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return plants.ScheduleCore{}, errors.New("insert failed, row affected = 0")
	}

	scheduleData := NewScheduleCore(plantInputGorm)

	return scheduleData, nil
}

func (repo *plantQuery) InsertTask(input []plants.TaskCore) error {
	var tasksInputGorm []Task
	for _, value := range input {
		taskGorm := NewTaskModel(value)
		tasksInputGorm = append(tasksInputGorm, taskGorm)
	}

	tx := repo.db.Create(&tasksInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
