package routelist

import (
	"code-be-docudigital/api/auth"

	"github.com/labstack/echo/v4"
)

func Authenticationroute(echo *echo.Echo) {
	echo.GET("/auth/client-credentials", auth.GetClientCredentials)

	echo.GET("/auth/tokens", auth.GetToken)

}
