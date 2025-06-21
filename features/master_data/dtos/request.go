package dtos

type InputMaster_data struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type InputJenisPembayaran struct {
	Name      string  `json:"name" form:"name" validate:"required"`
	Deskripsi *string `json:"deskripsi" form:"deskripsi"`
}

type InputTipeTransaksi struct {
	Name      string  `json:"name" form:"name" validate:"required"`
	Deskripsi *string `json:"deskripsi" form:"deskripsi"`
}
type InputStatusTransaksi struct {
	Name      string  `json:"name" form:"name" validate:"required"`
	Deskripsi *string `json:"deskripsi" form:"deskripsi"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
