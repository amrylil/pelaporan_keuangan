package handler

import (
	"net/http"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ctl *controller) GetTipeTransaksi(c *gin.Context) {
	var pagination dtos.Pagination
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Please provide valid pagination data!"))
		return
	}

	if pagination.Page < 1 || pagination.Size < 1 {
		pagination.Page = 1
		pagination.Size = 5
	}
	page := pagination.Page
	pageSize := pagination.Size

	tipeTransaksis, total, err := ctl.service.FindAllTipeTransaksi(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if tipeTransaksis == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is no TipeTransaksi data!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All TipeTransaksi Success",
		Data:       tipeTransaksis,
		Pagination: paginationData,
	})
}

func (ctl *controller) TipeTransaksiDetails(c *gin.Context) {
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(uint(tipeTransaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if tipeTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("TipeTransaksi Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    tipeTransaksi,
		Status:  true,
		Message: "Get TipeTransaksi Detail Success",
	})
}

func (ctl *controller) CreateTipeTransaksi(c *gin.Context) {
	var input dtos.InputTipeTransaksi

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

	err = ctl.service.CreateTipeTransaksi(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Create TipeTransaksi Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateTipeTransaksi(c *gin.Context) {
	var input dtos.InputTipeTransaksi
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(uint(tipeTransaksiID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if tipeTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("TipeTransaksi Not Found!"))
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"))
		return
	}

	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Bad Request!", gin.H{
			"error": errMap,
		}))
		return
	}

	err = ctl.service.ModifyTipeTransaksi(input, uint(tipeTransaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update TipeTransaksi Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteTipeTransaksi(c *gin.Context) {
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(uint(tipeTransaksiID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if tipeTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("TipeTransaksi Not Found!"))
		return
	}

	err = ctl.service.RemoveTipeTransaksi(uint(tipeTransaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Delete TipeTransaksi Success",
		Status:  true,
	})
}
