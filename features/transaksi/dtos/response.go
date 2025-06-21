package dtos

import "time"

type ResTransaksi struct {
	ID                uint      `json:"id"`
	Tanggal           string    `json:"tanggal"`
	IDTipeTransaksi   uint      `json:"id_tipe_transaksi"`
	Jumlah            float64   `json:"jumlah"`
	Keterangan        string    `json:"keterangan"`
	BuktiTransaksi    string    `json:"bukti_transaksi"`
	IDStatusTransaksi uint      `json:"id_status_transaksi"`
	KomentarManajer   string    `json:"komentar_manajer"`
	IDKategori        uint      `json:"id_kategori"`
	IDUser            uint      `json:"id_user"`
	IDJenisPembayaran uint      `json:"id_jenis_pembayaran"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// TransaksiDetailResponse - DTO untuk response transaksi dengan relasi
type TransaksiDetailResponse struct {
	ID              uint                    `json:"id"`
	Tanggal         string                  `json:"tanggal"`
	Jumlah          float64                 `json:"jumlah"`
	Keterangan      string                  `json:"keterangan"`
	BuktiTransaksi  string                  `json:"bukti_transaksi"`
	KomentarManajer string                  `json:"komentar_manajer"`
	TipeTransaksi   TipeTransaksiResponse   `json:"tipe_transaksi"`
	StatusTransaksi StatusTransaksiResponse `json:"status_transaksi"`
	Kategori        KategoriResponse        `json:"kategori"`
	User            UserResponse            `json:"user"`
	JenisPembayaran JenisPembayaranResponse `json:"jenis_pembayaran"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

// DTOs untuk relasi (sesuaikan dengan struktur model Anda)
type TipeTransaksiResponse struct {
	ID   uint64 `json:"id"`
	Nama string `json:"nama"`
}

type StatusTransaksiResponse struct {
	ID   uint64 `json:"id"`
	Nama string `json:"nama"`
}

type KategoriResponse struct {
	ID   uint64 `json:"id"`
	Nama string `json:"nama"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

type JenisPembayaranResponse struct {
	ID   uint64 `json:"id"`
	Nama string `json:"nama"`
}

type TransaksiListResponse struct {
	Data       []ResTransaksi     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type TransaksiStatsResponse struct {
	TotalTransaksi     int64                     `json:"total_transaksi"`
	TotalJumlah        float64                   `json:"total_jumlah"`
	RataRataJumlah     float64                   `json:"rata_rata_jumlah"`
	JumlahTertinggi    float64                   `json:"jumlah_tertinggi"`
	JumlahTerendah     float64                   `json:"jumlah_terendah"`
	PerKategori        []StatsPerKategori        `json:"per_kategori"`
	PerTipeTransaksi   []StatsPerTipeTransaksi   `json:"per_tipe_transaksi"`
	PerJenisPembayaran []StatsPerJenisPembayaran `json:"per_jenis_pembayaran"`
}

type StatsPerKategori struct {
	IDKategori     uint    `json:"id_kategori"`
	NamaKategori   string  `json:"nama_kategori"`
	TotalTransaksi int64   `json:"total_transaksi"`
	TotalJumlah    float64 `json:"total_jumlah"`
}

type StatsPerTipeTransaksi struct {
	IDTipeTransaksi   uint    `json:"id_tipe_transaksi"`
	NamaTipeTransaksi string  `json:"nama_tipe_transaksi"`
	TotalTransaksi    int64   `json:"total_transaksi"`
	TotalJumlah       float64 `json:"total_jumlah"`
}

type StatsPerJenisPembayaran struct {
	IDJenisPembayaran   uint    `json:"id_jenis_pembayaran"`
	NamaJenisPembayaran string  `json:"nama_jenis_pembayaran"`
	TotalTransaksi      int64   `json:"total_transaksi"`
	TotalJumlah         float64 `json:"total_jumlah"`
}

// ManagerCommentRequest - DTO untuk update komentar manajer
type ManagerCommentRequest struct {
	ID              uint   `json:"id" validate:"required"`
	KomentarManajer string `json:"komentar_manajer" validate:"required"`
}
