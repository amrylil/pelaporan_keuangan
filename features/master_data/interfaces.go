package master_data

import (
	"pelaporan_keuangan/features/master_data/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAllJenisPembayaran(page, size int) ([]JenisPembayaran, int64, error)
	InsertJenisPembayaran(newJenisPembayaran JenisPembayaran) error
	SelectJenisPembayaranByID(jenisPembayaranID uint64) (*JenisPembayaran, error)
	UpdateJenisPembayaran(jenisPembayaran JenisPembayaran) error
	DeleteJenisPembayaranByID(jenisPembayaranID uint64) error

	GetAllTipeTransaksi(page, size int) ([]TipeTransaksi, int64, error)
	InsertTipeTransaksi(newTipeTransaksi TipeTransaksi) error
	SelectTipeTransaksiByID(tipeTransaksiID uint64) (*TipeTransaksi, error)
	UpdateTipeTransaksi(tipeTransaksi TipeTransaksi) error
	DeleteTipeTransaksiByID(tipeTransaksi uint64) error

	GetAllStatusTransaksi(page, size int) ([]StatusTransaksi, int64, error)
	InsertStatusTransaksi(newStatusTransaksi StatusTransaksi) error
	SelectStatusTransaksiByID(statusTransaksiID uint64) (*StatusTransaksi, error)
	UpdateStatusTransaksi(statusTransaksi StatusTransaksi) error
	DeleteStatusTransaksiByID(statusTransaksi uint64) error
}

type Usecase interface {
	FindAllTipeTransaksi(page, size int) ([]dtos.ResTipeTransaksi, int64, error)
	FindTipeTransaksiByID(tipeTransaksiID uint64) (*dtos.ResTipeTransaksi, error)
	CreateTipeTransaksi(newTipeTransaksi dtos.InputTipeTransaksi) error
	ModifyTipeTransaksi(tipeTransaksiData dtos.InputTipeTransaksi, TipeTransaksiID uint64) error
	RemoveTipeTransaksi(TipeTransaksiID uint64) error

	FindAllJenisPembayaran(page, size int) ([]dtos.ResJenisPembayaran, int64, error)
	FindJenisPembayaranByID(jenisPembayaranID uint64) (*dtos.ResJenisPembayaran, error)
	CreateJenisPembayaran(newJenisPembayaran dtos.InputJenisPembayaran) error
	ModifyJenisPembayaran(jenisPembayaranData dtos.InputJenisPembayaran, JenisPembayaranID uint64) error
	RemoveJenisPembayaran(JenisPembayaranID uint64) error

	FindAllStatusTransaksi(page, size int) ([]dtos.ResStatusTransaksi, int64, error)
	FindStatusTransaksiByID(statusTransaksiID uint64) (*dtos.ResStatusTransaksi, error)
	CreateStatusTransaksi(newStatusTransaksi dtos.InputStatusTransaksi) error
	ModifyStatusTransaksi(statusTransaksiData dtos.InputStatusTransaksi, StatusTransaksiID uint64) error
	RemoveStatusTransaksi(statusTransaksiID uint64) error
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
