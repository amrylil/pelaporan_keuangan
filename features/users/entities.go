package users

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	IDRole   string `gorm:"column:id_role"`
}

func (Users) TableName() string {
	return "users"
}
