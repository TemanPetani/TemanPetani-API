package handler

import (
	"alta/temanpetani/features/templates"
)

type ScheduleTemplateRequest struct {
	Name string `json:"name" form:"name"`
}

type TaskTemplateRequest struct {
	Name      string `json:"name" form:"name"`
	StartDays uint   `json:"startDays" form:"startDays"`
}

func NewScheduleTemplateCore(templateRequest ScheduleTemplateRequest) templates.ScheduleTemplateCore {
	return templates.ScheduleTemplateCore{
		Name: templateRequest.Name,
	}
}

func NewTaskTemplateCore(templateRequest TaskTemplateRequest) templates.TaskTemplateCore {
	return templates.TaskTemplateCore{
		Name:      templateRequest.Name,
		StartDays: templateRequest.StartDays,
	}
}
