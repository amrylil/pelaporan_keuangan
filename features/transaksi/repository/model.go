package repository

import (
	"pelaporan_keuangan/features/transaksi"
	"pelaporan_keuangan/features/transaksi/dtos"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaksi.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]transaksi.Transaksi, int64, error) {
	var transaksis []transaksi.Transaksi
	var total int64

	// Count total records
	if err := mdl.db.Model(&transaksi.Transaksi{}).Count(&total).Error; err != nil {
		log.Error("Error counting transactions: ", err)
		return nil, 0, err
	}

	offset := (page - 1) * size

	// Preload relationships for better performance
	err := mdl.db.Preload("TipeTransaksi").
		Preload("StatusTransaksi").
		Preload("Kategori").
		Preload("User").
		Preload("JenisPembayaran").
		Offset(offset).
		Limit(size).
		Order("created_at DESC"). // Order by newest first
		Find(&transaksis).Error

	if err != nil {
		log.Error("Error fetching transactions: ", err)
		return nil, 0, err
	}

	return transaksis, total, nil
}

func (mdl *model) Insert(newTransaksi transaksi.Transaksi) error {
	err := mdl.db.Create(&newTransaksi).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(transaksiID uint) (*transaksi.Transaksi, error) {
	var transaksi transaksi.Transaksi
	err := mdl.db.Preload("TipeTransaksi").
		Preload("StatusTransaksi").
		Preload("Kategori").
		Preload("User").
		Preload("JenisPembayaran").
		First(&transaksi, transaksiID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil instead of error for not found
		}
		log.Error("Error fetching transaction: ", err)
		return nil, err
	}

	return &transaksi, nil
}

func (mdl *model) UpdatePartial(transaksiID uint, updates map[string]interface{}) error {
	err := mdl.db.Model(&transaksi.Transaksi{}).
		Where("id = ?", transaksiID).
		Updates(updates).Error

	if err != nil {
		log.Error("Error updating transaction: ", err)
		return err
	}

	return nil
}

func (mdl *model) GetWithFilter(filter dtos.TransaksiListRequest) ([]transaksi.Transaksi, int64, error) {
	var transaksis []transaksi.Transaksi
	var total int64

	query := mdl.db.Model(&transaksi.Transaksi{})

	// Apply filters
	if filter.Search != "" {
		query = query.Where("keterangan ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.IDTipeTransaksi != nil {
		query = query.Where("id_tipe_transaksi = ?", *filter.IDTipeTransaksi)
	}
	if filter.IDStatusTransaksi != nil {
		query = query.Where("id_status_transaksi = ?", *filter.IDStatusTransaksi)
	}
	if filter.IDKategori != nil {
		query = query.Where("id_kategori = ?", *filter.IDKategori)
	}
	if filter.IDUser != nil {
		query = query.Where("id_user = ?", *filter.IDUser)
	}
	if filter.TanggalMulai != nil && filter.TanggalAkhir != nil {
		query = query.Where("tanggal BETWEEN ? AND ?", *filter.TanggalMulai, *filter.TanggalAkhir)
	}

	// Count with filters
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (filter.Page - 1) * filter.Limit
	err := query.Preload("TipeTransaksi").
		Preload("StatusTransaksi").
		Preload("Kategori").
		Preload("User").
		Preload("JenisPembayaran").
		Offset(offset).
		Limit(filter.Limit).
		Order("created_at DESC").
		Find(&transaksis).Error

	if err != nil {
		log.Error("Error fetching filtered transactions: ", err)
		return nil, 0, err
	}

	return transaksis, total, nil
}

func (mdl *model) Update(transaksi transaksi.Transaksi) error {
	// Use Select to update all fields including zero values if needed
	err := mdl.db.Model(&transaksi).
		Where("id = ?", transaksi.ID).
		Updates(&transaksi).Error

	if err != nil {
		log.Error("Error updating transaction: ", err)
		return err
	}

	return nil
}

func (mdl *model) DeleteByID(transaksiID uint) error {
	err := mdl.db.Delete(&transaksi.Transaksi{}, transaksiID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) UpdateStatus(transaksiID uint, statusID int) error {
	err := mdl.db.Model(&transaksi.Transaksi{}).
		Where("id = ?", transaksiID).
		Update("id_status_transaksi", statusID).Error
	return err
}
