package log_audit

import (
	"gorm.io/gorm"
)

type Log_audit struct {
	gorm.Model

	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Log_audit) TableName() string {
	return "log_audit"
}
