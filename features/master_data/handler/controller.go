package handler

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/go-playground/validator/v10"
)

type controller struct {
	service master_data.Usecase
}

func New(service master_data.Usecase) master_data.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate
