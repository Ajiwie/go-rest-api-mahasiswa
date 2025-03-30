package config

import (
	"github.com/Ajiwie/go-rest-api-mahasiswa/exception"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(mahasiswaController controller.MahasiswaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/mahasiswa", mahasiswaController.FindAll)
	router.GET("/api/mahasiswa/:mahasiswaId", mahasiswaController.FindById)
	router.POST("/api/mahasiswa", mahasiswaController.Create)
	router.PUT("/api/mahasiswa/:mahasiswaId", mahasiswaController.Update)
	router.DELETE("/api/mahasiswa/:mahasiswaId", mahasiswaController.Delete)

	router.PanicHandler = exception.ErrHandler
	return router
}
