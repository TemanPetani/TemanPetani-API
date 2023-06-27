package handler

import (
	"alta/temanpetani/features/plants"
	"time"
)

type ScheduleResponse struct {
	ID        uint64          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	StartDate time.Time       `json:"startDate,omitempty"`
	Tasks     []TasksResponse `json:"tasks,omitempty"`
}

type TasksResponse struct {
	ID            uint64    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty"`
	CompletedDate time.Time `json:"completedDate,omitempty"`
}

func NewScheduleResponse(plant plants.ScheduleCore) ScheduleResponse {
	var tasksResponse []TasksResponse
	for _, value := range plant.Tasks {
		tasksResponse = append(tasksResponse, TasksResponse{
			ID:            value.ID,
			Name:          value.Name,
			StartDate:     value.StartDate,
			CompletedDate: value.CompletedDate,
		})
	}

	return ScheduleResponse{
		ID:        plant.ID,
		Name:      plant.Name,
		StartDate: plant.StartDate,
		Tasks:     tasksResponse,
	}
}

func NewTaskplantResponse(plant plants.TaskCore) TasksResponse {
	return TasksResponse{
		ID:            plant.ID,
		Name:          plant.Name,
		StartDate:     plant.StartDate,
		CompletedDate: plant.CompletedDate,
	}
}
