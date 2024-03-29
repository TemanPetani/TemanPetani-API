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
		return plants.ScheduleCore{}, errors.New("Gagal Menambahkan Data Tanaman")
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
		return errors.New("Gagal Menambahkan Data Tanaman")
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

func (repo *plantQuery) SelectRecentTask(scheduleId uint64) (plants.TaskCore, error) {
	var plantsData Task
	tx := repo.db.
		Where("schedule_id = ?", scheduleId).
		Where("start_date between start_date and curdate()").
		Order("start_date desc").
		First(&plantsData)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			tx = repo.db.Where("schedule_id = ?", scheduleId).First(&plantsData)
			if tx.Error != nil {
				return plants.TaskCore{}, tx.Error
			}
		} else {
			return plants.TaskCore{}, tx.Error
		}
	}

	plantCore := NewTaskCore(plantsData)

	return plantCore, nil
}

func (repo *plantQuery) SelectScheduleById(id uint64) (plants.ScheduleCore, error) {
	var plantGorm Schedule
	tx := repo.db.First(&plantGorm, id)
	if tx.Error != nil {
		return plants.ScheduleCore{}, tx.Error
	}

	plantCore := NewScheduleCore(plantGorm)
	return plantCore, nil
}

func (repo *plantQuery) SelectTasksNotification(userId uint64) ([]plants.TaskCore, error) {
	query := "select t.* from tasks as t " +
		"inner join schedules as s on t.schedule_id = s.id " +
		"inner join users as u on s.user_id = u.id " +
		"where  u.id = ? " +
		"and s.deleted_at is null " +
		"and t.completed_date is null " +
		"and t.start_date between s.start_date and curdate()"

	var plantsData []Task
	tx := repo.db.Raw(query, userId).Find(&plantsData)
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

func (repo *plantQuery) UpdateTaskById(taskId uint64, input plants.TaskCore) error {
	var plantGorm Task
	tx := repo.db.First(&plantGorm, taskId)
	if tx.Error != nil {
		return tx.Error
	}

	plantInputGorm := NewTaskModel(input)
	tx = repo.db.Model(&plantGorm).Updates(plantInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Memperbarui Data Pengguna")
	}

	return nil
}

func (repo *plantQuery) DeleteScheduleById(scheduleId uint64) error {
	var scheduleGorm Schedule
	tx := repo.db.First(&scheduleGorm, scheduleId)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Delete(&scheduleGorm, scheduleId)
	if tx.Error != nil {
		return tx.Error
	}

	var taskGorm Task
	tx = repo.db.Where("schedule_id = ?", scheduleId).Delete(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Gagal Memperbarui Data Pengguna")
	}

	return nil
}
