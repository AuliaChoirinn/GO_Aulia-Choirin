package kategori

import (

	"Code_Competence/config"
	"Code_Competence/models"

)

type Kategori struct {
	Nama      string
	Deskripsi string
}

func NewKategori(nama, deskripsi string) *Kategori {
	return &Kategori{
		Nama:      nama,
		Deskripsi: deskripsi,
	}
}

func (k *Kategori) GetNama() string {
	return k.Nama
}

func (k *Kategori) GetDeskripsi() string {
	return k.Deskripsi
}
