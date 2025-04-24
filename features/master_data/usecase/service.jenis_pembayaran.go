package usecase

import (
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

func (svc *service) FindAllJenisPembayaran(page, size int) ([]dtos.ResJenisPembayaran, int64, error) {
	var jenis_pembayarans []dtos.ResJenisPembayaran

	jenisPembayaranData, total, err := svc.model.GetAllJenisPembayaran(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, jenis_pembayaran := range jenisPembayaranData {
		var data dtos.ResJenisPembayaran

		if err := smapping.FillStruct(&data, smapping.MapFields(jenis_pembayaran)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		jenis_pembayarans = append(jenis_pembayarans, data)
	}

	return jenis_pembayarans, total, nil
}

func (svc *service) FindJenisPembayaranByID(jenisPembayaranID uint) (*dtos.ResJenisPembayaran, error) {
	res := dtos.ResJenisPembayaran{}
	jenis_pembayaran, err := svc.model.SelectJenisPembayaranByID(jenisPembayaranID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if jenis_pembayaran == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(jenis_pembayaran))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) CreateJenisPembayaran(newJenisPembayaran dtos.InputJenisPembayaran) error {
	jenis_pembayaran := master_data.JenisPembayaran{}

	err := smapping.FillStruct(&jenis_pembayaran, smapping.MapFields(newJenisPembayaran))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	jenis_pembayaran.ID = helpers.GenerateID()
	err = svc.model.InsertJenisPembayaran(jenis_pembayaran)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) ModifyJenisPembayaran(jenisPembayaran dtos.InputJenisPembayaran, jenisPembayaranID uint) error {
	newJenisPembayaran := master_data.JenisPembayaran{}

	err := smapping.FillStruct(&newJenisPembayaran, smapping.MapFields(jenisPembayaran))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newJenisPembayaran.ID = jenisPembayaranID
	err = svc.model.UpdateJenisPembayaran(newJenisPembayaran)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) RemoveJenisPembayaran(jenisPembayaranID uint) error {
	err := svc.model.DeleteJenisPembayaranByID(jenisPembayaranID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
