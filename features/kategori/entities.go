package kategori

import (
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model

	ID        uint64  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Deskripsi *string `gorm:"type:text"`
}

func (Kategori) TableName() string {
	return "kategori"
}
