package users

import (
	"pelaporan_keuangan/features/users/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Users, int64, error)
	Insert(newUsers Users) error
	SelectByID(usersID uint64) (*Users, error)
	Update(users Users) error
	DeleteByID(usersID uint64) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResUsers, int64, error)
	FindByID(usersID uint64) (*dtos.ResUsers, error)
	Create(newUsers dtos.InputUsers) error
	Modify(usersData dtos.InputUsers, usersID uint64) error
	Remove(usersID uint64) error
}

type Handler interface {
	GetUsers(c *gin.Context)
	UsersDetails(c *gin.Context)
	CreateUsers(c *gin.Context)
	UpdateUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
}
