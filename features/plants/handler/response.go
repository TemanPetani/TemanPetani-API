package handler

import (
	"alta/temanpetani/features/plants"
	"time"
)

type ScheduleResponse struct {
	ID        uint64          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	StartDate time.Time       `json:"startDate,omitempty"`
	Tasks     []TasksResponse `json:"activities,omitempty"`
}

type TasksResponse struct {
	ID            uint64     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	StartDate     time.Time  `json:"startDate,omitempty"`
	CompletedDate *time.Time `json:"completedDate"`
}

type FarmerScheduleResponse struct {
	ID           uint64 `json:"id,omitempty"`
	FarmerName   string `json:"farmerName,omitempty"`
	ScheduleName string `json:"scheduleName,omitempty"`
}

func NewScheduleResponse(plant plants.ScheduleCore) ScheduleResponse {
	var tasksResponse []TasksResponse
	for i, value := range plant.Tasks {
		tasksResponse = append(tasksResponse, TasksResponse{
			ID:            value.ID,
			Name:          value.Name,
			StartDate:     value.StartDate,
			CompletedDate: &value.CompletedDate,
		})
		if tasksResponse[i].CompletedDate.IsZero() {
			tasksResponse[i].CompletedDate = nil
		}
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
		CompletedDate: &plant.CompletedDate,
	}
}

func NewFarmerScheduleResponse(plant plants.ScheduleCore) FarmerScheduleResponse {
	return FarmerScheduleResponse{
		ID:           plant.ID,
		FarmerName:   plant.Farmer.FarmerName,
		ScheduleName: plant.Name,
	}
}
