package master_data

import (
	"pelaporan_keuangan/features/master_data/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Master_data, int64, error)
	Insert(newMaster_data Master_data) error
	SelectByID(master_dataID uint) (*Master_data, error)
	Update(master_data Master_data) error
	DeleteByID(master_dataID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResMaster_data, int64, error)
	FindByID(master_dataID uint) (*dtos.ResMaster_data, error)
	Create(newMaster_data dtos.InputMaster_data) error
	Modify(master_dataData dtos.InputMaster_data, master_dataID uint) error
	Remove(master_dataID uint) error
}

type Handler interface {
	GetMaster_data(c *gin.Context)
	Master_dataDetails(c *gin.Context)
	CreateMaster_data(c *gin.Context)
	UpdateMaster_data(c *gin.Context)
	DeleteMaster_data(c *gin.Context)
}
