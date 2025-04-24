package repository

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/labstack/gommon/log"
)

func (mdl *model) GetAllJenisPembayaran(page, size int) ([]master_data.JenisPembayaran, int64, error) {
	var jenis_pembayaran []master_data.JenisPembayaran
	var total int64

	if err := mdl.db.Model(&jenis_pembayaran).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&jenis_pembayaran).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return jenis_pembayaran, total, nil
}

func (mdl *model) InsertJenisPembayaran(newJenisPembayaran master_data.JenisPembayaran) error {
	err := mdl.db.Create(&newJenisPembayaran).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectJenisPembayaranByID(jenisPembayaranID uint) (*master_data.JenisPembayaran, error) {
	var jenis_pembayaran master_data.JenisPembayaran
	err := mdl.db.First(&jenis_pembayaran, jenisPembayaranID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &jenis_pembayaran, nil
}

func (mdl *model) UpdateJenisPembayaran(jenis_pembayaran master_data.JenisPembayaran) error {
	err := mdl.db.Updates(&jenis_pembayaran).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteJenisPembayaranByID(jenisPembayaranID uint) error {
	err := mdl.db.Delete(&master_data.JenisPembayaran{}, jenisPembayaranID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
