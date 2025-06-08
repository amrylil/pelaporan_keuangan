package handler

import (
	"net/http"
	"pelaporan_keuangan/features/transaksi"
	"pelaporan_keuangan/features/transaksi/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service transaksi.Usecase
}

func New(service transaksi.Usecase) transaksi.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetTransaksi(c *gin.Context) {
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

	transaksis, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if transaksis == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Transaksis!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Transaksis Success",
		Data:       transaksis,
		Pagination: paginationData,
	})
}

func (ctl *controller) TransaksiDetails(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	transaksi, err := ctl.service.FindByID(uint(transaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if transaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    transaksi,
		Status:  true,
		Message: " Get Transaksi Detail Success",
	})
}

func (ctl *controller) CreateTransaksi(c *gin.Context) {
	var input dtos.InputTransaksi

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
		Message: " Create Transaksi Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateTransaksi(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid transaction ID"))
		return
	}

	// Check if transaction exists
	existingTransaksi, err := ctl.service.FindByID(uint(transaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}
	if existingTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaction not found"))
		return
	}

	// Use UpdateTransaksiRequest for partial updates
	var input dtos.UpdateTransaksiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request data"))
		return
	}

	// Set the ID for update
	id := uint(transaksiID)
	input.ID = &id

	// Validate the input
	if err := validate.Struct(input); err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Validation failed", gin.H{
			"errors": errMap,
		}))
		return
	}

	// Call service with proper update DTO
	err = ctl.service.ModifyPartial(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update Transaction Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteTransaksi(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	transaksi, err := ctl.service.FindByID(uint(transaksiID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if transaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi Not Found!"))
		return
	}

	err = ctl.service.Remove(uint(transaksiID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Transaksi Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateTransaksiStatus(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid transaction ID"))
		return
	}

	var input struct {
		StatusID uint `json:"status_id" validate:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request data"))
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Status ID is required"))
		return
	}

	err = ctl.service.ModifyStatus(uint(transaksiID), int(input.StatusID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update Transaction Status Success",
		Status:  true,
	})
}
