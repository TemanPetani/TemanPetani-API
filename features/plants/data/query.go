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

func (repo *plantQuery) SelectAllSchedule() ([]plants.ScheduleCore, error) {
	query := ("select s.id as schedule_id, u.full_name as farmer_name, " +
		"s.name as schedule_name from schedules as s inner join users as u " +
		"on s.user_id = u.id where s.deleted_at is null order by s.updated_at desc")

	var plantsData []FarmerSchedule
	tx := repo.db.Raw(query).Scan(&plantsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var plantsCoreAll []plants.ScheduleCore
	for _, value := range plantsData {
		plantCore := NewFarmerSchedule(value)
		plantsCoreAll = append(plantsCoreAll, plantCore)
	}

	return plantsCoreAll, nil
}

func (repo *plantQuery) SelectAllFarmerSchedule(farmerId uint64) ([]plants.ScheduleCore, error) {
	var plantsData []Schedule
	tx := repo.db.Where("user_id = ?", farmerId).Find(&plantsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var plantsCoreAll []plants.ScheduleCore
	for _, value := range plantsData {
		plantCore := NewScheduleCore(value)
		plantsCoreAll = append(plantsCoreAll, plantCore)
	}
	return plantsCoreAll, nil
}

func (repo *plantQuery) SelectAllTasks(scheduleId uint64) ([]plants.TaskCore, error) {
	var plantsData []Task
	tx := repo.db.Where("schedule_id = ?", scheduleId).Find(&plantsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var plantsCoreAll []plants.TaskCore
	for _, value := range plantsData {
		plantCore := NewTaskCore(value)
		plantsCoreAll = append(plantsCoreAll, plantCore)
	}
	return plantsCoreAll, nil
}

func (repo *plantQuery) SelectScheduleById(id uint64) (plants.ScheduleCore, error) {
	var plantGorm Schedule
	tx := repo.db.First(&plantGorm, id)
	if tx.Error != nil {
		return plants.ScheduleCore{}, errors.New("error template not found")
	}

	plantCore := NewScheduleCore(plantGorm)
	return plantCore, nil
}
