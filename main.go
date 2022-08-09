package main

import (
	"locathlete-server/configs"
	"locathlete-server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.Init()
	defer configs.DbClose()
	routes.AthleteRoute(e)
	e.Logger.Fatal(e.Start(":" + configs.GetPort()))
}
