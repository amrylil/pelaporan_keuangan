package transaksi

import (
	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
