package master_data

import (
	"gorm.io/gorm"
)

type Master_data struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}
type JenisPembayaran struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Nama string `gorm:"type:varchar(255)"`
}
type TipeTransaksi struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Nama string `gorm:"type:varchar(255)"`
}
type StatusTransaksi struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Nama string `gorm:"type:varchar(255)"`
}

func (Master_data) TableName() string {
	return "master_data"
}
func (JenisPembayaran) TableName() string {
	return "jenis_pembayaran"
}
func (StatusTransaksi) TableName() string {
	return "status_pembayaran"
}
func (TipeTransaksi) TableName() string {
	return "tipe_pembayaran"
}
