package peremajaan

import (
	"code-be-docudigital/db"
	"time"

	// "fmt"

	"code-be-docudigital/model"

	// "github.com/fatih/structs"
	"net/http"
	"strconv"

	// "errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)


func SimpanPersonal(c echo.Context) error {
	db := db.Manager()

	u1 := uuid.Must(uuid.NewV4())
	Id := u1.String()
	// configuration := config.GetConfig()
	layoutISO := "2006-01-02"

	nama := c.FormValue("nama")
	usia,_ := strconv.Atoi("usia")
	tanggal_lahir, _ := time.Parse(layoutISO, c.FormValue("tanggal_lahir"))

	newPersonal := model.Personal{
		Id    			:  Id,
		Nama  			:  nama,
		Usia  			:  usia,
		TanggalLahir    :  tanggal_lahir,
	}

	// Insert Di tabel usulan
	if dbc := db.Debug().Create(&newPersonal); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}


	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Golongan Berhasil Diupdate",
		"Id": Id,
	})

}



func UpdatePersonal(c echo.Context) error {
	db := db.Manager()
	Id := c.FormValue("id")
	// configuration := config.GetConfig()
	layoutISO := "2006-01-02"

	nama := c.FormValue("sk_tanggal")
	usia,_ := strconv.Atoi("usia")
	tanggal_lahir, _ := time.Parse(layoutISO, c.FormValue("tanggal_lahir"))
	// Golongan := model.Golongan{}


	newPersonal := model.Personal{
		Nama  			:  nama,
		Usia  			:  usia,
		TanggalLahir    :  tanggal_lahir,
	}

	if dbc := db.Debug().Model(&newPersonal).Updates(newPersonal).Where("id = ?", Id); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}


	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Golongan Berhasil Diupdate",
	})

}

func GetPersonal(c echo.Context) error {

	db := db.Manager()
	if c.QueryParam("id") != "" {
		db = db.Where("id = ?", c.QueryParam("id"))
	}
	if c.QueryParam("nama") != "" {
		db = db.Where("nama = ?", c.QueryParam("nama"))
	}

	Personal := []model.Personal{}
	rows := db.Find(&Personal)
	return c.JSON(http.StatusOK, rows)

}

