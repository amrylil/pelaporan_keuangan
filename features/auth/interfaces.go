package auth

import (
	"pelaporan_keuangan/features/auth/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]User, int64, error)
	Insert(newUser User) error
	SelectByID(userID uint64) (*User, error)
	CheckEmailExist(email string) error
	GetUserByEmail(email string) (*User, error)
	Update(user User) error
	DeleteByID(userID uint64) error

	// auth
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResUser, int64, error)
	FindByID(userID uint64) (*dtos.ResUser, error)
	Create(newUser dtos.InputUser) error
	Modify(userData dtos.InputUser, userID uint64) error
	Remove(userID uint64) error

	Login(user dtos.LoginRequest) (*dtos.ResUser, error)
}

type Handler interface {
	GetUsers(c *gin.Context)
	UserDetails(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	Login(c *gin.Context)
}
