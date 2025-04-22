package repository

import (
	"pelaporan_keuangan/features/users"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]users.Users, int64, error) {
	var userss []users.Users
	var total int64

	if err := mdl.db.Model(&userss).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&userss).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return userss, total, nil
}

func (mdl *model) Insert(newUsers users.Users) error {
	err := mdl.db.Create(&newUsers).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(usersID uint) (*users.Users, error) {
	var users users.Users
	err := mdl.db.First(&users, usersID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &users, nil
}

func (mdl *model) Update(users users.Users) error {
	err := mdl.db.Updates(&users).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(usersID uint) error {
	err := mdl.db.Delete(&users.Users{}, usersID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
