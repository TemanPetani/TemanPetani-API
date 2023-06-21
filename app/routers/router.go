package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alta/temanpetani/app/middlewares"
	// _userData "alta/temanpetani/features/users/data"
	// _userHandler "alta/temanpetani/features/users/handler"
	// _userService "alta/temanpetani/features/users/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	// UserData := _userData.New(db)
	// UserService := _userService.New(UserData)
	// // UserHandler := _userHandler.New(UserService)


	e.POST("/users", echo.NotFoundHandler)
	e.GET("/users/profile", echo.NotFoundHandler, middlewares.JWTMiddleware())
	e.PUT("/users/profile", echo.NotFoundHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/profile", echo.NotFoundHandler, middlewares.JWTMiddleware())
	e.POST("/login", echo.NotFoundHandler)
} 