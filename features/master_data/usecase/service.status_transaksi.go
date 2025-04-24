package usecase

import (
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

func (svc *service) FindAllStatusTransaksi(page, size int) ([]dtos.ResStatusTransaksi, int64, error) {
	var status_transaksis []dtos.ResStatusTransaksi

	statusTransaksiData, total, err := svc.model.GetAllStatusTransaksi(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, status_transaksi := range statusTransaksiData {
		var data dtos.ResStatusTransaksi

		if err := smapping.FillStruct(&data, smapping.MapFields(status_transaksi)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		status_transaksis = append(status_transaksis, data)
	}

	return status_transaksis, total, nil
}

func (svc *service) FindStatusTransaksiByID(statusTransaksiID uint) (*dtos.ResStatusTransaksi, error) {
	res := dtos.ResStatusTransaksi{}
	status_transaksi, err := svc.model.SelectStatusTransaksiByID(statusTransaksiID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if status_transaksi == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(status_transaksi))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) CreateStatusTransaksi(newStatusTransaksi dtos.InputStatusTransaksi) error {
	status_transaksi := master_data.StatusTransaksi{}

	err := smapping.FillStruct(&status_transaksi, smapping.MapFields(newStatusTransaksi))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	status_transaksi.ID = helpers.GenerateID()
	err = svc.model.InsertStatusTransaksi(status_transaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) ModifyStatusTransaksi(statusTransaksi dtos.InputStatusTransaksi, statusTransaksiID uint) error {
	newStatusTransaksi := master_data.StatusTransaksi{}

	err := smapping.FillStruct(&newStatusTransaksi, smapping.MapFields(statusTransaksi))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newStatusTransaksi.ID = statusTransaksiID
	err = svc.model.UpdateStatusTransaksi(newStatusTransaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) RemoveStatusTransaksi(statusTransaksiID uint) error {
	err := svc.model.DeleteStatusTransaksiByID(statusTransaksiID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
