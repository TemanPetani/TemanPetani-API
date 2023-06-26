package handler

import (
	"alta/temanpetani/features/templates"
)

type ScheduleTemplateRequest struct {
	Name string `json:"name" form:"name"`
}

func NewScheduleTemplateCore(templateRequest ScheduleTemplateRequest) templates.ScheduleTemplateCore {
	return templates.ScheduleTemplateCore{
		Name: templateRequest.Name,
	}
}
