package service

import (
	"context"

	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/web"
)

type MahasiswaService interface {
	Create(ctx context.Context, request web.MahasiswaCreateRequest) web.MahasiswaResponse
	Update(ctx context.Context, request web.MahasiswaUpdateRequest) web.MahasiswaResponse
	Delete(ctx context.Context, mahasiswaId int)
	FindById(ctx context.Context, mahasiswaId int) web.MahasiswaResponse
	FindAll(ctx context.Context) []web.MahasiswaResponse
}
