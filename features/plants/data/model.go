package data

import (
	"alta/temanpetani/features/plants"
	_templateModel "alta/temanpetani/features/templates/data"
	_userModel "alta/temanpetani/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID         uint64    `gorm:"primarykey"`
	UserID     uint64    `gorm:"notNull"`
	TemplateID uint64    `gorm:"notNull"`
	Name       string    `gorm:"notNull"`
	StartDate  time.Time `gorm:"type:date;notNull"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt                  `gorm:"index"`
	User       _userModel.User                 `gorm:"foreignKey:UserID"`
	Template   _templateModel.ScheduleTemplate `gorm:"foreignKey:TemplateID"`
	Tasks      []Task                          `gorm:"foreignKey:ScheduleID"`
}

type Task struct {
	ID            uint64    `gorm:"primarykey"`
	ScheduleID    uint64    `gorm:"notNull"`
	Name          string    `gorm:"notNull"`
	StartDate     time.Time `gorm:"type:date;notNull"`
	CompletedDate time.Time `gorm:"type:date;default:null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type FarmerSchedule struct {
	ScheduleID   uint64
	FarmerName   string
	ScheduleName string
}

func NewScheduleCore(scheduleData Schedule) plants.ScheduleCore {
	return plants.ScheduleCore{
		ID:         scheduleData.ID,
		TemplateID: scheduleData.TemplateID,
		Name:       scheduleData.Name,
		StartDate:  scheduleData.StartDate,
		Farmer: plants.FarmerCore{
			FarmerID: scheduleData.UserID,
		},
	}
}

func NewScheduleModel(dataCore plants.ScheduleCore) Schedule {
	return Schedule{
		ID:         dataCore.ID,
		UserID:     dataCore.Farmer.FarmerID,
		TemplateID: dataCore.TemplateID,
		Name:       dataCore.Name,
		StartDate:  dataCore.StartDate,
	}
}

func NewTaskCore(taskData Task) plants.TaskCore {
	return plants.TaskCore{
		ID:            taskData.ID,
		ScheduleID:    taskData.ScheduleID,
		Name:          taskData.Name,
		StartDate:     taskData.StartDate,
		CompletedDate: taskData.CompletedDate,
	}
}

func NewTaskModel(dataCore plants.TaskCore) Task {
	return Task{
		ID:            dataCore.ID,
		ScheduleID:    dataCore.ScheduleID,
		Name:          dataCore.Name,
		StartDate:     dataCore.StartDate,
		CompletedDate: dataCore.CompletedDate,
	}
}

func NewFarmerSchedule(scheduleData FarmerSchedule) plants.ScheduleCore {
	return plants.ScheduleCore{
		ID:   scheduleData.ScheduleID,
		Name: scheduleData.ScheduleName,
		Farmer: plants.FarmerCore{
			FarmerName: scheduleData.FarmerName,
		},
	}
}
