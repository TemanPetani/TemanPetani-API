package handler

import (
	"alta/temanpetani/features/templates"
)

type ScheduleTemplateResponse struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TasksTemplateResponse struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDays uint   `json:"startDays,omitempty"`
}

func NewScheduleTemplateResponse(template templates.ScheduleTemplateCore) ScheduleTemplateResponse {
	return ScheduleTemplateResponse{
		ID:   template.ID,
		Name: template.Name,
	}
}

func NewTaskTemplateResponse(template templates.TaskTemplateCore) TasksTemplateResponse {
	return TasksTemplateResponse{
		ID:        template.ID,
		Name:      template.Name,
		StartDays: template.StartDays,
	}
}
