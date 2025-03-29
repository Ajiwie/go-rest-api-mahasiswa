package model

import "time"

type Mahasiswa struct {
	Id        int
	Nama      string
	NIM       string
	Jurusan   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
