package handler

import (
	"alta/temanpetani/features/plants"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type plantHandler struct {
	plantService plants.PlantServiceInterface
}

func New(service plants.PlantServiceInterface) *plantHandler {
	return &plantHandler{
		plantService: service,
	}
}

func (handler *plantHandler) CreateSchedule(c echo.Context) error {
	plantInput := ScheduleRequest{}
	errBind := c.Bind(&plantInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	plantCore := NewScheduleRequest(plantInput)
	plantCore.Farmer.FarmerID = userId
	err := handler.plantService.CreateSchedule(plantCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data, "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}

func (handler *plantHandler) GetAllSchedule(c echo.Context) error {
	results, err := handler.plantService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	var plantsResponse []FarmerScheduleResponse
	for _, value := range results {
		plantsResponse = append(plantsResponse, NewFarmerScheduleResponse(value))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", plantsResponse))
}

func (handler *plantHandler) GetAllFarmerSchedule(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	results, err := handler.plantService.GetAllFarmerSchedule(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	var plantsResponse []ScheduleResponse
	for _, value := range results {
		plantsResponse = append(plantsResponse, NewScheduleResponse(value))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", plantsResponse))
}

func (handler *plantHandler) GetScheduleById(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	result, err := handler.plantService.GetScheduleById(scheduleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	scheduleResponse := NewScheduleResponse(result)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", scheduleResponse))
}

func (handler *plantHandler) GetTasksNotification(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	results, err := handler.plantService.GetTasksNotification(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error read data, "+err.Error()))
	}

	var plantsResponse []TaskResponse
	for _, value := range results {
		plantsResponse = append(plantsResponse, NewTaskResponse(value))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", plantsResponse))
}

func (handler *plantHandler) UpdateTaskById(c echo.Context) error {
	paramId := c.Param("id")
	taskId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	plantInput := TaskRequest{}
	errBind := c.Bind(&plantInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	plantCore := NewTaskRequest(plantInput)
	err := handler.plantService.UpdateTaskById(taskId, plantCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error update data, "+err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *plantHandler) DeleteScheduleById(c echo.Context) error {
	paramId := c.Param("id")
	scheduleId, errParse := strconv.ParseUint(paramId, 10, 64)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error parse data"))
	}

	err := handler.plantService.DeleteScheduleById(scheduleId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
