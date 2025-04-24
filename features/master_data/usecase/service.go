package usecase

import (
	"pelaporan_keuangan/features/master_data"
)

type service struct {
	model master_data.Repository
}

func New(model master_data.Repository) master_data.Usecase {
	return &service{
		model: model,
	}
}
