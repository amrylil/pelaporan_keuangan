package kategori

import (
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Kategori) TableName() string {
	return "kategori"
}
