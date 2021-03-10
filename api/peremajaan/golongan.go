package peremajaan

import (
	"code-be-docudigital/db"
	"time"

	// "fmt"
	"encoding/json"
	"code-be-docudigital/helper"
	"code-be-docudigital/model"
	"code-be-docudigital/config"
	// "github.com/fatih/structs"
	"net/http"
	"strconv"

	// "errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

/*
* endpoint: /peremajaan/golongan/simpan-usul-old
* parameter:
*  mk_golongan_tahun:numeric  / required
*  mk_golongan_bulan:numeric / required
*  sk_nomor:  required
*  sk_tanggal:yyyy-mm-dd  required
*  tanggal_letter_bkn:yyyy-mm-dd  required
*  tmt_golongan:yyyy-mm-dd  required
*  nomor_letter_bkn:  required
*  jenis_kp_id:  required
*  pns_orang_id:  required
*  Golongan ID:  required
 */
func UsulGolonganSave(c echo.Context) error {
	db := db.Manager()
	// configuration := config.GetConfig()
	layoutISO := "2006-01-02"

	// Form Value, Di;ihat dari kolom request field yang ada pada spreadsheet
	mk_golongan_tahun, err := strconv.Atoi(c.FormValue("mk_golongan_tahun"))
	mk_golongan_bulan, err := strconv.Atoi(c.FormValue("mk_golongan_bulan"))
	sk_nomor := c.FormValue("sk_nomor")
	nomor_letter_bkn := c.FormValue("nomor_letter_bkn")
	jenis_kp_id := c.FormValue("jenis_kp_id")
	golongan_id := c.FormValue("golongan_id")
	pns_orang_id := c.FormValue("pns_orang_id")
	sk_tanggal, _ := time.Parse(layoutISO, c.FormValue("sk_tanggal"))
	tanggal_letter_bkn, _ := time.Parse(layoutISO, c.FormValue("tanggal_letter_bkn"))
	tmt_golongan, _ := time.Parse(layoutISO, c.FormValue("tmt_golongan"))
	// Golongan := model.Golongan{}
	if err != nil {
		response := map[string]interface{}{
			"error":   true,
			"message": err,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	u1 := uuid.Must(uuid.NewV4())
	ustrng := u1.String()

	newGolonganUsul := model.Golongan{
		Id:                 ustrng,
		Mk_golongan_tahun:  mk_golongan_tahun,
		Mk_golongan_bulan:  mk_golongan_bulan,
		Sk_nomor:           sk_nomor,
		Sk_tanggal:         sk_tanggal,
		Tanggal_letter_bkn: tanggal_letter_bkn,
		Tmt_golongan:       tmt_golongan,
		Nomor_letter_bkn:   nomor_letter_bkn,
		Jenis_kp_id:        jenis_kp_id,
		Pns_orang_id:       pns_orang_id,
		Golongan_id:        golongan_id,
	}

	jenisLayanan := "1"
	instansiId := "1"
	statusUsulan := 2

	newUsulanMonitoring := model.Usulan{
		IDUsulan:     u1,
		JenisLayanan: jenisLayanan,
		PnsID:        pns_orang_id,
		InstansiID:   instansiId,
		StatusUsulan: statusUsulan,
		TahapanUsulan: 	"Input Berkas",
		StatusUsulanNama: "Belum Lengkap",
	}

	// Insert Di tabel usulan
	if dbc := db.Debug().Create(&newUsulanMonitoring); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	if dbc := db.Debug().Create(&newGolonganUsul); dbc.Error != nil {
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

/*
* endpoint: /peremajaan/golongan/update-data
* parameter:
*  usulan_id:  required
*  mk_golongan_tahun:numeric  / required
*  mk_golongan_bulan:numeric / required
*  sk_nomor:  required
*  sk_tanggal:yyyy-mm-dd  required
*  tanggal_letter_bkn:yyyy-mm-dd  required
*  tmt_golongan:yyyy-mm-dd  required
*  nomor_letter_bkn:  required
*  jenis_kp_id:  required
*  pns_orang_id:  required
*  golongan_id:  required
 */

func UsulGolonganUpdate(c echo.Context) error {
	db := db.Manager()
	// Paramter utama, wajid ada
	usulan_id := c.FormValue("usulan_id")
	tipe := c.FormValue("tipe")
	IdRiwayatUpdate := c.FormValue("id_riwayat")

	if tipe == "U" {
		if c.FormValue("id_riwayat") == "" {
			return c.JSON(http.StatusOK, map[string]string{
				"error":   "true",
				"message": "id_riwayat harus diisi",
			})
		}
	} else {
		tipe = "I"
	}


	layoutISO := "2006-01-02"
	mkGolonganTahun, err := strconv.Atoi(c.FormValue("mk_golongan_tahun"))
	mkGolonganBulan, err := strconv.Atoi(c.FormValue("mk_golongan_bulan"))
	sk_tanggal, _ := time.Parse(layoutISO, c.FormValue("sk_tanggal"))
	tanggal_letter_bkn, _ := time.Parse(layoutISO, c.FormValue("tanggal_letter_bkn"))
	tmt_golongan, _ := time.Parse(layoutISO, c.FormValue("tmt_golongan"))
	if err != nil {
		response := map[string]interface{}{
			"error":   true,
			"message": err,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	newGolonganUsul := model.Golongan{
		Mk_golongan_tahun:  mkGolonganTahun,
		Mk_golongan_bulan:  mkGolonganBulan,
		Sk_nomor:           c.FormValue("sk_nomor"),
		Sk_tanggal:         sk_tanggal,
		Tanggal_letter_bkn: tanggal_letter_bkn,
		Tmt_golongan:       tmt_golongan,
		Nomor_letter_bkn:   c.FormValue("nomor_letter_bkn"),
		Jenis_kp_id:        c.FormValue("jenis_kp_id"),
		Golongan_id:        c.FormValue("golongan_id"),
		Pns_orang_id:       c.FormValue("pns_orang_id"),
		Tipe:               tipe,
		IdRiwayatUpdate:    IdRiwayatUpdate,
	}
	
	configuration := config.GetConfig()
	var ResultPns map[string]interface{}
	var ResultOrang map[string]interface{}
	url_pns := configuration.PROFILE_PNS_API + "/pns?id="+ c.FormValue("pns_orang_id")
	url_orang := configuration.PROFILE_PNS_API + "/orang?id="+ c.FormValue("pns_orang_id")
	ResultPns = helper.GetCurl(url_pns)
	ResultOrang = helper.GetCurl(url_orang)

	var NIP string
	var Nama string
	if(len(ResultPns["Value"].([]interface{})) > 0){
		NIP = ResultPns["Value"].([]interface{})[0].(map[string]interface{})["nip_baru"].(string)
	} else {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": "Data PNS Tidak Ditemukan",
		})
	}
	if(len(ResultOrang["Value"].([]interface{})) > 0){
		Nama = ResultOrang["Value"].([]interface{})[0].(map[string]interface{})["nama"].(string)
	}

	newUsulanMonitoring := model.Usulan{
		TipeUsulan: 		tipe,
		Nip:                NIP,
		Nama:               Nama,
		// DokumenUsulanLama:  riwayat_lama,
		IdRiwayatUpdate: IdRiwayatUpdate,
		StatusUsulan: 	1,
		TahapanUsulan: 	"Input Berkas",
		StatusUsulanNama: "Belum Lengkap",
	}

	JsonDokUsul, _ := json.Marshal(newGolonganUsul)
	if dbc := db.Debug().Model(&newUsulanMonitoring).Where("id = ?", usulan_id).Update("usulan_data", JsonDokUsul).Updates(newUsulanMonitoring); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	if dbc := db.Debug().Model(&newGolonganUsul).Where("id = ?", usulan_id).Updates(newGolonganUsul); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"id":      usulan_id,
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Golongan Berhasil Diupdate",
	})

}

func UsulDeleteGolongan(c echo.Context) error {
	db := db.Manager()

	IdRiwayatUpdate := c.FormValue("id_riwayat")
	PnsOrangId := c.FormValue("pns_orang_id")
	tipe := "D"

	// var counting int

	// db.Table("rw_golongan").Where("Id = ?", IdRiwayatUpdate).Count(&counting)
	// if counting == 0 {
	// 		// error handling...
	// 	response := map[string]interface{}{
	// 			"status": false,
	// 			"data":   "id_riwayat Not Found",
	// 	}
	// 	return c.JSON(http.StatusBadRequest, response)
	// }

	// var result model.Golongan

	// db.Debug().Table("rw_golongan").Where("Id = ?", IdRiwayatUpdate).Scan(&result)

	// //handler comparison PnsOrangId  with session
	// if result.Pns_orang_id != PnsOrangId {
	// 	// error handling...
	// 	response := map[string]interface{}{
	// 		"status": false,
	// 		"data":   "Invalid Data PNS IDData pns_orang_id tidak cocok",
	// 	}
	// 	return c.JSON(http.StatusBadRequest, response)
	// }

	jenisLayananNama := "Golongan"
	jenisLayanan := "8"
	subLayanan := "39"
	detailLayanan := "121"
	instansiId := "1"
	statusUsulan := 0
	tglUsulan := time.Now()
	namaTableUsul := "rw_golongan_usul"

	layoutISO := "2006-01-02"

	mk_golongan_tahun := 0
	mk_golongan_bulan := 0
	sk_nomor := ""
	sk_tanggal, _ := time.Parse(layoutISO, "2000-01-01")
	tanggal_letter_bkn, _ := time.Parse(layoutISO, "2000-01-01")
	tmt_golongan, _ := time.Parse(layoutISO, "2000-01-01")
	nomor_letter_bkn := ""
	jenis_kp_id := ""
	golongan_id := "0"

	u1 := uuid.Must(uuid.NewV4())
	ustrng := u1.String()

	// json_dok_usulan := map[string]interface{}{}
	// json_dok_usulan["data"] = result
	newGolonganUsul := model.Golongan{
		Id:                 ustrng,
		Mk_golongan_tahun:  mk_golongan_tahun,
		Mk_golongan_bulan:  mk_golongan_bulan,
		Sk_nomor:           sk_nomor,
		Sk_tanggal:         sk_tanggal,
		Tanggal_letter_bkn: tanggal_letter_bkn,
		Tmt_golongan:       tmt_golongan,
		Nomor_letter_bkn:   nomor_letter_bkn,
		Jenis_kp_id:        jenis_kp_id,
		Pns_orang_id:       PnsOrangId,
		Golongan_id:        golongan_id,
		Tipe:               tipe,
		IdRiwayatUpdate:    IdRiwayatUpdate,
	}

	newUsulanMonitoring := model.Usulan{
		IDUsulan:         u1,
		JenisLayanan:     jenisLayanan,
		JenisLayananNama: jenisLayananNama,
		SubLayanan:       subLayanan,
		DetailLayanan:    detailLayanan,
		PnsID:            PnsOrangId,
		InstansiID:       instansiId,
		StatusUsulan:     statusUsulan,
		TglUsulan:        tglUsulan,
		NamaTableUsul:    namaTableUsul,
		TipeUsulan:       tipe,
		IdRiwayatUpdate:  IdRiwayatUpdate,
		// DokumenUsulan:    json_dok_usulan,
		// DokumenUsulanLama: json_dok_usulan,
	}
	// fmt.Println(json_dokumen);

	// Insert Di tabel usulan
	if dbc := db.Debug().Create(&newUsulanMonitoring); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Insert Di tabel rw_golongan usul

	if dbc := db.Debug().Create(&newGolonganUsul); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Penghapusan Riwayat Golongan Berhasil Disimpan",
		"id":      ustrng,
	})

}

/*
* endpoint: /peremajaan/golongan/update-data
* parameter:
*  usulan_id:  required
*  mk_golongan_tahun:numeric  / required
*  mk_golongan_bulan:numeric / required
*  sk_nomor:  required
*  sk_tanggal:yyyy-mm-dd  required
*  tanggal_letter_bkn:yyyy-mm-dd  required
*  tmt_golongan:yyyy-mm-dd  required
*  nomor_letter_bkn:  required
*  jenis_kp_id:  required
*  pns_orang_id:  required
*  golongan_id:  required
 */

func UpdatePejabatApproval(c echo.Context) error {
	db := db.Manager()
	nama_layanan := c.FormValue("nama_layanan")
	// Paramter utama, wajid ada
	usulan_id := c.FormValue("usulan_id")
	pns_orang_approval_id := c.FormValue("pns_orang_approval_id")
	pns_orang_approval_kedua_id := c.FormValue("pns_orang_approval_kedua_id")

	if dbc := db.Debug().Table("rw_"+nama_layanan+"_usul").Where("id = ?", usulan_id).Updates(map[string]interface{}{"pns_orang_approval_id": pns_orang_approval_id, "pns_orang_approval_kedua_id": pns_orang_approval_kedua_id}); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"id":      usulan_id,
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Golongan Berhasil Diupdate",
	})

}

