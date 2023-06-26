package service

import (
	"alta/temanpetani/features/templates"

	"github.com/go-playground/validator/v10"
)

type templateService struct {
	templateData templates.TemplateDataInterface
	validate     *validator.Validate
}

func New(repo templates.TemplateDataInterface) templates.TemplateServiceInterface {
	return &templateService{
		templateData: repo,
		validate:     validator.New(),
	}
}

func (service *templateService) CreateSchedule(input templates.ScheduleTemplateCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.templateData.InsertSchedule(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (service *templateService) CreateTask(input templates.TaskTemplateCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.templateData.InsertTask(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (service *templateService) GetAllSchedule() ([]templates.ScheduleTemplateCore, error) {
	data, err := service.templateData.SelectAllSchedule()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *templateService) GetScheduleById(id uint64) (templates.ScheduleTemplateCore, error) {
	data, err := service.templateData.SelectScheduleById(id)
	if err != nil {
		return templates.ScheduleTemplateCore{}, err
	}

	tasks, err := service.templateData.SelectAllTasks(data.ID)
	if err != nil {
		return templates.ScheduleTemplateCore{}, err
	}

	data.Tasks = tasks

	return data, err
}
