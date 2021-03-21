package routelist

import (
	"code-be-docudigital/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func Barang(echo *echo.Group) {

	echo.POST("/latihanbarang/simpan", peremajaan.SimpanBarang)
	echo.POST("/latihanbarang/update", peremajaan.UpdateBarang)
	echo.GET("/latihanbarang", peremajaan.GetBarang)
	echo.PUT("/latihanbarang/simpanbarang/", peremajaan.SimpanBarang)
	echo.PUT("/latihanbarang/update", peremajaan.UpdateBarang)
	echo.DELETE("/latihanbarang/delete", peremajaan.DeleteBarang)
}
