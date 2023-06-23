package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	_userData "alta/temanpetani/features/users/data"
	_userHandler "alta/temanpetani/features/users/handler"
	_userService "alta/temanpetani/features/users/service"
	"alta/temanpetani/utils/middlewares"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	initUserRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)

	e.POST("/login", userHandler.Login)
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users", userHandler.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandler.DeleteUserById, middlewares.JWTMiddleware())
}
