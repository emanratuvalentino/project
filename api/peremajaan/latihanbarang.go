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

func SimpanBarang(c echo.Context) error {
	db := db.Manager()

	u1 := uuid.Must(uuid.NewV4())
	Id := u1.String()
	// configuration := config.GetConfig()
	layoutISO := "2006-01-02"

	barang := c.FormValue("barang")
	berat, _ := strconv.Atoi(c.FormValue("berat"))
	tanggal_masuk, _ := time.Parse(layoutISO, c.FormValue("tanggal_masuk"))

	newPersonal := model.Personal1{
		Id:           Id,
		Barang:       barang,
		Berat:        berat,
		TanggalMasuk: tanggal_masuk,
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
		"message": "Data Usulan Barang Berhasil Diupdate",
		"Id":      Id,
	})

}

func UpdateBarang(c echo.Context) error {
	db := db.Manager()
	Id := c.FormValue("id")
	// configuration := config.GetConfig()
	layoutISO := "2006-01-02"

	barang := c.FormValue("barang")
	berat, _ := strconv.Atoi("berat")
	tanggal_masuk, _ := time.Parse(layoutISO, c.FormValue("tanggal_masuk"))
	// Golongan := model.Golongan{}

	newPersonal := model.Personal1{
		Barang:       barang,
		Berat:        berat,
		TanggalMasuk: tanggal_masuk,
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

func GetPersonal1(c echo.Context) error {

	db := db.Manager()
	if c.QueryParam("id") != "" {
		db = db.Where("id = ?", c.QueryParam("id"))
	}
	if c.QueryParam("nama") != "" {
		db = db.Where("nama = ?", c.QueryParam("nama"))
	}

	Personal := []model.Personal1{}
	rows := db.Find(&Personal)
	return c.JSON(http.StatusOK, rows)

}
