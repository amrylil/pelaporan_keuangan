package repository

import (
	"pelaporan_keuangan/features/kategori"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) kategori.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]kategori.Kategori, int64, error) {
	var kategoris []kategori.Kategori
	var total int64

	if err := mdl.db.Model(&kategoris).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&kategoris).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return kategoris, total, nil
}

func (mdl *model) Insert(newKategori kategori.Kategori) error {
	err := mdl.db.Create(&newKategori).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(kategoriID uint64) (*kategori.Kategori, error) {
	var kategori kategori.Kategori
	err := mdl.db.First(&kategori, kategoriID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &kategori, nil
}

func (mdl *model) Update(kategori kategori.Kategori) error {
	err := mdl.db.Updates(&kategori).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(kategoriID uint64) error {
	err := mdl.db.Delete(&kategori.Kategori{}, kategoriID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
