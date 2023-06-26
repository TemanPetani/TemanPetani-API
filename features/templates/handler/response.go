package handler

import (
	"alta/temanpetani/features/templates"
)

type ScheduleTemplateResponse struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewScheduleTemplateResponse(template templates.ScheduleTemplateCore) ScheduleTemplateResponse {
	return ScheduleTemplateResponse{
		ID:   template.ID,
		Name: template.Name,
	}
}
