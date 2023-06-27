package plants

import "time"

type ScheduleCore struct {
	ID         uint64
	TemplateID uint64    `validate:"required"`
	Name       string    `validate:"required"`
	StartDate  time.Time `validate:"required"`
	Tasks      []TaskCore
}

type TaskCore struct {
	ID            uint64
	ScheduleID    uint64
	Name          string
	StartDate     time.Time
	CompletedDate time.Time
}

type PlantDataInterface interface {
	InsertSchedule(input ScheduleCore) (ScheduleCore, error)
	InsertTask(input []TaskCore) error
}

type PlantServiceInterface interface {
	CreateSchedule(input ScheduleCore) error
}
