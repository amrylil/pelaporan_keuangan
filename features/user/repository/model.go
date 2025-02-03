package repository

import (
	"blueprint_golang/features/user"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) ([]user.User, error) {
	var users []user.User

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&users).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (mdl *model) Insert(newUser user.User) error {
	err := mdl.db.Create(&newUser).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(userID int) (*user.User, error) {
	var user user.User
	err := mdl.db.First(&user, userID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &user, nil
}

func (mdl *model) Update(user user.User) error {
	err := mdl.db.Updates(&user).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(userID int) error {
	err := mdl.db.Delete(&user.User{}, userID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
