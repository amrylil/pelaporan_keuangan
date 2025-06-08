package auth

import (
	"pelaporan_keuangan/features/auth/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Auth, int64, error)
	Insert(newAuth Auth) error
	SelectByID(authID uint) (*Auth, error)
	Update(auth Auth) error
	DeleteByID(authID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResAuth, int64, error)
	FindByID(authID uint) (*dtos.ResAuth, error)
	Create(newAuth dtos.InputAuth) error
	Modify(authData dtos.InputAuth, authID uint) error
	Remove(authID uint) error
}

type Handler interface {
	GetAuth(c *gin.Context)
	AuthDetails(c *gin.Context)
	CreateAuth(c *gin.Context)
	UpdateAuth(c *gin.Context)
	DeleteAuth(c *gin.Context)
}
