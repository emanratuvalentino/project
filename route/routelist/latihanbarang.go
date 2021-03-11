package routelist

import (
	"code-be-docudigital/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func LatihanBarang(echo *echo.Group) {

	echo.POST("/latihanbarang/simpan", peremajaan.SimpanBarang)
	echo.POST("/latihanbarang/update", peremajaan.UpdateBarang)
	echo.GET("/latihanbarang", peremajaan.GetPersonal1)

}
