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
	beratkg, _ := strconv.Atoi(c.FormValue("berat_kg"))
	tanggal_masuk, _ := time.Parse(layoutISO, c.FormValue("tanggal_masuk"))

	newBarang := model.Barang{
		Id:           Id,
		Barang:       barang,
		Berat_kg:     beratkg,
		TanggalMasuk: tanggal_masuk,
	}

	// Insert Di tabel usulan
	if dbc := db.Debug().Create(&newBarang); dbc.Error != nil {
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
	beratkg, _ := strconv.Atoi("berat_kg")
	tanggalmasuk, _ := time.Parse(layoutISO, c.FormValue("tanggal_masuk"))
	// Barang := model.Golongan{}

	newBarang := model.Barang{
		Barang:       barang,
		Berat_kg:     beratkg,
		TanggalMasuk: tanggalmasuk,
	}

	if dbc := db.Debug().Model(&newBarang).Update(newBarang).Where("id = ?", Id); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Barang Berhasil Diupdate",
	})

}

func GetBarang(c echo.Context) error {

	db := db.Manager()
	if c.QueryParam("id") != "" {
		db = db.Where("id = ?", c.QueryParam("id"))
	}
	if c.QueryParam("barang") != "" {
		db = db.Where("barang = ?", c.QueryParam("barang"))
	}

	Barang := []model.Barang{}
	rows := db.Find(&Barang)
	return c.JSON(http.StatusOK, rows)

}
