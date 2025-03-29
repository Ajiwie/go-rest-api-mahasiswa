package helper

import (
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/model"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/web"
)

func ToMahasiswaResponse(mahasiswa model.Mahasiswa) web.MahasiswaResponse {
	return web.MahasiswaResponse{
		Id:      mahasiswa.Id,
		Nama:    mahasiswa.Nama,
		NIM:     mahasiswa.NIM,
		Jurusan: mahasiswa.Jurusan,
	}
}

func ToMahasiswaResponses(mahasiswa []model.Mahasiswa) []web.MahasiswaResponse {
	var mahasiswaResponses []web.MahasiswaResponse
	for _, mahasmahasiswa := range mahasiswa {
		mahasiswaResponses = append(mahasiswaResponses, web.MahasiswaResponse(mahasmahasiswa))
	}
	return mahasiswaResponses
}
