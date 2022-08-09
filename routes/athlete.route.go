package routes

import (
	"locathlete-server/controllers"

	"github.com/labstack/echo/v4"
)

func AthleteRoute(e *echo.Echo) {
	e.GET("/athletes", controllers.GetAthletes)
}
