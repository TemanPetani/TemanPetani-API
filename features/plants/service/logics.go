package service

import (
	"alta/temanpetani/features/plants"
	"alta/temanpetani/features/templates"

	"github.com/go-playground/validator/v10"
)

type plantService struct {
	plantData    plants.PlantDataInterface
	templateData templates.TemplateDataInterface
	validate     *validator.Validate
}

func New(plantRepo plants.PlantDataInterface, templateRepo templates.TemplateDataInterface) plants.PlantServiceInterface {
	return &plantService{
		plantData:    plantRepo,
		templateData: templateRepo,
		validate:     validator.New(),
	}
}

func NewTaskCore(scheduleData plants.ScheduleCore, taskData templates.TaskTemplateCore) plants.TaskCore {
	return plants.TaskCore{
		ScheduleID: scheduleData.ID,
		Name:       taskData.Name,
		StartDate:  scheduleData.StartDate.AddDate(0, 0, int(taskData.StartDays)),
	}
}

func (service *plantService) CreateSchedule(input plants.ScheduleCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	scheduleData, errInsert := service.plantData.InsertSchedule(input)
	if errInsert != nil {
		return errInsert
	}

	results, errSelect := service.templateData.SelectAllTasks(input.TemplateID)
	if errSelect != nil {
		return errSelect
	}

	var tasksInput []plants.TaskCore
	for _, value := range results {
		task := NewTaskCore(scheduleData, value)
		tasksInput = append(tasksInput, task)
	}

	errTasksInsert := service.plantData.InsertTask(tasksInput)
	if errTasksInsert != nil {
		return errTasksInsert
	}

	return nil
}
