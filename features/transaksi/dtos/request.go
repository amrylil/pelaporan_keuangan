package dtos

type InputTransaksi struct {
	Tanggal           string  `json:"tanggal"            binding:"required,datetime=2006-01-02"` // yyyy-mm-dd
	NamaTransaksi     string  `json:"nama_transaksi"     binding:"required,max=255"`
	Jumlah            float64 `json:"jumlah"             binding:"required,gt=0"`
	Keterangan        string  `json:"keterangan"         binding:"omitempty"`
	BuktiTransaksi    string  `json:"bukti_transaksi"    binding:"omitempty,url"` // atau path lokal
	IDTipeTransaksi   uint    `json:"id_tipe_transaksi"  binding:"required"`
	IDKategori        uint    `json:"id_kategori"        binding:"required"`
	IDJenisPembayaran uint    `json:"id_jenis_pembayaran" binding:"required"`
}

// UpdateTransaksiRequest - DTO untuk update transaksi
type UpdateTransaksiRequest struct {
	ID                *uint    `json:"id" validate:"required"`
	Tanggal           *string  `json:"tanggal,omitempty"`
	IDTipeTransaksi   *uint    `json:"id_tipe_transaksi,omitempty"`
	Jumlah            *float64 `json:"jumlah,omitempty"`
	Keterangan        *string  `json:"keterangan,omitempty"`
	BuktiTransaksi    *string  `json:"bukti_transaksi,omitempty"`
	IDStatusTransaksi *uint    `json:"id_status_transaksi,omitempty"`
	KomentarManajer   *string  `json:"komentar_manajer,omitempty"`
	IDKategori        *uint    `json:"id_kategori,omitempty"`
	IDUser            *uint    `json:"id_user,omitempty"`
	IDJenisPembayaran *uint    `json:"id_jenis_pembayaran,omitempty"`
}

type TransaksiListRequest struct {
	Page              int      `query:"page" validate:"min=1" example:"1"`
	Limit             int      `query:"limit" validate:"min=1,max=100" example:"10"`
	Search            string   `query:"search" example:"makan"`
	IDTipeTransaksi   *uint    `query:"id_tipe_transaksi"`
	IDStatusTransaksi *uint    `query:"id_status_transaksi"`
	IDKategori        *uint    `query:"id_kategori"`
	IDUser            *uint    `query:"id_user"`
	IDJenisPembayaran *uint    `query:"id_jenis_pembayaran"`
	TanggalMulai      *string  `query:"tanggal_mulai" example:"2024-01-01"`
	TanggalAkhir      *string  `query:"tanggal_akhir" example:"2024-12-31"`
	JumlahMin         *float64 `query:"jumlah_min"`
	JumlahMax         *float64 `query:"jumlah_max"`
}

type TransaksiStatsRequest struct {
	TanggalMulai      string `query:"tanggal_mulai" validate:"required" example:"2024-01-01"`
	TanggalAkhir      string `query:"tanggal_akhir" validate:"required" example:"2024-12-31"`
	IDTipeTransaksi   *uint  `query:"id_tipe_transaksi"`
	IDKategori        *uint  `query:"id_kategori"`
	IDUser            *uint  `query:"id_user"`
	IDJenisPembayaran *uint  `query:"id_jenis_pembayaran"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
