package kategori

import (
	"pelaporan_keuangan/features/kategori/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Kategori, int64, error)
	Insert(newKategori Kategori) error
	SelectByID(kategoriID uint64) (*Kategori, error)
	Update(kategori Kategori) error
	DeleteByID(kategoriID uint64) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResKategori, int64, error)
	FindByID(kategoriID uint64) (*dtos.ResKategori, error)
	Create(newKategori dtos.InputKategori) error
	Modify(kategoriData dtos.InputKategori, kategoriID uint64) error
	Remove(kategoriID uint64) error
}

type Handler interface {
	GetKategori(c *gin.Context)
	KategoriDetails(c *gin.Context)
	CreateKategori(c *gin.Context)
	UpdateKategori(c *gin.Context)
	DeleteKategori(c *gin.Context)
}
