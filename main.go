package main

import (
	"alta/temanpetani/app/config"
	"alta/temanpetani/app/database"
	"alta/temanpetani/app/routers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config := config.ReadEnv()
	db := database.InitDB(config)
	if errMigrate := database.InitMigration(db); errMigrate != nil {
		log.Fatal(errMigrate.Error())
	}
	routers.InitRouters(db, e)

	e.Logger.Fatal(e.Start(":8080"))
}
