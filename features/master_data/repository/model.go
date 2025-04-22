package repository

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/labstack/gommon/log"
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

func (mdl *model) GetAll(page, size int) ([]master_data.Master_data, int64, error) {
	var master_datas []master_data.Master_data
	var total int64

	if err := mdl.db.Model(&master_datas).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&master_datas).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return master_datas, total, nil
}

func (mdl *model) Insert(newMaster_data master_data.Master_data) error {
	err := mdl.db.Create(&newMaster_data).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(master_dataID uint) (*master_data.Master_data, error) {
	var master_data master_data.Master_data
	err := mdl.db.First(&master_data, master_dataID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &master_data, nil
}

func (mdl *model) Update(master_data master_data.Master_data) error {
	err := mdl.db.Updates(&master_data).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(master_dataID uint) error {
	err := mdl.db.Delete(&master_data.Master_data{}, master_dataID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
