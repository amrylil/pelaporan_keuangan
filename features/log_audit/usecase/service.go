package usecase

import (
	"pelaporan_keuangan/features/log_audit"
	"pelaporan_keuangan/features/log_audit/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model log_audit.Repository
}

func New(model log_audit.Repository) log_audit.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResLog_audit, int64, error) {
	var log_audits []dtos.ResLog_audit

	log_auditsEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, log_audit := range log_auditsEnt {
		var data dtos.ResLog_audit

		if err := smapping.FillStruct(&data, smapping.MapFields(log_audit)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		log_audits = append(log_audits, data)
	}

	return log_audits, total, nil
}

func (svc *service) FindByID(log_auditID uint64) (*dtos.ResLog_audit, error) {
	res := dtos.ResLog_audit{}
	log_audit, err := svc.model.SelectByID(log_auditID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if log_audit == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(log_audit))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newLog_audit dtos.InputLog_audit) error {
	log_audit := log_audit.Log_audit{}

	err := smapping.FillStruct(&log_audit, smapping.MapFields(newLog_audit))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	log_audit.ID = helpers.GenerateID()
	err = svc.model.Insert(log_audit)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(log_auditData dtos.InputLog_audit, log_auditID uint64) error {
	newLog_audit := log_audit.Log_audit{}

	err := smapping.FillStruct(&newLog_audit, smapping.MapFields(log_auditData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newLog_audit.ID = log_auditID
	err = svc.model.Update(newLog_audit)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(log_auditID uint64) error {
	err := svc.model.DeleteByID(log_auditID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
