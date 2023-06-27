package handler

import (
	"alta/temanpetani/features/plants"
	"time"
)

type ScheduleRequest struct {
	TemplateID uint64    `json:"templateId" form:"templateId"`
	Name       string    `json:"name" form:"name"`
	StartDate  time.Time `json:"startDate" form:"startDate"`
}

func NewScheduleCore(plantRequest ScheduleRequest) plants.ScheduleCore {
	return plants.ScheduleCore{
		TemplateID: plantRequest.TemplateID,
		Name:       plantRequest.Name,
		StartDate:  plantRequest.StartDate,
	}
}
