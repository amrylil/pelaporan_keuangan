package master_data

import (
	"gorm.io/gorm"
)

type Master_data struct {
	gorm.Model

	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}
type JenisPembayaran struct {
	gorm.Model

	ID        uint64  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255)"`
	Deskripsi *string `gorm:"type:text"`
}
type TipeTransaksi struct {
	gorm.Model

	ID        uint64  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255)"`
	Deskripsi *string `gorm:"type:text"`
}
type StatusTransaksi struct {
	gorm.Model

	Name      string  `gorm:"type:varchar(255)"`
	Deskripsi *string `gorm:"type:text"`
}

func (Master_data) TableName() string {
	return "master_data"
}
func (JenisPembayaran) TableName() string {
	return "jenis_pembayaran"
}
func (StatusTransaksi) TableName() string {
	return "status_transaksi"
}
func (TipeTransaksi) TableName() string {
	return "tipe_transaksi"
}
