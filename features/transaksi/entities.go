package transaksi

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model // ini akan menyertakan: ID, CreatedAt, UpdatedAt, DeletedAt

	Tanggal           string  `gorm:"column:tanggal;type:date"`
	IDTipeTransaksi   uint    `gorm:"column:id_tipe_transaksi"`
	Jumlah            float64 `gorm:"column:jumlah"`
	Keterangan        string  `gorm:"column:keterangan;type:text"`
	BuktiTransaksi    string  `gorm:"column:bukti_transaksi"`
	IDStatusTransaksi uint    `gorm:"column:id_status_transaksi"`
	KomentarManajer   string  `gorm:"column:komentar_manajer;type:text"`
	IDKategori        uint    `gorm:"column:id_kategori"`
	IDUser            uint    `gorm:"column:id_user"`
	IDJenisPembayaran uint    `gorm:"column:id_jenis_pembayaran"`

	// Relasi opsional
	// TipeTransaksi   TipeTransaksi   `gorm:"foreignKey:IDTipeTransaksi"`
	// StatusTransaksi StatusTransaksi `gorm:"foreignKey:IDStatusTransaksi"`
	// Kategori        Kategori        `gorm:"foreignKey:IDKategori"`
	// User            User            `gorm:"foreignKey:IDUser"`
	// JenisPembayaran JenisPembayaran `gorm:"foreignKey:IDJenisPembayaran"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
