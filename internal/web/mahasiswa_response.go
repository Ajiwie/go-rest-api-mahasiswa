package web

import "time"

type MahasiswaResponse struct {
	Id        int       `json:"id"`
	Nama      string    `json:"nama"`
	NIM       string    `json:"nim"`
	Jurusan   string    `json:"jurusan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
