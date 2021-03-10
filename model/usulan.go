package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type Usulan struct {
	// gorm.Model
	IDUsulan               uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();column:id"`
	JenisLayanan           string    `json:"jenis_layanan"`
	SubLayanan             string    `json:"sub_layanan"`
	DetailLayanan          string    `json:"detail_layanan"`
	DetailLayananNama      string    `json:"detail_layanan_nama"`
	PnsID                  string    `json:"pns_id"`
	UsulanData             JSONB     `json:"usulan_data" sql:"type:jsonb"`
	StatusUsulan           int       `json:"status_usulan"`
	StatusUsulanNama       string    `json:"status_usulan_nama"`
	DokumenUsulan          JSONB     `json:"dokumen_usulan" sql:"type:jsonb"`
	DokumenUsulanLama      JSONB     `json:"dokumen_usulan_lama" sql:"type:jsonb"`
	TglUsulan              time.Time `json:"tgl_usulan" gorm:"type:time"`
	TglPengirimanKeLayanan time.Time `json:"tgl_pengiriman_ke_layanan" gorm:"type:time;column:tgl_pengiriman_kelayanan" `
	TglUpdateDariLayanan   time.Time `json:"tgl_update_dari_layanan" gorm:"type:time;column:tgl_update_layanan" `
	InstansiID             string    `json:"instansi_id"`
	InstansiNama           string    `json:"instansi_nama"`
	Keterangan             string    `json:"keterangan"`
	StatusAktif            int       `json:"status_aktif"`
	JenisLayananNama       string    `json:"jenis_layanan_nama"`
	TglSuratUsulan         time.Time `json:"tgl_surat_usulan" gorm:"type:time"`
	TglSuratKeluar         time.Time `json:"tgl_surat_keluar" gorm:"type:time"`
	NoSuratUsulan          string    `json:"no_surat_usulan"`
	NoSuratKeluar          string    `json:"no_surat_keluar"`
	Nama                   string    `json:"nama"`
	Nip                    string    `json:"nip"`
	NamaTableUsul          string    `json:"nama_table_usul"`
	TipeUsulan             string    `json:"tipe_usulan"`
	TahapanUsulan          string    `json:"tahapan_usulan"`
	IdRiwayatUpdate        string    `json:"id_riwayat_update"`
	NamaTabelRiwayat       string    `json:"nama_tabel_riwayat"`
	PathSuratUsulan        string    `json:"path_surat_usulan"`
	Sumber                 string    `json:"sumber"`
	DeviceToken            string    `json:"device_token"`
	ProvinsiNama           string    `json:"provinsi_nama"`
	TipePegawai			   string    `json:"tipe_pegawai"`
}

type UsulanPeremajaanFull struct {
	// gorm.Model
	IDUsulan               uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();column:id"`
	JenisLayanan           *int    `json:"jenis_layanan"`
	SubLayanan             *int    `json:"sub_layanan"`
	DetailLayanan          *int    `json:"detail_layanan"`
	DetailLayananNama      *string    `json:"detail_layanan_nama"`
	PnsID                  *string    `json:"pns_id"`
	UsulanData             *string    `json:"usulan_data" sql:"type:jsonb"`
	StatusUsulan           *int       `json:"status_usulan"`
	StatusUsulanNama       *string    `json:"status_usulan_nama"`
	DokumenUsulan          *string    `json:"dokumen_usulan" sql:"type:jsonb"`
	DokumenUsulanLama      *string    `json:"dokumen_usulan_lama" sql:"type:jsonb"`
	TglUsulan              *string 	 `json:"tgl_usulan" gorm:"type:time"`
	TglPengirimanKeLayanan *string 	 `json:"tgl_pengiriman_ke_layanan" gorm:"type:time;column:tgl_pengiriman_kelayanan" `
	TglUpdateDariLayanan   *string 	 `json:"tgl_update_dari_layanan" gorm:"type:time;column:tgl_update_layanan" `
	InstansiID             *string    `json:"instansi_id"`
	InstansiNama           *string    `json:"instansi_nama"`
	Keterangan             *string    `json:"keterangan"`
	StatusAktif            *int       `json:"status_aktif"`
	JenisLayananNama       *string    `json:"jenis_layanan_nama"`
	TglSuratUsulan         *string 	 `json:"tgl_surat_usulan" gorm:"type:time"`
	TglSuratKeluar         *string 	 `json:"tgl_surat_keluar" gorm:"type:time"`
	NoSuratUsulan          *string    `json:"no_surat_usulan"`
	NoSuratKeluar          *string    `json:"no_surat_keluar"`
	Nama                   *string    `json:"nama"`
	Nip                    *string    `json:"nip"`
	NamaTableUsul          *string    `json:"nama_table_usul"`
	TipeUsulan             *string    `json:"tipe_usulan"`
	TahapanUsulan          *string    `json:"tahapan_usulan"`
	IdRiwayatUpdate        *string    `json:"id_riwayat_update"`
	NamaTabelRiwayat       *string    `json:"nama_tabel_riwayat"`
	PathSuratUsulan        *string    `json:"path_surat_usulan"`
	Sumber                 *string    `json:"sumber"`
	DeviceToken            *string    `json:"device_token"`
	ProvinsiNama           *string    `json:"provinsi_nama"`
	TipePegawai			   *string    `json:"tipe_pegawai"`
}


