package product

import (
	"blueprint_golang/features/product/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Product, int64, error)
	Insert(newProduct Product) error
	SelectByID(productID uint) (*Product, error)
	Update(product Product) error
	DeleteByID(productID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResProduct, int64, error)
	FindByID(productID uint) (*dtos.ResProduct, error)
	Create(newProduct dtos.InputProduct) error
	Modify(productData dtos.InputProduct, productID uint) error
	Remove(productID uint) error
}

type Handler interface {
	GetProducts(c *gin.Context)
	ProductDetails(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
