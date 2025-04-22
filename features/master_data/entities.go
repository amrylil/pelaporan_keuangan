package master_data

import (
	"gorm.io/gorm"
)

type Master_data struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Master_data) TableName() string {
	return "master_data"
}
