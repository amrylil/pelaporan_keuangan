package users

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	ID       uint64 `gorm:"primaryKey"`
	Nama     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
	TipeAkun string `gorm:"column:tipe_akun"`
}

func (Users) TableName() string {
	return "users"
}
