package models

type User struct {
	Nama    string `json:"nama" validate:"required,min=3,max=40"`
	Email   string `json:"email" validate:"required,email,min=6,max=32"`
	Tanggal string `json:"tanggal" validate:"required"`
	Kota    string `json:"kota" validate:"required,min=5,max=25"`
	Negara  string `json:"negara" validate:"required,min=5,max=25"`
}
