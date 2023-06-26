package data

import (
	"alta/temanpetani/features/templates"
	"time"

	"gorm.io/gorm"
)

type ScheduleTemplate struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string `gorm:"unique;notNull"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Tasks     []TaskTemplate `gorm:"foreignKey:ScheduleID"`
}

type TaskTemplate struct {
	ID         uint64 `gorm:"primarykey"`
	ScheduleID uint64
	Name       string `gorm:"unique;notNull"`
	StartDays  uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func NewScheduleTemplateCore(scheduleData ScheduleTemplate) templates.ScheduleTemplateCore {
	return templates.ScheduleTemplateCore{
		ID:   scheduleData.ID,
		Name: scheduleData.Name,
	}
}

func NewScheduleTemplateModel(dataCore templates.ScheduleTemplateCore) ScheduleTemplate {
	return ScheduleTemplate{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func NewTaskTemplateCore(taskData TaskTemplate) templates.TaskTemplateCore {
	return templates.TaskTemplateCore{
		ID:         taskData.ID,
		ScheduleID: taskData.ID,
		Name:       taskData.Name,
		StartDays:  taskData.StartDays,
	}
}

func NewTaskTemplateModel(dataCore templates.TaskTemplateCore) TaskTemplate {
	return TaskTemplate{
		ID:         dataCore.ID,
		ScheduleID: dataCore.ScheduleID,
		Name:       dataCore.Name,
		StartDays:  dataCore.StartDays,
	}
}
