package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Docudigital API v0.1.0")

}
