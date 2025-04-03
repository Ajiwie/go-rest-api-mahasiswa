package main

import (
	"net/http"

	"github.com/Ajiwie/go-rest-api-mahasiswa/config"
	"github.com/Ajiwie/go-rest-api-mahasiswa/helper"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/controller"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/middleware"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/repository"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/service"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.Database()
	validate := validator.New()
	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepository, db, validate)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaService)
	router := config.NewRouter(mahasiswaController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
