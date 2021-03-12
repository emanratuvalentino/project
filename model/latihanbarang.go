package model

import (
	"time"
)

type Barang struct {
	Id           string    `json:"id"    valid:"required"`
	Barang       string    `json:"barang"  valid:"required"`
	Berat_kg     int       `json:"berat_kg"  valid:"required"`
	TanggalMasuk time.Time `json:"tanggal_masuk"  valid:"required"`
}

func (Barang) TableName() string {
	return "latihan_barang"
}
