package handler

import (
	"log"
	"net/http"
	"pelaporan_keuangan/features/kategori"
	"pelaporan_keuangan/features/kategori/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service kategori.Usecase
}

func New(service kategori.Usecase) kategori.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetKategori(c *gin.Context) {
	var pagination dtos.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Please provide valid pagination data!"))
		return
	}

	if pagination.Page < 1 || pagination.Size < 1 {
		pagination.Page = 1
		pagination.Size = 5
	}
	page := pagination.Page
	pageSize := pagination.Size

	kategoris, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Kategoris Success",
		Data:       kategoris,
		Pagination: paginationData,
	})
}

func (ctl *controller) KategoriDetails(c *gin.Context) {
	kategoriID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	kategori, err := ctl.service.FindByID(kategoriID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if kategori == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Kategori Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    kategori,
		Status:  true,
		Message: " Get Kategori Detail Success",
	})
}

func (ctl *controller) CreateKategori(c *gin.Context) {
	var input dtos.InputKategori

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"))
		return
	}

	validate = validator.New()

	err := validate.Struct(input)

	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Bad Request!", gin.H{
			"error": errMap,
		}))
		return
	}

	err = ctl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create Kategori Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateKategori(c *gin.Context) {
	var input dtos.InputKategori

	log.Print("id sebelum parse : ", c.Param("id"))
	kategoriID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("ID kategori tidak valid"))
		return
	}
	log.Print("setelah : ", kategoriID)
	_, err = ctl.service.FindByID(kategoriID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Kategori dengan ID tersebut tidak ditemukan"))
		return
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Request body tidak valid"))
		return
	}
	validate := validator.New()
	err = validate.Struct(input)

	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Input tidak valid!", gin.H{
			"errors": errMap,
		}))
		return
	}

	err = ctl.service.Modify(input, kategoriID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Gagal mengupdate kategori"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update Kategori Success",
		Status:  true,
	})
}
func (ctl *controller) DeleteKategori(c *gin.Context) {
	kategoriID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	kategori, err := ctl.service.FindByID(kategoriID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if kategori == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Kategori Not Found!"))
		return
	}

	err = ctl.service.Remove(kategoriID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Kategori Success",
		Status:  true,
	})
}
