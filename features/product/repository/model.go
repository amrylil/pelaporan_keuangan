package repository

import (
	"blueprint_golang/features/product"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]product.Product, int64, error) {
	var products []product.Product
	var total int64

	if err := mdl.db.Model(&products).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&products).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return products, total, nil
}

func (mdl *model) Insert(newProduct product.Product) error {
	err := mdl.db.Create(&newProduct).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(productID uint) (*product.Product, error) {
	var product product.Product
	err := mdl.db.First(&product, productID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &product, nil
}

func (mdl *model) Update(product product.Product) error {
	err := mdl.db.Updates(&product).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(productID uint) error {
	err := mdl.db.Delete(&product.Product{}, productID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
