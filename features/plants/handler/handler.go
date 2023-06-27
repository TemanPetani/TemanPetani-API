package handler

import (
	"alta/temanpetani/features/plants"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"net/http"
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

	plantCore := NewScheduleCore(plantInput)
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
