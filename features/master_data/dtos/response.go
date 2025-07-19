package dtos

type ResMaster_data struct {
	Name string `json:"name"`
}

type ResJenisPembayaran struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi,omitempty"`
}
type ResTipeTransaksi struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi,omitempty"`
}
type ResStatusTransaksi struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi,omitempty"`
}
