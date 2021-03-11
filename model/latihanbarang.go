package model

import "time"

type Personal1 struct {
	Id           string    `jason:"id" 				valid:"required"`
	Barang       string    `jason:"barang" 			valid:"required"`
	Berat        int       `jason:"berat" 			valid:"required"`
	TanggalMasuk time.Time `jason:"tanggal_masuk" 	valid:"required"`
}

func (Personal1) TableName() string {
	return "latihanbarang"
}
