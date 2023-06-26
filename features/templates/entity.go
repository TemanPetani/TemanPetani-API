package templates

type ScheduleTemplateCore struct {
	ID    uint64
	Name  string
	Tasks []TaskTemplateCore
}

type TaskTemplateCore struct {
	ID         uint64
	ScheduleID uint64
	Name       string
	StartDays  uint
}

type TemplateDataInterface interface {
	InsertSchedule(input ScheduleTemplateCore) error
}

type TemplateServiceInterface interface {
	CreateSchedule(input ScheduleTemplateCore) error
}
