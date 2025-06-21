package auth

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model

	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	TipeAkun string `gorm:"type:varchar(20)"`
}

func (Auth) TableName() string {
	return "auth"
}
