package handler

import (
	"alta/temanpetani/features/templates"
	"alta/temanpetani/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateHandler struct {
	templateService templates.TemplateServiceInterface
}

func New(service templates.TemplateServiceInterface) *TemplateHandler {
	return &TemplateHandler{
		templateService: service,
	}
}

func (handler *TemplateHandler) CreateScheduleTemplate(c echo.Context) error {
	templateInput := ScheduleTemplateRequest{}
	errBind := c.Bind(&templateInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	templateCore := NewScheduleTemplateCore(templateInput)
	err := handler.templateService.CreateSchedule(templateCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data, "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}
