package routelist

import (
	"code-be-docudigital/api"

	"github.com/labstack/echo/v4"
)

func Homeroute(echo *echo.Echo) {
	echo.GET("/", api.Home)
}
