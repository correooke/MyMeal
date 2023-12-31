package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAlive(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server is alive!")
}
