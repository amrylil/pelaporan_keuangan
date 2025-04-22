package usecase

import (
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model master_data.Repository
}

func New(model master_data.Repository) master_data.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResMaster_data, int64, error) {
	var master_datas []dtos.ResMaster_data

	master_datasEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, master_data := range master_datasEnt {
		var data dtos.ResMaster_data

		if err := smapping.FillStruct(&data, smapping.MapFields(master_data)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		master_datas = append(master_datas, data)
	}

	return master_datas, total, nil
}

func (svc *service) FindByID(master_dataID uint) (*dtos.ResMaster_data, error) {
	res := dtos.ResMaster_data{}
	master_data, err := svc.model.SelectByID(master_dataID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if master_data == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(master_data))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newMaster_data dtos.InputMaster_data) error {
	master_data := master_data.Master_data{}

	err := smapping.FillStruct(&master_data, smapping.MapFields(newMaster_data))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	master_data.ID = helpers.GenerateID()
	err = svc.model.Insert(master_data)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(master_dataData dtos.InputMaster_data, master_dataID uint) error {
	newMaster_data := master_data.Master_data{}

	err := smapping.FillStruct(&newMaster_data, smapping.MapFields(master_dataData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newMaster_data.ID = master_dataID
	err = svc.model.Update(newMaster_data)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(master_dataID uint) error {
	err := svc.model.DeleteByID(master_dataID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
