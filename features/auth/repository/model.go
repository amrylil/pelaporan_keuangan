package repository

import (
	"pelaporan_keuangan/features/auth"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]auth.Auth, int64, error) {
	var auths []auth.Auth
	var total int64

	if err := mdl.db.Model(&auths).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&auths).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return auths, total, nil
}

func (mdl *model) Insert(newAuth auth.Auth) error {
	err := mdl.db.Create(&newAuth).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(authID uint64) (*auth.Auth, error) {
	var auth auth.Auth
	err := mdl.db.First(&auth, authID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &auth, nil
}

func (mdl *model) Update(auth auth.Auth) error {
	err := mdl.db.Updates(&auth).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(authID uint64) error {
	err := mdl.db.Delete(&auth.Auth{}, authID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
