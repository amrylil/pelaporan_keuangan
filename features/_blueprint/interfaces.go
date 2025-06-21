package _blueprint

import (
	"pelaporan_keuangan/features/_blueprint/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Placeholder, int64, error)
	Insert(newPlaceholder Placeholder) error
	SelectByID(placeholderID uint64) (*Placeholder, error)
	Update(placeholder Placeholder) error
	DeleteByID(placeholderID uint64) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResPlaceholder, int64, error)
	FindByID(placeholderID uint64) (*dtos.ResPlaceholder, error)
	Create(newPlaceholder dtos.InputPlaceholder) error
	Modify(placeholderData dtos.InputPlaceholder, placeholderID uint64) error
	Remove(placeholderID uint64) error
}

type Handler interface {
	GetPlaceholder(c *gin.Context)
	PlaceholderDetails(c *gin.Context)
	CreatePlaceholder(c *gin.Context)
	UpdatePlaceholder(c *gin.Context)
	DeletePlaceholder(c *gin.Context)
}
