package usecase

import (
	"pelaporan_keuangan/features/kategori"
	"pelaporan_keuangan/features/kategori/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model kategori.Repository
}

func New(model kategori.Repository) kategori.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResKategori, int64, error) {
	var kategoris []dtos.ResKategori

	kategorisEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, kategori := range kategorisEnt {
		var data dtos.ResKategori

		if err := smapping.FillStruct(&data, smapping.MapFields(kategori)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		kategoris = append(kategoris, data)
	}

	return kategoris, total, nil
}

func (svc *service) FindByID(kategoriID uint64) (*dtos.ResKategori, error) {
	res := dtos.ResKategori{}
	kategori, err := svc.model.SelectByID(kategoriID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if kategori == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(kategori))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newKategori dtos.InputKategori) error {
	kategori := kategori.Kategori{}

	err := smapping.FillStruct(&kategori, smapping.MapFields(newKategori))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	kategori.ID = helpers.GenerateID()
	err = svc.model.Insert(kategori)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(kategoriData dtos.InputKategori, kategoriID uint64) error {
	newKategori := kategori.Kategori{}

	err := smapping.FillStruct(&newKategori, smapping.MapFields(kategoriData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newKategori.ID = kategoriID
	err = svc.model.Update(newKategori)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(kategoriID uint64) error {
	err := svc.model.DeleteByID(kategoriID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
