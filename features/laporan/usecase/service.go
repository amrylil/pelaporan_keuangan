package usecase

import (
	"pelaporan_keuangan/features/laporan"
	"pelaporan_keuangan/features/laporan/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model laporan.Repository
}

func New(model laporan.Repository) laporan.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResLaporan, int64, error) {
	var laporans []dtos.ResLaporan

	laporansEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, laporan := range laporansEnt {
		var data dtos.ResLaporan

		if err := smapping.FillStruct(&data, smapping.MapFields(laporan)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		laporans = append(laporans, data)
	}

	return laporans, total, nil
}

func (svc *service) FindByID(laporanID uint64) (*dtos.ResLaporan, error) {
	res := dtos.ResLaporan{}
	laporan, err := svc.model.SelectByID(laporanID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if laporan == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(laporan))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newLaporan dtos.InputLaporan) error {
	laporan := laporan.Laporan{}

	err := smapping.FillStruct(&laporan, smapping.MapFields(newLaporan))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	laporan.ID = helpers.GenerateID()
	err = svc.model.Insert(laporan)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(laporanData dtos.InputLaporan, laporanID uint64) error {
	newLaporan := laporan.Laporan{}

	err := smapping.FillStruct(&newLaporan, smapping.MapFields(laporanData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newLaporan.ID = laporanID
	err = svc.model.Update(newLaporan)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(laporanID uint64) error {
	err := svc.model.DeleteByID(laporanID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
