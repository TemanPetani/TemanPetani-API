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

func (service *plantService) GetAllSchedule() ([]plants.ScheduleCore, error) {
	data, err := service.plantData.SelectAllSchedule()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *plantService) GetAllFarmerSchedule(farmerId uint64) ([]plants.ScheduleCore, error) {
	data, err := service.plantData.SelectAllFarmerSchedule(farmerId)
	if err != nil {
		return nil, err
	}

	for i, value := range data {
		task, errSelect := service.plantData.SelectRecentTask(value.ID)
		if errSelect != nil {
			return nil, errSelect
		}
		data[i].Tasks = append(data[i].Tasks, task)
	}
	return data, err
}

func (service *plantService) GetScheduleById(id uint64) (plants.ScheduleCore, error) {
	data, err := service.plantData.SelectScheduleById(id)
	if err != nil {
		return plants.ScheduleCore{}, err
	}

	tasks, err := service.plantData.SelectAllTasks(data.ID)
	if err != nil {
		return plants.ScheduleCore{}, err
	}

	data.Tasks = tasks

	return data, err
}
