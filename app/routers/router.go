package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_userData "alta/temanpetani/features/users/data"
	_userHandler "alta/temanpetani/features/users/handler"
	_userService "alta/temanpetani/features/users/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandler := _userHandler.New(UserService)

	e.POST("/login", UserHandler.Login)
	e.POST("/users", UserHandler.CreateUser)
}
