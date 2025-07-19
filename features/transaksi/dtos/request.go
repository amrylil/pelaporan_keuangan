package dtos

type InputTransaksi struct {
	Tanggal           string  `form:"tanggal"             validate:"required,datetime=2006-01-02"`
	NamaTransaksi     string  `form:"nama_transaksi"      validate:"required,max=255"`
	Jumlah            float64 `form:"jumlah"              validate:"required,gt=0"`
	Keterangan        string  `form:"keterangan"`
	IDTipeTransaksi   uint64  `form:"id_tipe_transaksi"   validate:"required"`
	IDKategori        uint64  `form:"id_kategori"         validate:"required"`
	IDJenisPembayaran uint64  `form:"id_jenis_pembayaran" validate:"required"`
	IDUser            uint64  `validate:"required"`
	IDStatusTransaksi uint64  `form:"id_status_transaksi" validate:"required"`
	BuktiTransaksi    string
}

// UpdateTransaksiRequest - DTO untuk update transaksi
type UpdateTransaksiRequest struct {
	ID                *uint64  `json:"id" validate:"required"`
	Tanggal           *string  `json:"tanggal,omitempty"`
	IDTipeTransaksi   *uint64  `json:"id_tipe_transaksi,omitempty"`
	Jumlah            *float64 `json:"jumlah,omitempty"`
	Keterangan        *string  `json:"keterangan,omitempty"`
	BuktiTransaksi    *string  `json:"bukti_transaksi,omitempty"`
	IDStatusTransaksi *uint64  `json:"id_status_transaksi,omitempty"`
	KomentarManajer   *string  `json:"komentar_manajer,omitempty"`
	IDKategori        *uint64  `json:"id_kategori,omitempty"`
	IDUser            *uint64  `json:"id_user,omitempty"`
	IDJenisPembayaran *uint64  `json:"id_jenis_pembayaran,omitempty"`
}

type TransaksiListRequest struct {
	Page              int      `query:"page" validate:"min=1" example:"1"`
	Limit             int      `query:"limit" validate:"min=1,max=100" example:"10"`
	Search            string   `query:"search" example:"makan"`
	IDTipeTransaksi   *uint64  `query:"id_tipe_transaksi"`
	IDStatusTransaksi *uint64  `query:"id_status_transaksi"`
	IDKategori        *uint64  `query:"id_kategori"`
	IDUser            *uint64  `query:"id_user"`
	IDJenisPembayaran *uint64  `query:"id_jenis_pembayaran"`
	TanggalMulai      *string  `query:"tanggal_mulai" example:"2024-01-01"`
	TanggalAkhir      *string  `query:"tanggal_akhir" example:"2024-12-31"`
	JumlahMin         *float64 `query:"jumlah_min"`
	JumlahMax         *float64 `query:"jumlah_max"`
}

type TransaksiStatsRequest struct {
	TanggalMulai      string  `query:"tanggal_mulai" validate:"required" example:"2024-01-01"`
	TanggalAkhir      string  `query:"tanggal_akhir" validate:"required" example:"2024-12-31"`
	IDTipeTransaksi   *uint64 `query:"id_tipe_transaksi"`
	IDKategori        *uint64 `query:"id_kategori"`
	IDUser            *uint64 `query:"id_user"`
	IDJenisPembayaran *uint64 `query:"id_jenis_pembayaran"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
