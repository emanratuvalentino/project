package model

import (
	"time"
)


type Personal struct {
	Id              string       `json:"id"    valid:"required"`
	Nama            string       `json:"nama"  valid:"required"`
	Usia            int          `json:"usia"  valid:"required"`
	TanggalLahir    time.Time    `json:"tanggal_lahir"  valid:"required"`
}


func (Personal) TableName() string {
	return "latihan_sapk"
}