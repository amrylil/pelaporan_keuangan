package repository

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/labstack/gommon/log"
)

func (mdl *model) GetAllStatusTransaksi(page, size int) ([]master_data.StatusTransaksi, int64, error) {
	var status_transaksi []master_data.StatusTransaksi
	var total int64

	if err := mdl.db.Model(&status_transaksi).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&status_transaksi).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return status_transaksi, total, nil
}

func (mdl *model) InsertStatusTransaksi(newStatusTransaksi master_data.StatusTransaksi) error {
	err := mdl.db.Create(&newStatusTransaksi).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectStatusTransaksiByID(statusTransaksiID uint64) (*master_data.StatusTransaksi, error) {
	var status_transaksi master_data.StatusTransaksi
	err := mdl.db.First(&status_transaksi, statusTransaksiID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &status_transaksi, nil
}

func (mdl *model) UpdateStatusTransaksi(status_transaksi master_data.StatusTransaksi) error {
	err := mdl.db.Updates(&status_transaksi).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteStatusTransaksiByID(statusTransaksiID uint64) error {
	err := mdl.db.Delete(&master_data.StatusTransaksi{}, statusTransaksiID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
