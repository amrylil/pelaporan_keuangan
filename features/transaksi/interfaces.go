package transaksi

import (
	"pelaporan_keuangan/features/transaksi/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Transaksi, int64, error)
	Insert(newTransaksi Transaksi) error
	SelectByID(transaksiID uint) (*Transaksi, error)
	Update(transaksi Transaksi) error
	UpdatePartial(transaksiID uint, updates map[string]interface{}) error // NEW
	DeleteByID(transaksiID uint) error
	UpdateStatus(transaksiID uint, statusID int) error
	GetWithFilter(filter dtos.TransaksiListRequest) ([]Transaksi, int64, error) // NEW
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResTransaksi, int64, error)
	FindByID(transaksiID uint) (*dtos.ResTransaksi, error)
	Create(newTransaksi dtos.InputTransaksi) error
	Modify(transaksiData dtos.InputTransaksi, transaksiID uint) error
	ModifyPartial(updateData dtos.UpdateTransaksiRequest) error // NEW
	Remove(transaksiID uint) error
	ModifyStatus(transaksiID uint, statusID int) error
	FindWithFilter(filter dtos.TransaksiListRequest) ([]dtos.ResTransaksi, int64, error) // NEW
}

type Handler interface {
	GetTransaksi(c *gin.Context)
	TransaksiDetails(c *gin.Context)
	CreateTransaksi(c *gin.Context)
	UpdateTransaksi(c *gin.Context)
	DeleteTransaksi(c *gin.Context)
}