type UsulanMonitoring struct {
	IDUsulan               uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();column:id"`
	StatusUsulan           int        `json:"status_usulan"`
	StatusUsulanNama       string     `json:"status_usulan_nama"`
	TglUsulan              *time.Time `json:"tgl_usulan" gorm:"type:time"`
	Keterangan             string     `json:"keterangan"`
	NoSuratUsulan          string     `json:"no_surat_usulan"`
	ProvinsiNama           string     `json:"provinsi_nama"`
	InstansiNama           string     `json:"instansi_nama"`
	UsulanData             JSONB      `json:"usulan_data" sql:"type:jsonb"`
	DokumenUsulan          JSONB      `json:"dokumen_usulan" sql:"type:jsonb"`
	JenisLayananNama       string     `json:"jenis_layanan_nama"`
	Nama                   string     `json:"nama"`
	Sumber                 string     `json:"sumber"`
	Nip                    string     `json:"nip"`
}

type UsulanDetail struct {
	IDUsulan               uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();column:id"`
	UsulanData             JSONB     `json:"usulan_data" sql:"type:jsonb"`
	DokumenUsulan          JSONB      `json:"dokumen_usulan" sql:"type:jsonb"`
	TglUsulan              *time.Time `json:"tgl_usulan" gorm:"type:time"`
	JenisLayananNama       string     `json:"jenis_layanan_nama"`
	Nama                   string     `json:"nama"`
	Nip                    string     `json:"nip"`
}

type UsulanUpload struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	JenisLayanan     int       `json:"jenis_layanan"`
	SubLayanan       int       `json:"sub_layanan"`
	DetailLayanan    int       `json:"detail_layanan"`
	NoSuratUsulan    string    `json:"no_surat_usulan"`
	PnsID            string    `json:"pns_id"`
	InstansiID       string    `json:"instansi_id"`
	JenisLayananNama string    `json:"jenis_layanan_nama"`
	Nip              string    `json:"nip"`
	Nama             string    `json:"nama"`
	InstansiNama     string    `json:"instansi_nama"`
	ProvinsiNama     string    `json:"provinsi_nama"`
	DokumenUsulan    string    `gorm:"type:jsonb" json:"dokumen_usulan" form:"dokumen_usulan"`
	UsulanData       JSONB     `gorm:"type:jsonb" json:"usulan_data"`
	StatusUsulan     int       `json:"status_usulan"`
	NamaTableUsul          string    `json:"nama_table_usul"`
	
}

type FilterMonitUsulan struct {
	NoSuratUsulan    string `json:"no_surat_usulan,omitempty" form:"no_surat_usulan" query:"no_surat_usulan"`
	DetailLayanan    string `json:"detail_layanan,omitempty" query:"detail_layanan"`
	Nip              string `json:"nip,omitempty" query:"nip"`
	Nama             string `json:"nama,omitempty" query:"nama"`
	StatusUsulan     int    `json:"status_usulan,omitempty" query:"status_usulan"`
	TglUsulan        string `json:"tgl_usulan,omitempty" query:"tgl_usulan"`
	TipePegawai      string `json:"tipe_pegawai,omitempty" query:"tipe_pegawai"`
}

type DokumenUsulanPeremajaan struct {
	DokId   string `json:"dok_id"`
	DokNama string `json:"dok_nama"`
	DokUri string `json:"dok_uri"`
}

type UpdateMonitUsulan struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"id" form:"id" query:"id"`
}

func (Usulan) TableName() string {
	return "usulan"
}
func (UsulanUpload) TableName() string {
	return "usulan"
}

func (FilterMonitUsulan) TableName() string {
	return "usulan"
}

func (UsulanMonitoring) TableName() string {
	return "usulan"
}
func (UsulanPeremajaanFull) TableName() string {
	return "usulan"
}
