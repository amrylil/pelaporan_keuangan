package usecase

import (
	"errors"
	"pelaporan_keuangan/features/transaksi"
	"pelaporan_keuangan/features/transaksi/dtos"
	"pelaporan_keuangan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model transaksi.Repository
}

func New(model transaksi.Repository) transaksi.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResTransaksi, int64, error) {
	var transaksis []dtos.ResTransaksi

	transaksisEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, transaksi := range transaksisEnt {
		var data dtos.ResTransaksi

		if err := smapping.FillStruct(&data, smapping.MapFields(transaksi)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		transaksis = append(transaksis, data)
	}

	return transaksis, total, nil
}

func (svc *service) FindByID(transaksiID uint) (*dtos.ResTransaksi, error) {
	res := dtos.ResTransaksi{}
	transaksi, err := svc.model.SelectByID(transaksiID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if transaksi == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(transaksi))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newTransaksi dtos.InputTransaksi) error {
	transaksiEntity := transaksi.Transaksi{}

	err := smapping.FillStruct(&transaksiEntity, smapping.MapFields(newTransaksi))
	if err != nil {
		log.Error("Error mapping input to entity: ", err)
		return errors.New("invalid input data")
	}

	transaksiEntity.ID = helpers.GenerateID()
	err = svc.model.Insert(transaksiEntity)

	if err != nil {
		log.Error("Error creating transaction: ", err)
		return errors.New("failed to create transaction")
	}

	return nil
}

func (svc *service) Modify(transaksiData dtos.InputTransaksi, transaksiID uint) error {
	newTransaksi := transaksi.Transaksi{}

	err := smapping.FillStruct(&newTransaksi, smapping.MapFields(transaksiData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newTransaksi.ID = transaksiID
	err = svc.model.Update(newTransaksi)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(transaksiID uint) error {
	err := svc.model.DeleteByID(transaksiID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) ModifyStatus(transaksiID uint, statusID int) error {
	if err := svc.model.UpdateStatus(transaksiID, statusID); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (svc *service) ModifyPartial(updateData dtos.UpdateTransaksiRequest) error {
	if updateData.ID == nil {
		return errors.New("transaction ID is required")
	}

	// Check if transaction exists
	existing, err := svc.model.SelectByID(*updateData.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("transaction not found")
	}

	// Build updates map only for non-nil fields
	updates := make(map[string]interface{})

	if updateData.Tanggal != nil {
		updates["tanggal"] = *updateData.Tanggal
	}
	if updateData.IDTipeTransaksi != nil {
		updates["id_tipe_transaksi"] = *updateData.IDTipeTransaksi
	}
	if updateData.Jumlah != nil {
		updates["jumlah"] = *updateData.Jumlah
	}
	if updateData.Keterangan != nil {
		updates["keterangan"] = *updateData.Keterangan
	}
	if updateData.BuktiTransaksi != nil {
		updates["bukti_transaksi"] = *updateData.BuktiTransaksi
	}
	if updateData.IDStatusTransaksi != nil {
		updates["id_status_transaksi"] = *updateData.IDStatusTransaksi
	}
	if updateData.KomentarManajer != nil {
		updates["komentar_manajer"] = *updateData.KomentarManajer
	}
	if updateData.IDKategori != nil {
		updates["id_kategori"] = *updateData.IDKategori
	}
	if updateData.IDUser != nil {
		updates["id_user"] = *updateData.IDUser
	}
	if updateData.IDJenisPembayaran != nil {
		updates["id_jenis_pembayaran"] = *updateData.IDJenisPembayaran
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return svc.model.UpdatePartial(*updateData.ID, updates)
}

// FindWithFilter - New method for filtered search
func (svc *service) FindWithFilter(filter dtos.TransaksiListRequest) ([]dtos.ResTransaksi, int64, error) {
	var transaksis []dtos.ResTransaksi

	// Set defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	transaksisEnt, total, err := svc.model.GetWithFilter(filter)
	if err != nil {
		log.Error("Error fetching filtered transactions: ", err)
		return nil, 0, err
	}

	for _, transaksi := range transaksisEnt {
		var data dtos.ResTransaksi

		if err := smapping.FillStruct(&data, smapping.MapFields(transaksi)); err != nil {
			log.Error("Error mapping entity to response: ", err)
			return nil, 0, err
		}

		transaksis = append(transaksis, data)
	}

	return transaksis, total, nil
}
