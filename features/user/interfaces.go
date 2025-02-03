package user

import (
	"blueprint_golang/features/user/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Paginate(page, size int) ([]User, error)
	Insert(newUser User) error
	SelectByID(userID int) (*User, error)
	Update(user User) error
	DeleteByID(userID int) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResUser, error)
	FindByID(userID int) (*dtos.ResUser, error)
	Create(newUser dtos.InputUser) error
	Modify(userData dtos.InputUser, userID int) error
	Remove(userID int) error
}

type Handler interface {
	GetUsers(c *gin.Context)
	UserDetails(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
