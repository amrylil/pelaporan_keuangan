package usecase

import (
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

func (svc *service) FindAllTipeTransaksi(page, size int) ([]dtos.ResTipeTransaksi, int64, error) {
	var tipe_transaksis []dtos.ResTipeTransaksi

	tipeTransaksiData, total, err := svc.model.GetAllTipeTransaksi(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, tipe_transaksi := range tipeTransaksiData {
		var data dtos.ResTipeTransaksi

		if err := smapping.FillStruct(&data, smapping.MapFields(tipe_transaksi)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		tipe_transaksis = append(tipe_transaksis, data)
	}

	return tipe_transaksis, total, nil
}

func (svc *service) FindTipeTransaksiByID(tipeTransaksiID uint64) (*dtos.ResTipeTransaksi, error) {
	res := dtos.ResTipeTransaksi{}
	tipe_transaksi, err := svc.model.SelectTipeTransaksiByID(tipeTransaksiID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if tipe_transaksi == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(tipe_transaksi))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) CreateTipeTransaksi(newTipeTransaksi dtos.InputTipeTransaksi) error {
	tipe_transaksi := master_data.TipeTransaksi{}

	err := smapping.FillStruct(&tipe_transaksi, smapping.MapFields(newTipeTransaksi))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	tipe_transaksi.ID = helpers.GenerateID()
	err = svc.model.InsertTipeTransaksi(tipe_transaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) ModifyTipeTransaksi(tipeTransaksi dtos.InputTipeTransaksi, tipeTransaksiID uint64) error {
	newTipeTransaksi := master_data.TipeTransaksi{}

	err := smapping.FillStruct(&newTipeTransaksi, smapping.MapFields(tipeTransaksi))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newTipeTransaksi.ID = tipeTransaksiID
	err = svc.model.UpdateTipeTransaksi(newTipeTransaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) RemoveTipeTransaksi(tipeTransaksiID uint64) error {
	err := svc.model.DeleteTipeTransaksiByID(tipeTransaksiID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
