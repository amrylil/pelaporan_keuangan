package usecase

import (
	"pelaporan_keuangan/features/transaksi"
	"pelaporan_keuangan/features/transaksi/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model transaksi.Repository
}

func New(model transaksi.Repository) transaksi.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResTransaksi, int64, error) {
	var transaksis []dtos.ResTransaksi

	transaksisEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, transaksi := range transaksisEnt {
		var data dtos.ResTransaksi

		if err := smapping.FillStruct(&data, smapping.MapFields(transaksi)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		transaksis = append(transaksis, data)
	}

	return transaksis, total, nil
}

func (svc *service) FindByID(transaksiID uint) (*dtos.ResTransaksi, error) {
	res := dtos.ResTransaksi{}
	transaksi, err := svc.model.SelectByID(transaksiID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if transaksi == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(transaksi))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newTransaksi dtos.InputTransaksi) error {
	transaksi := transaksi.Transaksi{}

	err := smapping.FillStruct(&transaksi, smapping.MapFields(newTransaksi))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	transaksi.ID = helpers.GenerateID()
	err = svc.model.Insert(transaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(transaksiData dtos.InputTransaksi, transaksiID uint) error {
	newTransaksi := transaksi.Transaksi{}

	err := smapping.FillStruct(&newTransaksi, smapping.MapFields(transaksiData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newTransaksi.ID = transaksiID
	err = svc.model.Update(newTransaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(transaksiID uint) error {
	err := svc.model.DeleteByID(transaksiID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
