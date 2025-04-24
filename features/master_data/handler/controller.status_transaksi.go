package handler

import (
	"net/http"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ctl *controller) GetStatusTransaksi(c *gin.Context) {
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

	statusTransaksis, total, err := ctl.service.FindAllStatusTransaksi(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if statusTransaksis == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No StatusTransaksis!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All StatusTransaksis Success",
		Data:       statusTransaksis,
		Pagination: paginationData,
	})
}

func (ctl *controller) StatusTransaksiDetails(c *gin.Context) {
	statusTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	statusTransaksi, err := ctl.service.FindStatusTransaksiByID(uint(statusTransaksiID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if statusTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("StatusTransaksi Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    statusTransaksi,
		Status:  true,
		Message: " Get StatusTransaksi Detail Success",
	})
}

func (ctl *controller) CreateStatusTransaksi(c *gin.Context) {
	var input dtos.InputStatusTransaksi

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

	err = ctl.service.CreateStatusTransaksi(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create StatusTransaksi Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateStatusTransaksi(c *gin.Context) {
	var input dtos.InputStatusTransaksi
	statusTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	statusTransaksi, err := ctl.service.FindStatusTransaksiByID(uint(statusTransaksiID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if statusTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("StatusTransaksi Not Found!"))
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

	err = ctl.service.ModifyStatusTransaksi(input, uint(statusTransaksiID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update StatusTransaksi Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteStatusTransaksi(c *gin.Context) {
	statusTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	statusTransaksi, err := ctl.service.FindStatusTransaksiByID(uint(statusTransaksiID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if statusTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("StatusTransaksi Not Found!"))
		return
	}

	err = ctl.service.RemoveStatusTransaksi(uint(statusTransaksiID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete StatusTransaksi Success",
		Status:  true,
	})
}
