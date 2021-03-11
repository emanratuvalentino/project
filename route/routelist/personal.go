package routelist

import (
	"code-be-docudigital/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func Personal(echo *echo.Group) {

	echo.POST("/personal/simpan", peremajaan.SimpanPersonal)
	echo.POST("/personal/update", peremajaan.UpdatePersonal)
	echo.GET("/personal", peremajaan.GetPersonal)
}
