package handler

import (
	"net/http"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetTipeTransaksi godoc
// @Summary Get all tipe transaksi with pagination
// @Description Retrieve all tipe transaksi data with pagination support
// @Tags MasterData-TipeTransaksi
// @Accept json
// @Produce json
// @Param pagination body dtos.Pagination true "Pagination parameters"
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResTipeTransaksi,pagination=helpers.Pagination} "Success get all tipe transaksi"
// @Failure 400 {object} helpers.ResponseError "Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No tipe transaksi found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /tipe-transaksi [post]
func (ctl *controller) GetTipeTransaksi(c *gin.Context) {
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

	tipeTransaksis, total, err := ctl.service.FindAllTipeTransaksi(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
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

// TipeTransaksiDetails godoc
// @Summary Get tipe transaksi by ID
// @Description Retrieve detailed information of a specific tipe transaksi by ID
// @Tags MasterData-TipeTransaksi
// @Accept json
// @Produce json
// @Param id path int true "Tipe Transaksi ID"
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtos.ResTipeTransaksi} "Success get tipe transaksi detail"
// @Failure 400 {object} helpers.ResponseError "Invalid ID parameter"
// @Failure 404 {object} helpers.ResponseError "Tipe transaksi not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /tipe-transaksi/{id} [get]
func (ctl *controller) TipeTransaksiDetails(c *gin.Context) {
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(tipeTransaksiID)
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

// CreateTipeTransaksi godoc
// @Summary Create new tipe transaksi
// @Description Create a new tipe transaksi with the provided data
// @Tags MasterData-TipeTransaksi
// @Accept json
// @Produce json
// @Param input body dtos.InputTipeTransaksi true "Tipe transaksi input data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Success create tipe transaksi"
// @Failure 400 {object} helpers.ResponseError "Invalid request data or validation error"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /tipe-transaksi [post]
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

// UpdateTipeTransaksi godoc
// @Summary Update existing tipe transaksi
// @Description Update an existing tipe transaksi with the provided data
// @Tags MasterData-TipeTransaksi
// @Accept json
// @Produce json
// @Param id path int true "Tipe Transaksi ID"
// @Param input body dtos.InputTipeTransaksi true "Tipe transaksi update data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Success update tipe transaksi"
// @Failure 400 {object} helpers.ResponseError "Invalid request data or validation error"
// @Failure 404 {object} helpers.ResponseError "Tipe transaksi not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /tipe-transaksi/{id} [put]
func (ctl *controller) UpdateTipeTransaksi(c *gin.Context) {
	var input dtos.InputTipeTransaksi
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(tipeTransaksiID)
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

	err = ctl.service.ModifyTipeTransaksi(input, tipeTransaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update TipeTransaksi Success",
		Status:  true,
	})
}

// DeleteTipeTransaksi godoc
// @Summary Delete tipe transaksi
// @Description Delete an existing tipe transaksi by ID
// @Tags MasterData-TipeTransaksi
// @Accept json
// @Produce json
// @Param id path int true "Tipe Transaksi ID"
// @Success 200 {object} helpers.ResponseCUDSuccess "Success delete tipe transaksi"
// @Failure 400 {object} helpers.ResponseError "Invalid ID parameter"
// @Failure 404 {object} helpers.ResponseError "Tipe transaksi not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /tipe-transaksi/{id} [delete]
func (ctl *controller) DeleteTipeTransaksi(c *gin.Context) {
	tipeTransaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	tipeTransaksi, err := ctl.service.FindTipeTransaksiByID(tipeTransaksiID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if tipeTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("TipeTransaksi Not Found!"))
		return
	}

	err = ctl.service.RemoveTipeTransaksi(tipeTransaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Delete TipeTransaksi Success",
		Status:  true,
	})
}
