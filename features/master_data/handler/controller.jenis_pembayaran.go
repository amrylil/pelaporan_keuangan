package handler

import (
	"net/http"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetJenisPembayaran godoc
// @Summary Get all payment types
// @Description Get all payment types with pagination
// @Tags MasterData-JenisPembayaran
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(5)
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResJenisPembayaran,pagination=helpers.Pagination} "Get all payment types success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No payment types found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /jenis-pembayaran [get]
func (ctl *controller) GetJenisPembayaran(c *gin.Context) {
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

	JenisPembayarans, total, err := ctl.service.FindAllJenisPembayaran(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if JenisPembayarans == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No JenisPembayarans!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All JenisPembayarans Success",
		Data:       JenisPembayarans,
		Pagination: paginationData,
	})
}

// JenisPembayaranDetails godoc
// @Summary Get payment type details
// @Description Get payment type details by ID
// @Tags MasterData-JenisPembayaran
// @Accept json
// @Produce json
// @Param id path int true "Payment Type ID"
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtos.ResJenisPembayaran} "Get payment type detail success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid payment type ID"
// @Failure 404 {object} helpers.ResponseError "Payment type not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /jenis-pembayaran/{id} [get]
func (ctl *controller) JenisPembayaranDetails(c *gin.Context) {
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(JenisPembayaranID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if JenisPembayaran == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("JenisPembayaran Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    JenisPembayaran,
		Status:  true,
		Message: " Get JenisPembayaran Detail Success",
	})
}

// CreateJenisPembayaran godoc
// @Summary Create new payment type
// @Description Create a new payment type
// @Tags MasterData-JenisPembayaran
// @Accept json
// @Produce json
// @Param request body dtos.InputJenisPembayaran true "Payment type data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Create payment type success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid input data or validation error"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /jenis-pembayaran [post]
func (ctl *controller) CreateJenisPembayaran(c *gin.Context) {
	var input dtos.InputJenisPembayaran

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

	err = ctl.service.CreateJenisPembayaran(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create JenisPembayaran Success",
		Status:  true,
	})
}

// UpdateJenisPembayaran godoc
// @Summary Update payment type
// @Description Update an existing payment type
// @Tags MasterData-JenisPembayaran
// @Accept json
// @Produce json
// @Param id path int true "Payment Type ID"
// @Param request body dtos.InputJenisPembayaran true "Update payment type data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update payment type success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid payment type ID or request data"
// @Failure 404 {object} helpers.ResponseError "Payment type not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /jenis-pembayaran/{id} [put]
func (ctl *controller) UpdateJenisPembayaran(c *gin.Context) {
	var input dtos.InputJenisPembayaran
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(JenisPembayaranID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if JenisPembayaran == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("JenisPembayaran Not Found!"))
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

	err = ctl.service.ModifyJenisPembayaran(input, JenisPembayaranID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update JenisPembayaran Success",
		Status:  true,
	})
}

// DeleteJenisPembayaran godoc
// @Summary Delete payment type
// @Description Delete a specific payment type by ID
// @Tags MasterData-JenisPembayaran
// @Accept json
// @Produce json
// @Param id path int true "Payment Type ID"
// @Success 200 {object} helpers.ResponseCUDSuccess "Delete payment type success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid payment type ID"
// @Failure 404 {object} helpers.ResponseError "Payment type not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /jenis-pembayaran/{id} [delete]
func (ctl *controller) DeleteJenisPembayaran(c *gin.Context) {
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(JenisPembayaranID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if JenisPembayaran == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("JenisPembayaran Not Found!"))
		return
	}

	err = ctl.service.RemoveJenisPembayaran(JenisPembayaranID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete JenisPembayaran Success",
		Status:  true,
	})
}
