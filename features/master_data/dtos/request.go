package dtos

type InputMaster_data struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type InputJenisPembayaran struct {
	Nama string `json:"nama_jenis_pembayaran" form:"nama_jenis_pembayaran" validate:"required"`
}

type InputTipeTransaksi struct {
	Nama string `json:"name" form:"name" validate:"required"`
}
type InputStatusTransaksi struct {
	Nama string `json:"name" form:"name" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
