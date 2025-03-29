package web

type MahasiswaUpdateRequest struct {
	Id   int    `validate:"required"`
	Nama string `validate:"required,max=200,min=1" json:"nama"`
}
