package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	_userData "alta/temanpetani/features/users/data"
	_userHandler "alta/temanpetani/features/users/handler"
	_userService "alta/temanpetani/features/users/service"

	_productData "alta/temanpetani/features/products/data"
	_productHandler "alta/temanpetani/features/products/handler"
	_productService "alta/temanpetani/features/products/service"

	_templateData "alta/temanpetani/features/templates/data"
	_templateHandler "alta/temanpetani/features/templates/handler"
	_templateService "alta/temanpetani/features/templates/service"

	_plantData "alta/temanpetani/features/plants/data"
	_plantHandler "alta/temanpetani/features/plants/handler"
	_plantService "alta/temanpetani/features/plants/service"

	"alta/temanpetani/utils/middlewares"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	initUserRouter(db, e)
	initProductRouter(db, e)
	initTemplateRouter(db, e)
	initPlantsRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)

	e.POST("/login", userHandler.Login)
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/profile", userHandler.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users/profile", userHandler.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/profile", userHandler.DeleteUserById, middlewares.JWTMiddleware())
}

func initProductRouter(db *gorm.DB, e *echo.Echo) {
	productData := _productData.New(db)
	productService := _productService.New(productData)
	productHandler := _productHandler.New(productService)

	e.POST("/products", productHandler.PostProductHandler, middlewares.JWTMiddleware())
	e.GET("/products", productHandler.GetAllProductsHandler)
	e.PUT("/products/:id/images", productHandler.PutImageProductHandler, middlewares.JWTMiddleware())

}

func initTemplateRouter(db *gorm.DB, e *echo.Echo) {
	templateData := _templateData.New(db)
	templateService := _templateService.New(templateData)
	templateHandler := _templateHandler.New(templateService)

	e.POST("/templates", templateHandler.CreateScheduleTemplate, middlewares.JWTMiddleware())
	e.POST("/templates/:id/tasks", templateHandler.CreateTaskTemplate, middlewares.JWTMiddleware())
	e.GET("/templates", templateHandler.GetAllSchedule, middlewares.JWTMiddleware())
	e.GET("/templates/:id", templateHandler.GetScheduleById, middlewares.JWTMiddleware())
	e.PUT("/templates/:id", templateHandler.UpdateScheduleById, middlewares.JWTMiddleware())
	e.PUT("/templates/tasks/:id", templateHandler.UpdateTaskById, middlewares.JWTMiddleware())
	e.DELETE("/templates/:id", templateHandler.DeleteScheduleById, middlewares.JWTMiddleware())
	e.DELETE("/templates/tasks/:id", templateHandler.DeleteTaskById, middlewares.JWTMiddleware())
}

func initPlantsRouter(db *gorm.DB, e *echo.Echo) {
	templateData := _templateData.New(db)
	plantData := _plantData.New(db)
	plantService := _plantService.New(plantData, templateData)
	plantHandler := _plantHandler.New(plantService)

	e.POST("/plants", plantHandler.CreateSchedule, middlewares.JWTMiddleware())
	e.GET("/plants", plantHandler.GetAllSchedule, middlewares.JWTMiddleware())
	e.GET("/plants/:id", plantHandler.GetScheduleById, middlewares.JWTMiddleware())
	e.GET("/users/plants", plantHandler.GetAllFarmerSchedule, middlewares.JWTMiddleware())
}
