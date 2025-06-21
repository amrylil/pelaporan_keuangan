package dtos

type ResKategori struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi,omitempty"`
}
