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

type TaskRequest struct {
	CompletedDate time.Time `json:"completedDate" form:"completedDate"`
}

func NewScheduleRequest(request ScheduleRequest) plants.ScheduleCore {
	return plants.ScheduleCore{
		TemplateID: request.TemplateID,
		Name:       request.Name,
		StartDate:  request.StartDate,
	}
}

func NewTaskRequest(request TaskRequest) plants.TaskCore {
	return plants.TaskCore{
		CompletedDate: request.CompletedDate,
	}
}
