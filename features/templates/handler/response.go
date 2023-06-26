package handler

import (
	"alta/temanpetani/features/templates"
)

type ScheduleTemplateResponse struct {
	ID    uint64                  `json:"id,omitempty"`
	Name  string                  `json:"name,omitempty"`
	Tasks []TasksTemplateResponse `json:"tasks,omitempty"`
}

type TasksTemplateResponse struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDays uint   `json:"startDays,omitempty"`
}

func NewScheduleTemplateResponse(template templates.ScheduleTemplateCore) ScheduleTemplateResponse {
	var tasksResponse []TasksTemplateResponse
	for _, value := range template.Tasks {
		tasksResponse = append(tasksResponse, TasksTemplateResponse{
			ID:        value.ID,
			Name:      value.Name,
			StartDays: value.StartDays,
		})
	}

	return ScheduleTemplateResponse{
		ID:    template.ID,
		Name:  template.Name,
		Tasks: tasksResponse,
	}
}

func NewTaskTemplateResponse(template templates.TaskTemplateCore) TasksTemplateResponse {
	return TasksTemplateResponse{
		ID:        template.ID,
		Name:      template.Name,
		StartDays: template.StartDays,
	}
}
