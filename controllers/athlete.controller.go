package controllers

import (
	"net/http"

	"locathlete-server/services"

	"github.com/labstack/echo/v4"
)

func GetAthletes(c echo.Context) error {
	athletes, err := services.FindPaginatedAtheletes(c.Request().Context())
	if err != nil {
		c.Logger().Fatal("Error", err)
	}
	return c.JSON(http.StatusOK, athletes)
}
