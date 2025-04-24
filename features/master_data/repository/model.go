package repository

import (
	"pelaporan_keuangan/features/master_data"

	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) master_data.Repository {
	return &model{
		db: db,
	}
}
