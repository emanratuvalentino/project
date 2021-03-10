package routelist

import (
	"code-be-docudigital/api/peremajaan"

	"github.com/labstack/echo/v4"
)

func Peremajaan(echo *echo.Group) {

	echo.POST("/peremajaan/golongan/simpan-usul", peremajaan.BlankUsulGolonganSave)
	echo.POST("/peremajaan/golongan/update-data", peremajaan.UsulGolonganUpdate)
	echo.POST("/peremajaan/golongan/hapus", peremajaan.UsulDeleteGolongan)



}
