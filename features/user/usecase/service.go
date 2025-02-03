package usecase

import (
	"blueprint_golang/features/user"
	"blueprint_golang/features/user/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model user.Repository
}

func New(model user.Repository) user.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResUser, error) {
	var users []dtos.ResUser

	usersEnt, err := svc.model.Paginate(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	for _, user := range usersEnt {
		var data dtos.ResUser

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
		}

		users = append(users, data)
	}

	return users, nil
}

func (svc *service) FindByID(userID int) (*dtos.ResUser, error) {
	res := dtos.ResUser{}
	user, err := svc.model.SelectByID(userID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newUser dtos.InputUser) error {
	user := user.User{}

	err := smapping.FillStruct(&user, smapping.MapFields(newUser))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	err = svc.model.Insert(user)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(userData dtos.InputUser, userID int) error {
	newUser := user.User{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newUser.ID = userID
	err = svc.model.Update(newUser)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(userID int) error {
	err := svc.model.DeleteByID(userID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
