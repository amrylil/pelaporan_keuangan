package usecase

import (
	"pelaporan_keuangan/features/users"
	"pelaporan_keuangan/features/users/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model users.Repository
}

func New(model users.Repository) users.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResUsers, int64, error) {
	var userss []dtos.ResUsers

	userssEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, users := range userssEnt {
		var data dtos.ResUsers

		if err := smapping.FillStruct(&data, smapping.MapFields(users)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		userss = append(userss, data)
	}

	return userss, total, nil
}

func (svc *service) FindByID(usersID uint) (*dtos.ResUsers, error) {
	res := dtos.ResUsers{}
	users, err := svc.model.SelectByID(usersID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if users == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(users))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newUsers dtos.InputUsers) error {
	users := users.Users{}

	err := smapping.FillStruct(&users, smapping.MapFields(newUsers))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	users.ID = helpers.GenerateID()
	err = svc.model.Insert(users)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(usersData dtos.InputUsers, usersID uint) error {
	newUsers := users.Users{}

	err := smapping.FillStruct(&newUsers, smapping.MapFields(usersData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newUsers.ID = usersID
	err = svc.model.Update(newUsers)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(usersID uint) error {
	err := svc.model.DeleteByID(usersID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
