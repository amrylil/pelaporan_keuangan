package dtos

type InputKategori struct {
	Name      string  `json:"name" form:"name" validate:"required,min=3"`
	Deskripsi *string `json:"deskripsi" form:"deskripsi"`
}

// DTO untuk data pagination (tidak berubah)
type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
