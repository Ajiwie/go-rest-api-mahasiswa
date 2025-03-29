package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ajiwie/go-rest-api-mahasiswa/helper"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/model"
)

type MahasiswaRepositoryImpl struct {
}

func NewMahasiswaRepository() *MahasiswaRepositoryImpl {
	return &MahasiswaRepositoryImpl{}
}

func (r *MahasiswaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa {
	sql := "INSERT INTO Mahasiswa (nama, nim, jurusan) VALUES (?,?,?)"
	result, err := tx.ExecContext(ctx, sql, mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Jurusan)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	mahasiswa.Id = int(id)
	return mahasiswa
}

func (r *MahasiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa {
	sql := "UPDATE mahasiswa set nama = ? where id = ?)"
	_, err := tx.ExecContext(ctx, sql, mahasiswa.Nama, mahasiswa.Id)
	helper.PanicIfErr(err)
	return mahasiswa
}

func (r *MahasiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) {
	sql := "DELETE from mahasiswa where id= ?"
	_, err := tx.ExecContext(ctx, sql, mahasiswa.Id)
	helper.PanicIfErr(err)
}

func (r *MahasiswaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, mahasiwaId int) (model.Mahasiswa, error) {
	sql := "SELECT id, nama, nim, jurusan FROM mahasiswa where id=?"
	rows, err := tx.QueryContext(ctx, sql, mahasiwaId)
	helper.PanicIfErr(err)
	defer rows.Close()

	mahasiswa := model.Mahasiswa{}
	if rows.Next() {
		err := rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.NIM, &mahasiswa.Jurusan)
		helper.PanicIfErr(err)
		return mahasiswa, nil
	} else {
		return mahasiswa, errors.New("id mahasiswa tidak ditemukan")
	}

}

func (r *MahasiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Mahasiswa {
	sql := "SELECT * FROM mahasiswa"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var mahasiswa []model.Mahasiswa
	for rows.Next() {
		mhsw := model.Mahasiswa{}
		err := rows.Scan(&mhsw.Id, &mhsw.Nama, &mhsw.NIM, &mhsw.Jurusan)
		helper.PanicIfErr(err)
		mahasiswa = append(mahasiswa, mhsw)
	}
	return mahasiswa
}
