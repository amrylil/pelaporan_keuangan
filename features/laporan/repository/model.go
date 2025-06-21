package repository

import (
	"pelaporan_keuangan/features/laporan"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) laporan.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]laporan.Laporan, int64, error) {
	var laporans []laporan.Laporan
	var total int64

	if err := mdl.db.Model(&laporans).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&laporans).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return laporans, total, nil
}

func (mdl *model) Insert(newLaporan laporan.Laporan) error {
	err := mdl.db.Create(&newLaporan).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(laporanID uint64) (*laporan.Laporan, error) {
	var laporan laporan.Laporan
	err := mdl.db.First(&laporan, laporanID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &laporan, nil
}

func (mdl *model) Update(laporan laporan.Laporan) error {
	err := mdl.db.Updates(&laporan).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(laporanID uint64) error {
	err := mdl.db.Delete(&laporan.Laporan{}, laporanID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
