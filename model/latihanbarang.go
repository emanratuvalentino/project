package model

import "time"

type Personal1 struct {
	Id           string    `json:"id"    valid:"required"`
	Barang       string    `json:"barang"  valid:"required"`
	Berat        int       `json:"berat"  valid:"required"`
	TanggalMasuk time.Time `json:"tanggal_masuk"  valid:"required"`
}

func (Personal1) TableName() string {
	return "latihan_barang"
}
