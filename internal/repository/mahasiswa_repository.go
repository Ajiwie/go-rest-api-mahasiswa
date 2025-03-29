package repository

import (
	"context"
	"database/sql"

	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/model"
)

type MahasiswaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa
	Update(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa
	Delete(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa)
	FindById(ctx context.Context, tx *sql.Tx, mahasiwaId int) (model.Mahasiswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Mahasiswa
}
