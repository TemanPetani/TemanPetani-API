package handler

import (
	"alta/temanpetani/features/users"
	"alta/temanpetani/utils/helpers"
	"alta/temanpetani/utils/middlewares"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	dataLogin, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error login, internal server error"))
		}
	}

	response := NewAuthResponse(dataLogin, token)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("login successful", response))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error insert data, "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success insert data"))
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	result, err := handler.userService.GetById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error read data, "+err.Error()))
	}

	userResponse := NewUserResponse(result)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success read data", userResponse))
}

func (handler *UserHandler) UpdateUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.UpdateById(userId, userCore)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error update data, "+err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *UserHandler) DeleteUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	err := handler.userService.DeleteById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("error delete data, "+err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}

func (handler *UserHandler) UpdateUserPicture(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error read data, "+errExtract.Error()))
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		return helpers.StatusBadRequestResponse(c, "error get file image: "+err.Error())
	}

	errUpdate := handler.userService.UpdateImage(userId, picture)
	if errUpdate != nil {
		return helpers.StatusInternalServerError(c, errUpdate.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}
