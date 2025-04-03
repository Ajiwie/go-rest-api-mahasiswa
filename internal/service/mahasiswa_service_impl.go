package service

import (
	"context"
	"database/sql"

	"github.com/Ajiwie/go-rest-api-mahasiswa/exception"
	"github.com/Ajiwie/go-rest-api-mahasiswa/helper"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/model"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/repository"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/web"
	"github.com/go-playground/validator/v10"
)

type MahasiswaServiceImpl struct {
	Repo     repository.MahasiswaRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewMahasiswaService(repo repository.MahasiswaRepository, db *sql.DB, validate *validator.Validate) *MahasiswaServiceImpl {
	return &MahasiswaServiceImpl{Repo: repo, DB: db, Validate: validate}
}

func (s *MahasiswaServiceImpl) Create(ctx context.Context, request web.MahasiswaCreateRequest) web.MahasiswaResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := model.Mahasiswa{
		Nama:    request.Nama,
		NIM:     request.Nim,
		Jurusan: request.Jurusan,
	}

	mahasiswa = s.Repo.Create(ctx, tx, mahasiswa)
	return helper.ToMahasiswaResponse(mahasiswa)
}

func (s *MahasiswaServiceImpl) Update(ctx context.Context, request web.MahasiswaUpdateRequest) web.MahasiswaResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := s.Repo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	mahasiswa.Nama = request.Nama
	mahasiswa = s.Repo.Update(ctx, tx, mahasiswa)
	return helper.ToMahasiswaResponse(mahasiswa)
}

func (s *MahasiswaServiceImpl) Delete(ctx context.Context, mahasiswaId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := s.Repo.FindById(ctx, tx, mahasiswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.Repo.Delete(ctx, tx, mahasiswa)
}

func (s *MahasiswaServiceImpl) FindById(ctx context.Context, mahasiswaId int) web.MahasiswaResponse {

	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := s.Repo.FindById(ctx, tx, mahasiswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMahasiswaResponse(mahasiswa)
}

func (s *MahasiswaServiceImpl) FindAll(ctx context.Context) []web.MahasiswaResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := s.Repo.FindAll(ctx, tx)
	return helper.ToMahasiswaResponses(mahasiswa)
}
