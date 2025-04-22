package users

import (
	"pelaporan_keuangan/features/users/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Users, int64, error)
	Insert(newUsers Users) error
	SelectByID(usersID uint) (*Users, error)
	Update(users Users) error
	DeleteByID(usersID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResUsers, int64, error)
	FindByID(usersID uint) (*dtos.ResUsers, error)
	Create(newUsers dtos.InputUsers) error
	Modify(usersData dtos.InputUsers, usersID uint) error
	Remove(usersID uint) error
}

type Handler interface {
	GetUsers(c *gin.Context)
	UsersDetails(c *gin.Context)
	CreateUsers(c *gin.Context)
	UpdateUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
}
