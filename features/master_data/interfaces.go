package master_data

import (
	"pelaporan_keuangan/features/master_data/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAllJenisPembayaran(page, size int) ([]JenisPembayaran, int64, error)
	InsertJenisPembayaran(newJenisPembayaran JenisPembayaran) error
	SelectJenisPembayaranByID(jenisPembayaranID uint) (*JenisPembayaran, error)
	UpdateJenisPembayaran(jenisPembayaran JenisPembayaran) error
	DeleteJenisPembayaranByID(jenisPembayaranID uint) error

	GetAllTipeTransaksi(page, size int) ([]TipeTransaksi, int64, error)
	InsertTipeTransaksi(newTipeTransaksi TipeTransaksi) error
	SelectTipeTransaksiByID(tipeTransaksiID uint) (*TipeTransaksi, error)
	UpdateTipeTransaksi(tipeTransaksi TipeTransaksi) error
	DeleteTipeTransaksiByID(tipeTransaksi uint) error

	GetAllStatusTransaksi(page, size int) ([]StatusTransaksi, int64, error)
	InsertStatusTransaksi(newStatusTransaksi StatusTransaksi) error
	SelectStatusTransaksiByID(statusTransaksiID uint) (*StatusTransaksi, error)
	UpdateStatusTransaksi(statusTransaksi StatusTransaksi) error
	DeleteStatusTransaksiByID(statusTransaksi uint) error
}

type Usecase interface {
	FindAllTipeTransaksi(page, size int) ([]dtos.ResTipeTransaksi, int64, error)
	FindTipeTransaksiByID(tipeTransaksiID uint) (*dtos.ResTipeTransaksi, error)
	CreateTipeTransaksi(newTipeTransaksi dtos.InputTipeTransaksi) error
	ModifyTipeTransaksi(tipeTransaksiData dtos.InputTipeTransaksi, TipeTransaksiID uint) error
	RemoveTipeTransaksi(TipeTransaksiID uint) error

	FindAllJenisPembayaran(page, size int) ([]dtos.ResJenisPembayaran, int64, error)
	FindJenisPembayaranByID(jenisPembayaranID uint) (*dtos.ResJenisPembayaran, error)
	CreateJenisPembayaran(newJenisPembayaran dtos.InputJenisPembayaran) error
	ModifyJenisPembayaran(jenisPembayaranData dtos.InputJenisPembayaran, JenisPembayaranID uint) error
	RemoveJenisPembayaran(JenisPembayaranID uint) error

	FindAllStatusTransaksi(page, size int) ([]dtos.ResStatusTransaksi, int64, error)
	FindStatusTransaksiByID(statusTransaksiID uint) (*dtos.ResStatusTransaksi, error)
	CreateStatusTransaksi(newStatusTransaksi dtos.InputStatusTransaksi) error
	ModifyStatusTransaksi(statusTransaksiData dtos.InputStatusTransaksi, StatusTransaksiID uint) error
	RemoveStatusTransaksi(statusTransaksiID uint) error
}

type Handler interface {
	GetJenisPembayaran(c *gin.Context)
	JenisPembayaranDetails(c *gin.Context)
	CreateJenisPembayaran(c *gin.Context)
	UpdateJenisPembayaran(c *gin.Context)
	DeleteJenisPembayaran(c *gin.Context)

	GetTipeTransaksi(c *gin.Context)
	TipeTransaksiDetails(c *gin.Context)
	CreateTipeTransaksi(c *gin.Context)
	UpdateTipeTransaksi(c *gin.Context)
	DeleteTipeTransaksi(c *gin.Context)

	GetStatusTransaksi(c *gin.Context)
	StatusTransaksiDetails(c *gin.Context)
	CreateStatusTransaksi(c *gin.Context)
	UpdateStatusTransaksi(c *gin.Context)
	DeleteStatusTransaksi(c *gin.Context)
}
