package usecase

import (
	"blueprint_golang/features/product"
	"blueprint_golang/features/product/dtos"
	"blueprint_golang/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model product.Repository
}

func New(model product.Repository) product.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResProduct, int64, error) {
	var products []dtos.ResProduct

	productsEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, product := range productsEnt {
		var data dtos.ResProduct

		if err := smapping.FillStruct(&data, smapping.MapFields(product)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		products = append(products, data)
	}

	return products, total, nil
}

func (svc *service) FindByID(productID uint) (*dtos.ResProduct, error) {
	res := dtos.ResProduct{}
	product, err := svc.model.SelectByID(productID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(product))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newProduct dtos.InputProduct) error {
	product := product.Product{}

	err := smapping.FillStruct(&product, smapping.MapFields(newProduct))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	product.ID = helpers.GenerateID()
	err = svc.model.Insert(product)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(productData dtos.InputProduct, productID uint) error {
	newProduct := product.Product{}

	err := smapping.FillStruct(&newProduct, smapping.MapFields(productData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newProduct.ID = productID
	err = svc.model.Update(newProduct)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(productID uint) error {
	err := svc.model.DeleteByID(productID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
