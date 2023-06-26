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
	SelectAllSchedule() ([]ScheduleTemplateCore, error)
	SelectAllTasks(scheduleId uint64) ([]TaskTemplateCore, error)
	SelectScheduleById(id uint64) (ScheduleTemplateCore, error)
	UpdateScheduleById(id uint64, input ScheduleTemplateCore) error
	DeleteScheduleById(id uint64) error
	UpdateTaskById(id uint64, input TaskTemplateCore) error
	DeleteTaskById(id uint64) error
}

type TemplateServiceInterface interface {
	CreateSchedule(input ScheduleTemplateCore) error
	CreateTask(input TaskTemplateCore) error
	GetAllSchedule() ([]ScheduleTemplateCore, error)
	GetScheduleById(id uint64) (ScheduleTemplateCore, error)
	UpdateScheduleById(id uint64, input ScheduleTemplateCore) error
	DeleteScheduleById(id uint64) error
	UpdateTaskById(id uint64, input TaskTemplateCore) error
	DeleteTaskById(id uint64) error
}
