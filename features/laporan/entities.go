package laporan

import (
	"gorm.io/gorm"
)

type Laporan struct {
	gorm.Model

	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Laporan) TableName() string {
	return "laporan"
}
