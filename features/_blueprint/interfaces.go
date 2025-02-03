package _blueprint

import (
	"blueprint_golang/features/_blueprint/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Paginate(page, size int) ([]Placeholder, error)
	Insert(newPlaceholder Placeholder) error
	SelectByID(placeholderID int) (*Placeholder, error)
	Update(placeholder Placeholder) error
	DeleteByID(placeholderID int) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResPlaceholder, error)
	FindByID(placeholderID int) (*dtos.ResPlaceholder, error)
	Create(newPlaceholder dtos.InputPlaceholder) error
	Modify(placeholderData dtos.InputPlaceholder, placeholderID int) error
	Remove(placeholderID int) error
}

type Handler interface {
	GetPlaceholders(c *gin.Context)
	PlaceholderDetails(c *gin.Context)
	CreatePlaceholder(c *gin.Context)
	UpdatePlaceholder(c *gin.Context)
	DeletePlaceholder(c *gin.Context)
}
