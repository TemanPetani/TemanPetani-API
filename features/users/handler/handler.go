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
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errBind.Error()))
	}

	dataLogin, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "Email") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "Password") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
		}
	}

	response := NewAuthResponse(dataLogin, token)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Berhasil Login", response))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errBind.Error()))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "Password") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "Duplicate") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Berhasil Membuat Pengguna Baru"))
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(errExtract.Error()))
	}

	result, err := handler.userService.GetById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse(err.Error()))
	}

	userResponse := NewUserResponse(result)
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Berhasil Mendapatkan Data Pengguna", userResponse))
}

func (handler *UserHandler) UpdateUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(errExtract.Error()))
	}

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errBind.Error()))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.UpdateById(userId, userCore)
	if err != nil {
		if strings.Contains(err.Error(), "Password") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "Duplicate") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusNotFound, helpers.FailedResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Berhasil Memperbarui Data Pengguna"))
}

func (handler *UserHandler) DeleteUserById(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(errExtract.Error()))
	}

	err := handler.userService.DeleteById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Berhasil Menghapus Data Pengguna"))
}

func (handler *UserHandler) UpdateUserPicture(c echo.Context) error {
	userId, _, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(errExtract.Error()))
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	errUpdate := handler.userService.UpdateImage(userId, picture)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Berhasil Memperbarui Data Pengguna"))
}
