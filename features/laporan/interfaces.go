package laporan

import (
	"pelaporan_keuangan/features/laporan/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Laporan, int64, error)
	Insert(newLaporan Laporan) error
	SelectByID(laporanID uint) (*Laporan, error)
	Update(laporan Laporan) error
	DeleteByID(laporanID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResLaporan, int64, error)
	FindByID(laporanID uint) (*dtos.ResLaporan, error)
	Create(newLaporan dtos.InputLaporan) error
	Modify(laporanData dtos.InputLaporan, laporanID uint) error
	Remove(laporanID uint) error
}

type Handler interface {
	GetLaporan(c *gin.Context)
	LaporanDetails(c *gin.Context)
	CreateLaporan(c *gin.Context)
	UpdateLaporan(c *gin.Context)
	DeleteLaporan(c *gin.Context)
}
