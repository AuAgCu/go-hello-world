package handlers

import (
	"net/http"

	"app/models"

	"github.com/labstack/echo"
)

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}
