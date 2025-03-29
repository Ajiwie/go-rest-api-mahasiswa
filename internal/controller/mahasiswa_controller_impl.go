package controller

import (
	"net/http"
	"strconv"

	"github.com/Ajiwie/go-rest-api-mahasiswa/helper"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/service"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/web"
	"github.com/julienschmidt/httprouter"
)

type MahasiswaControllerImpl struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService) *MahasiswaControllerImpl {
	return &MahasiswaControllerImpl{MahasiswaService: mahasiswaService}
}

func (c *MahasiswaControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	mahasiswaCreateRequest := web.MahasiswaCreateRequest{}
	helper.ReadFromRequestBody(r, &mahasiswaCreateRequest)

	mahasiswaResponse := c.MahasiswaService.Create(r.Context(), mahasiswaCreateRequest)
	WebResponse := web.MahasiswaWebResponse{
		Code:   200,
		Status: "OK",
		Data:   mahasiswaResponse,
	}

	helper.WriterToResponseBody(w, WebResponse)
}

func (c *MahasiswaControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	mahasiswaUpdateRequest := web.MahasiswaUpdateRequest{}
	helper.ReadFromRequestBody(r, &mahasiswaUpdateRequest)

	mahasiswaId := params.ByName("mahasiswaId")
	id, err := strconv.Atoi(mahasiswaId)
	helper.PanicIfErr(err)

	mahasiswaUpdateRequest.Id = id

	mahasiswaResponse := c.MahasiswaService.Update(r.Context(), mahasiswaUpdateRequest)
	WebResponse := web.MahasiswaWebResponse{
		Code:   200,
		Status: "OK",
		Data:   mahasiswaResponse,
	}

	helper.WriterToResponseBody(w, WebResponse)

}