func UpdateSpesimen(c echo.Context) error {
	db := db.Manager()

	usulan_id := c.FormValue("usulan_id")
	pns_orang_approval_id := c.FormValue("spesimen_id")

	if dbc := db.Debug().Table("usulan").Where("id = ?", usulan_id).Updates(map[string]interface{}{"pns_orang_approval_id": pns_orang_approval_id}); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"id":      usulan_id,
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Spesimen berhasil diperbarui",
	})

}

func VerifikasiBerkas(c echo.Context) error {
	db := db.Manager()

	usulan_id := c.FormValue("usulan_id")

	if dbc := db.Debug().Table("usulan").Where("id = ?", usulan_id).Updates(map[string]interface{}{"status_usulan": 2}); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"id":      usulan_id,
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Berkas Berhasil Diverifikasi",
	})

}

/*
* endpoint: /peremajaan/golongan/update-data
* parameter:
*  pns_orang_id:  required
*  golongan_id:  required
 */
func BlankUsulGolonganSave(c echo.Context) error {
	db := db.Manager()
	layoutISO := "2006-01-02"
	mk_golongan_tahun := 0
	mk_golongan_bulan := 0
	sk_nomor := ""
	sk_tanggal, _ := time.Parse(layoutISO, "2000-01-01")
	tanggal_letter_bkn, _ := time.Parse(layoutISO, "2000-01-01")
	tmt_golongan, _ := time.Parse(layoutISO, "2000-01-01")
	nomor_letter_bkn := ""
	jenis_kp_id := ""
	golongan_id := "0"
	tipe := "I"
	pns_orang_id := c.FormValue("pns_orang_id")
	// Golongan := model.Golongan{}

	u1 := uuid.Must(uuid.NewV4())
	ustrng := u1.String()

	newGolonganUsul := model.Golongan{
		Id:                 ustrng,
		Mk_golongan_tahun:  mk_golongan_tahun,
		Mk_golongan_bulan:  mk_golongan_bulan,
		Sk_nomor:           sk_nomor,
		Sk_tanggal:         sk_tanggal,
		Tanggal_letter_bkn: tanggal_letter_bkn,
		Tmt_golongan:       tmt_golongan,
		Nomor_letter_bkn:   nomor_letter_bkn,
		Jenis_kp_id:        jenis_kp_id,
		Pns_orang_id:       pns_orang_id,
		Golongan_id:        golongan_id,
		Tipe:               tipe,
	}

	// Tambahan Insert Setelah Insert Data Ke Tabel *_usul, bentuknya sama semua,
	//untuk nilai jenis layanan akan silahkan lihat di file excelnya,

	jenisLayananNama := "Golongan"
	jenisLayanan := "8"
	subLayanan := "39"
	detailLayanan := "121"
	instansiId := "1"
	statusUsulan := 1
	tglUsulan := time.Now()
	// isian namaTableUsul berdasarkan nama tabel yang digunakan
	namaTableUsul := "rw_golongan_usul"
	namaTabelRiwayat := "rw_golongan"
	sumber := c.FormValue("sumber")

	sumberUsulan := map[string]string{
		"mysapk":   "MYSAPK",
		"instansi": "SIASN INSTANSI",
		"dms":      "DMS",
		"simpeg":   "SIMPEG",
	}

	value, found := sumberUsulan[sumber]
	if !found {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": "nilai " + value + " pada field 'sumber' tidak valid",
		})
	}

	DeviceToken := c.FormValue("device_token")
	if sumber == "mysapk" && DeviceToken == "" {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": "device_token Harus Diisi Jika Sumber dari MYSAPK",
		})
	}

	configuration := config.GetConfig()
	var ResultPns map[string]interface{}
	var ResultOrang map[string]interface{}
	url_pns := configuration.PROFILE_PNS_API + "/pns?id="+ c.FormValue("pns_orang_id")
	url_orang := configuration.PROFILE_PNS_API + "/orang?id="+ c.FormValue("pns_orang_id")
	ResultPns = helper.GetCurl(url_pns)
	ResultOrang = helper.GetCurl(url_orang)

	var NIP string
	var Nama string
	
	if(len(ResultPns["Value"].([]interface{})) > 0){
		NIP = ResultPns["Value"].([]interface{})[0].(map[string]interface{})["nip_baru"].(string)
		instansiId = ResultPns["Value"].([]interface{})[0].(map[string]interface{})["instansi_kerja_id"].(string)
	} else {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": "Data PNS Tidak Ditemukan",
		})
	}
	if(len(ResultOrang["Value"].([]interface{})) > 0){
		Nama = ResultOrang["Value"].([]interface{})[0].(map[string]interface{})["nama"].(string)
	}
	BlankUsulan := make(map[string]interface{})

	newUsulanMonitoring := model.Usulan{
		IDUsulan:         u1,
		JenisLayanan:     jenisLayanan,
		UsulanData:       BlankUsulan,
		JenisLayananNama: jenisLayananNama,
		SubLayanan:       subLayanan,
		DetailLayanan:    detailLayanan,
		PnsID:            pns_orang_id,
		InstansiID:       instansiId,
		StatusUsulan:     statusUsulan,
		TglUsulan:        tglUsulan,
		NamaTableUsul:    namaTableUsul,
		NamaTabelRiwayat: namaTabelRiwayat,
		Sumber:           sumberUsulan[sumber],
		Nip:                NIP,
		Nama:               Nama,
		DeviceToken: 	DeviceToken,
	}

	// Insert Di tabel usulan
	if dbc := db.Debug().Create(&newUsulanMonitoring); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Insert Di tabel rw_golongan usul

	if dbc := db.Debug().Create(&newGolonganUsul); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Data Usulan Golongan Berhasil Disimpan",
		"id":      ustrng,
	})

}

