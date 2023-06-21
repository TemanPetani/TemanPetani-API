package main

import (
	"alta/temanpetani/app/config"
	"alta/temanpetani/app/database"
	"alta/temanpetani/app/migration"
	"alta/temanpetani/app/routers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.ReadEnv()
	database := database.InitDB(config)
	if errMigrate := migration.InitMigration(database); errMigrate != nil {
		log.Fatal(errMigrate.Error())
	}

	echo := echo.New()
	echo.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	echo.Pre(middleware.RemoveTrailingSlash())
	echo.Use(middleware.CORS())

	routers.InitRouters(database, echo)
	
	echo.Logger.Fatal(echo.Start(":8080"))
}	