package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Product) TableName() string {
	return "products"
}
