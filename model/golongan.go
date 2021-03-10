package model

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func (Golongan) TableName() string {
	return "rw_golongan_usul"
}

/*
Aturan Pembuatan Model :
1. Diambil dari seluruh field yang ada didalam tabel, tidak perlu ambil field ncistime,
2. Cantumkan valid:"required", JIKA nama kolom tersebut ada didalam request field (cek spreadsheet)
3. Lihat script pembuatan tabel di dalam peremajaan-data.sql, jika ada kolom yang didefiniskan NOT NULL,
   maka cantumkan default:0 didalam blok gorm, contoh : gorm:"column:jumlah_kredit_utama; default:0"
*/
type Golongan struct {
	Id                     string    `gorm:"column:id" json:"id"  valid:"required"`
	Jenis_kp_id            string    `gorm:"column:jenis_kp_id" json:"jenis_kp_id"  valid:"required" sql:"DEFAULT:NULL"`
	Golongan_id            string    `gorm:"column:golongan_id" json:"golongan_id" sql:"DEFAULT:'1'"`
	Pns_orang_id           string    `gorm:"column:pns_orang_id" json:"pns_orang_id"  valid:"required" sql:"DEFAULT:NULL"`
	Sk_nomor               string    `gorm:"column:sk_nomor" json:"sk_nomor"  valid:"required" sql:"DEFAULT:NULL"`
	Sk_tanggal             time.Time `gorm:"column:sk_tanggal" json:"sk_tanggal"  valid:"required" sql:"DEFAULT:NULL"`
	Nomor_letter_bkn       string    `gorm:"column:nomor_letter_bkn" json:"nomor_letter_bkn"  valid:"required" sql:"DEFAULT:NULL"`
	Tanggal_letter_bkn     time.Time `gorm:"column:tanggal_letter_bkn" json:"tanggal_letter_bkn"  valid:"required" sql:"DEFAULT:NULL"`
	Tmt_golongan           time.Time `gorm:"column:tmt_golongan" json:"tmt_golongan"  valid:"required" sql:"DEFAULT:NULL"`
	Jumlah_kredit_utama    float32   `gorm:"column:jumlah_kredit_utama; default:0" json:"jumlah_kredit_utama"`
	Jumlah_kredit_tambahan float32   `gorm:"column:jumlah_kredit_tambahan; default:0" json:"jumlah_kredit_tambahan"`

	// Asal_id string `gorm:"column:asal_id" json:"asal_id" sql:"DEFAULT:NULL"`
	// Asal_nama string `gorm:"column:asal_nama" json:"asal_nama" sql:"DEFAULT:NULL"`
	Status_selesai    string `gorm:"column:status_selesai" json:"status_selesai" sql:"DEFAULT:NULL"`
	Mk_golongan_tahun int    `gorm:"column:mk_golongan_tahun" json:"mk_golongan_tahun"  valid:"required" sql:"DEFAULT:NULL"`
	Mk_golongan_bulan int    `gorm:"column:mk_golongan_bulan" json:"mk_golongan_bulan"  valid:"required" sql:"DEFAULT:NULL"`

	Tanggal_usul                string         `gorm:"column:tanggal_usul" json:"tanggal_usul" sql:"DEFAULT:current_timestamp"`
	Status_usul                 string         `gorm:"column:status_usul" json:"status_usul" sql:"DEFAULT:NULL"`
	Keterangan_status           string         `gorm:"column:keterangan_status" json:"keterangan_status" sql:"DEFAULT:NULL"`
	Pns_orang_approval_id       string         `gorm:"column:pns_orang_approval_id" json:"pns_orang_approval_id" sql:"DEFAULT:NULL"`
	Tanggal_approval            string         `gorm:"column:tanggal_approval" json:"tanggal_approval" sql:"DEFAULT:NULL"`
	Path                        string         `gorm:"column:path" json:"path" sql:"DEFAULT:NULL"`
	Pns_orang_approval_kedua_id string         `gorm:"column:pns_orang_approval_kedua_id" json:"pns_orang_approval_kedua_id" sql:"DEFAULT:NULL"`
	Tanggal_approval_kedua      string         `gorm:"column:tanggal_approval_kedua" json:"tanggal_approval_kedua" sql:"DEFAULT:NULL"`
	Tipe                        string         `gorm:"column:tipe" json:"tipe" sql:"DEFAULT:NULL"`
	Id_riwayat_update           string         `gorm:"column:id_riwayat_update" json:"id_riwayat_update" sql:"DEFAULT:NULL"`
	Is_open                     string         `gorm:"column:is_open" json:"is_open" sql:"DEFAULT:NULL"`
	Golongan_nama               postgres.Jsonb `gorm:"column:golongan_nama" json:"golongan_nama" sql:"DEFAULT:NULL"`
	IdRiwayatUpdate             string         `gorm:"column:id_riwayat_update" json:"id_riwayat_update" sql:"DEFAULT:NULL"`
	Jenis_kp_nama               string         `gorm:"column:jenis_kp_nama" json:"jenis_kp_nama" sql:"DEFAULT:NULL"`
	JumlahKreditUtama           float32        `json:"jumlah_kredit_utama" gorm:"column:jumlah_kredit_utama;"`
	JumlahKreditUTambahan       float32        `json:"jumlah_kredit_tambahan" gorm:"column:jumlah_kredit_tambahan;"`
}
