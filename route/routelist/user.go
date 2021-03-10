package routelist

import (
	"code-be-docudigital/api"

	"github.com/labstack/echo/v4"
)

func Userroute(echo *echo.Group) {

	/*-------------------------------------------------
		EXAMPLE ROUTES
		Both of this routes are examples only. DO NOT
		rewrite those. Comment them for production!
		examples are:
		1. Get users data from db
		2. Get Token payload data (multiple example)
		Check both of those API for the detail
	--------------------------------------------------*/
	echo.GET("/users", api.GetUsers)

	echo.GET("/jwt-payload", api.GetJWTPayload)
}
