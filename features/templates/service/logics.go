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
