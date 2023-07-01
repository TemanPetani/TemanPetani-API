package handler

import (
	"alta/temanpetani/features/plants"
	"time"
)

type ScheduleResponse struct {
	ID        uint64         `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	StartDate time.Time      `json:"startDate,omitempty"`
	Tasks     []TaskResponse `json:"activities,omitempty"`
}

type TaskResponse struct {
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
	var tasksResponse []TaskResponse
	for _, value := range plant.Tasks {
		tasksResponse = append(tasksResponse, NewTaskResponse(value))
	}

	return ScheduleResponse{
		ID:        plant.ID,
		Name:      plant.Name,
		StartDate: plant.StartDate,
		Tasks:     tasksResponse,
	}
}

func NewTaskResponse(plant plants.TaskCore) TaskResponse {
	var completedDate *time.Time
	if plant.CompletedDate.IsZero() {
		completedDate = nil
	} else {
		completedDate = &plant.CompletedDate
	}

	return TaskResponse{
		ID:            plant.ID,
		Name:          plant.Name,
		StartDate:     plant.StartDate,
		CompletedDate: completedDate,
	}
}

func NewFarmerScheduleResponse(plant plants.ScheduleCore) FarmerScheduleResponse {
	return FarmerScheduleResponse{
		ID:           plant.ID,
		FarmerName:   plant.Farmer.FarmerName,
		ScheduleName: plant.Name,
	}
}
