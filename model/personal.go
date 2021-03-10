package model

import (
	"time"
)


type Personal struct {
	Id              string       `json:"id"    valid:"required"`
	Nama            string       `json:"nama"  valid:"required"`
	Usia            int          `json:"nama"  valid:"required"`
	TanggalLahir    time.Time    `json:"nama"  valid:"required"`
}


func (Personal) TableName() string {
	return "latihan_sapk"
}