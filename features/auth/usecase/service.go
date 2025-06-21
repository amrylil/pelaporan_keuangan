package usecase

import (
	"pelaporan_keuangan/features/auth"
	"pelaporan_keuangan/features/auth/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model auth.Repository
}

func New(model auth.Repository) auth.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResAuth, int64, error) {
	var auths []dtos.ResAuth

	authsEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, auth := range authsEnt {
		var data dtos.ResAuth

		if err := smapping.FillStruct(&data, smapping.MapFields(auth)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		auths = append(auths, data)
	}

	return auths, total, nil
}

func (svc *service) FindByID(authID uint64) (*dtos.ResAuth, error) {
	res := dtos.ResAuth{}
	auth, err := svc.model.SelectByID(authID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if auth == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(auth))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newAuth dtos.InputAuth) error {
	auth := auth.Auth{}

	err := smapping.FillStruct(&auth, smapping.MapFields(newAuth))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	auth.ID = helpers.GenerateID()
	err = svc.model.Insert(auth)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(authData dtos.InputAuth, authID uint64) error {
	newAuth := auth.Auth{}

	err := smapping.FillStruct(&newAuth, smapping.MapFields(authData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newAuth.ID = authID
	err = svc.model.Update(newAuth)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(authID uint64) error {
	err := svc.model.DeleteByID(authID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
