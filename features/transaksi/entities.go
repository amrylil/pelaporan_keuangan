package transaksi

import (
	"pelaporan_keuangan/features/kategori"
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/users"

	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model

	ID uint64 `gorm:"primaryKey"`

	Tanggal           string  `gorm:"column:tanggal;type:date"`
	IDTipeTransaksi   uint64  `gorm:"column:id_tipe_transaksi"`
	Jumlah            float64 `gorm:"column:jumlah"`
	Keterangan        string  `gorm:"column:keterangan;type:text"`
	BuktiTransaksi    string  `gorm:"column:bukti_transaksi"`
	IDStatusTransaksi uint64  `gorm:"column:id_status_transaksi"`
	KomentarManajer   string  `gorm:"column:komentar_manajer;type:text"`
	IDKategori        uint64  `gorm:"column:id_kategori"`
	IDUser            uint64  `gorm:"column:id_user"`
	IDJenisPembayaran uint64  `gorm:"column:id_jenis_pembayaran"`

	TipeTransaksi   master_data.TipeTransaksi   `gorm:"foreignKey:IDTipeTransaksi"`
	StatusTransaksi master_data.StatusTransaksi `gorm:"foreignKey:IDStatusTransaksi"`
	Kategori        kategori.Kategori           `gorm:"foreignKey:IDKategori"`
	User            users.User                  `gorm:"foreignKey:IDUser"`
	JenisPembayaran master_data.JenisPembayaran `gorm:"foreignKey:IDJenisPembayaran"`
}

func (Transaksi) TableName() string {
	return "transaksi"

}
