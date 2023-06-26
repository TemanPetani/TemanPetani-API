package handler

import (
	"alta/temanpetani/features/templates"
	"alta/temanpetani/utils/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type templateHandler struct {
	templateService templates.TemplateServiceInterface
}

func New(service templates.TemplateServiceInterface) *templateHandler {
	return &templateHandler{
		templateService: service,
	}
}

func (handler *templateHandler) CreateScheduleTemplate(c echo.Context) error {
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

func (handler *templateHandler) CreateTaskTemplate(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	templateInput := TaskTemplateRequest{}
	errBind := c.Bind(&templateInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	templateCore := NewTaskTemplateCore(templateInput)
	templateCore.ScheduleID = scheduleId
	err := handler.templateService.CreateTask(templateCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data, "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}

func (handler *templateHandler) GetAllSchedule(c echo.Context) error {
	results, err := handler.templateService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	var templatesResponse []ScheduleTemplateResponse
	for _, value := range results {
		templatesResponse = append(templatesResponse, NewScheduleTemplateResponse(value))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", templatesResponse))
}

func (handler *templateHandler) GetScheduleById(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	result, err := handler.templateService.GetScheduleById(scheduleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	scheduleResponse := NewScheduleTemplateResponse(result)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", scheduleResponse))
}

func (handler *templateHandler) UpdateScheduleById(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	templateInput := ScheduleTemplateRequest{}
	errBind := c.Bind(&templateInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	templateCore := NewScheduleTemplateCore(templateInput)
	err := handler.templateService.UpdateScheduleById(scheduleId, templateCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error update data, "+err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *templateHandler) DeleteScheduleById(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	err := handler.templateService.DeleteScheduleById(scheduleId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}

func (handler *templateHandler) UpdateTaskById(c echo.Context) error {
	paramId := c.Param("id")
	taskId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	templateInput := TaskTemplateRequest{}
	errBind := c.Bind(&templateInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	templateCore := NewTaskTemplateCore(templateInput)
	err := handler.templateService.UpdateTaskById(taskId, templateCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error update data, "+err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *templateHandler) DeleteTaskById(c echo.Context) error {
	paramId := c.Param("id")
	taskId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	err := handler.templateService.DeleteTaskById(taskId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
