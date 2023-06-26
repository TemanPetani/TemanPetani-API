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
