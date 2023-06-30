package plants

import (
	"time"
)

type ScheduleCore struct {
	ID         uint64
	TemplateID uint64    `validate:"required"`
	Name       string    `validate:"required"`
	StartDate  time.Time `validate:"required"`
	Farmer     FarmerCore
	Tasks      []TaskCore
}

type TaskCore struct {
	ID            uint64
	ScheduleID    uint64
	Name          string
	StartDate     time.Time
	CompletedDate time.Time
}

type FarmerCore struct {
	FarmerID   uint64
	FarmerName string
}

type PlantDataInterface interface {
	InsertSchedule(input ScheduleCore) (ScheduleCore, error)
	InsertTask(input []TaskCore) error
	SelectAllSchedule() ([]ScheduleCore, error)
	SelectAllFarmerSchedule(farmerId uint64) ([]ScheduleCore, error)
	SelectAllTasks(scheduleId uint64) ([]TaskCore, error)
	SelectRecentTask(scheduleId uint64) (TaskCore, error)
	SelectScheduleById(scheduleId uint64) (ScheduleCore, error)
	SelectTasksNotification(userId uint64) ([]TaskCore, error)
	UpdateTaskById(taskId uint64, input TaskCore) error
	DeleteScheduleById(scheduleId uint64) error
}

type PlantServiceInterface interface {
	CreateSchedule(input ScheduleCore) error
	GetAllSchedule() ([]ScheduleCore, error)
	GetAllFarmerSchedule(farmerId uint64) ([]ScheduleCore, error)
	GetScheduleById(scheduleId uint64) (ScheduleCore, error)
	GetTasksNotification(userId uint64) ([]TaskCore, error)
	UpdateTaskById(taskId uint64, input TaskCore) error
	DeleteScheduleById(scheduleId uint64) error
}
