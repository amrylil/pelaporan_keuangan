package dtos

type ResKategori struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi,omitempty"`
}
