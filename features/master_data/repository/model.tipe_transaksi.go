package repository

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/labstack/gommon/log"
)

func (mdl *model) GetAllTipeTransaksi(page, size int) ([]master_data.TipeTransaksi, int64, error) {
	var tipe_transaksi []master_data.TipeTransaksi
	var total int64

	if err := mdl.db.Model(&tipe_transaksi).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&tipe_transaksi).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return tipe_transaksi, total, nil
}

func (mdl *model) InsertTipeTransaksi(newTipeTransaksi master_data.TipeTransaksi) error {
	err := mdl.db.Create(&newTipeTransaksi).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectTipeTransaksiByID(tipeTransaksiID uint) (*master_data.TipeTransaksi, error) {
	var tipe_transaksi master_data.TipeTransaksi
	err := mdl.db.First(&tipe_transaksi, tipeTransaksiID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &tipe_transaksi, nil
}

func (mdl *model) UpdateTipeTransaksi(tipe_transaksi master_data.TipeTransaksi) error {
	err := mdl.db.Updates(&tipe_transaksi).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteTipeTransaksiByID(tipeTransaksiID uint) error {
	err := mdl.db.Delete(&master_data.TipeTransaksi{}, tipeTransaksiID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
