package web

type MahasiswaCreateRequest struct {
	Nama    string `validate:"required,min=1,max=100" json:"nama"`
	Nim     string `validate:"required,min=1,max=100" json:"nim"`
	Jurusan string `validate:"required,min=1,max=100" json:"jurusan"`
}
