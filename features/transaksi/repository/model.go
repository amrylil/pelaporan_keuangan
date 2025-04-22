package repository

import (
	"pelaporan_keuangan/features/transaksi"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaksi.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]transaksi.Transaksi, int64, error) {
	var transaksis []transaksi.Transaksi
	var total int64

	if err := mdl.db.Model(&transaksis).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&transaksis).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return transaksis, total, nil
}

func (mdl *model) Insert(newTransaksi transaksi.Transaksi) error {
	err := mdl.db.Create(&newTransaksi).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(transaksiID uint) (*transaksi.Transaksi, error) {
	var transaksi transaksi.Transaksi
	err := mdl.db.First(&transaksi, transaksiID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &transaksi, nil
}

func (mdl *model) Update(transaksi transaksi.Transaksi) error {
	err := mdl.db.Updates(&transaksi).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(transaksiID uint) error {
	err := mdl.db.Delete(&transaksi.Transaksi{}, transaksiID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
