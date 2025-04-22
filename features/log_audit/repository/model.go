package repository

import (
	"pelaporan_keuangan/features/log_audit"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) log_audit.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]log_audit.Log_audit, int64, error) {
	var log_audits []log_audit.Log_audit
	var total int64

	if err := mdl.db.Model(&log_audits).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&log_audits).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return log_audits, total, nil
}

func (mdl *model) Insert(newLog_audit log_audit.Log_audit) error {
	err := mdl.db.Create(&newLog_audit).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(log_auditID uint) (*log_audit.Log_audit, error) {
	var log_audit log_audit.Log_audit
	err := mdl.db.First(&log_audit, log_auditID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &log_audit, nil
}

func (mdl *model) Update(log_audit log_audit.Log_audit) error {
	err := mdl.db.Updates(&log_audit).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(log_auditID uint) error {
	err := mdl.db.Delete(&log_audit.Log_audit{}, log_auditID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