/*
* endpoint: /peremajaan/golongan/get-verifikasi-data
* parameter:
*  pns_orang_id:  required
*  golongan_id:  required
 */
func GolonganUsulGetValidasiData(c echo.Context) error {
	db := db.Manager()
	usulan_id := c.Param("usulan_id")
	rwGolonganUsul := model.Golongan{}

	db.Where("id = ?", usulan_id).First(&rwGolonganUsul)

	if dbc := db.Debug().Where("id = ?", usulan_id).First(&rwGolonganUsul); dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	type golongan struct {
		MkGolonganTahun  bool
		MkGolonganBulan  bool
		SkNomor          bool
		SkTanggal        bool
		TanggalLetterBkn bool
		TmtGolongan      bool
		NomorLetterBkn   bool
		JenisKpId        bool
		PnsOrangId       bool
		GolonganId       bool
	}

	//seeding pengguna by raw-survei
	var resultGolongan = golongan{}
	resultGolongan.MkGolonganTahun = false
	resultGolongan.MkGolonganBulan = false
	resultGolongan.SkNomor = false
	resultGolongan.SkTanggal = false
	resultGolongan.TanggalLetterBkn = false
	resultGolongan.TmtGolongan = false
	resultGolongan.NomorLetterBkn = false
	resultGolongan.JenisKpId = false
	resultGolongan.PnsOrangId = false
	resultGolongan.GolonganId = false

	if strconv.Itoa(rwGolonganUsul.Mk_golongan_tahun) != "" {
		resultGolongan.MkGolonganTahun = true
	}

	if strconv.Itoa(rwGolonganUsul.Mk_golongan_bulan) != "" {
		resultGolongan.MkGolonganBulan = true
	}

	if rwGolonganUsul.Sk_nomor != "" {
		resultGolongan.SkNomor = true
	}

	if rwGolonganUsul.Sk_tanggal.String() != "" {
		resultGolongan.SkTanggal = true
	}

	if rwGolonganUsul.Tanggal_letter_bkn.String() != "" {
		resultGolongan.TanggalLetterBkn = true
	}

	if rwGolonganUsul.Tmt_golongan.String() != "" {
		resultGolongan.TmtGolongan = true
	}

	if rwGolonganUsul.Nomor_letter_bkn != "" {
		resultGolongan.NomorLetterBkn = true
	}

	if rwGolonganUsul.Jenis_kp_id != "" {
		resultGolongan.JenisKpId = true
	}
	if rwGolonganUsul.Pns_orang_id != "" {
		resultGolongan.PnsOrangId = true
	}
	if rwGolonganUsul.Golongan_id != "" {
		resultGolongan.GolonganId = true
	}

	dataAll := map[string]interface{}{}
	dataAll["detail"] = rwGolonganUsul
	dataAll["validasi"] = resultGolongan

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"data":    dataAll,
		"message": "Get RwGolongan Succes",
	})

}
