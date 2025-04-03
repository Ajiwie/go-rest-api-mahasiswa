package web

type MahasiswaCreateRequest struct {
	Nama    string `validate:"required,max=200,min=1" json:"nama"`
	Nim     string `validate:"required,max=200,min=1" json:"nim"`
	Jurusan string `validate:"required,max=200,min=1" json:"jurusan"`
}
