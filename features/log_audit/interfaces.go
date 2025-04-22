package log_audit

import (
	"pelaporan_keuangan/features/log_audit/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Log_audit, int64, error)
	Insert(newLog_audit Log_audit) error
	SelectByID(log_auditID uint) (*Log_audit, error)
	Update(log_audit Log_audit) error
	DeleteByID(log_auditID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResLog_audit, int64, error)
	FindByID(log_auditID uint) (*dtos.ResLog_audit, error)
	Create(newLog_audit dtos.InputLog_audit) error
	Modify(log_auditData dtos.InputLog_audit, log_auditID uint) error
	Remove(log_auditID uint) error
}

type Handler interface {
	GetLog_audit(c *gin.Context)
	Log_auditDetails(c *gin.Context)
	CreateLog_audit(c *gin.Context)
	UpdateLog_audit(c *gin.Context)
	DeleteLog_audit(c *gin.Context)
}
