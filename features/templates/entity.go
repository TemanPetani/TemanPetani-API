package templates

type ScheduleTemplateCore struct {
	ID    uint64
	Name  string `validate:"required"`
	Tasks []TaskTemplateCore
}

type TaskTemplateCore struct {
	ID         uint64
	ScheduleID uint64 `validate:"required"`
	Name       string `validate:"required"`
	StartDays  uint   `validate:"required"`
}

type TemplateDataInterface interface {
	InsertSchedule(input ScheduleTemplateCore) error
	InsertTask(input TaskTemplateCore) error
}

type TemplateServiceInterface interface {
	CreateSchedule(input ScheduleTemplateCore) error
	CreateTask(input TaskTemplateCore) error
}
