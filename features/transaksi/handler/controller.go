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

// GetTransaksi godoc
// @Summary Get all transactions
// @Description Get all transactions with pagination
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(5)
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResTransaksi,pagination=helpers.Pagination} "Get all transactions success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No transactions found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi [get]
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

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Transaksis Success",
		Data:       transaksis,
		Pagination: paginationData,
	})
}

// TransaksiDetails godoc
// @Summary Get transaction details
// @Description Get transaction details by ID
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtos.ResTransaksi} "Get transaction detail success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid transaction ID"
// @Failure 404 {object} helpers.ResponseError "Transaction not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi/{id} [get]
func (ctl *controller) TransaksiDetails(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	transaksi, err := ctl.service.FindByID(transaksiID)
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

// CreateTransaksi godoc
// @Summary Create new transaction
// @Description Create a new financial transaction
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param request body dtos.InputTransaksi true "Transaction data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Create transaction success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid input data or validation error"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi [post]
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

// UpdateTransaksi godoc
// @Summary Update transaction
// @Description Update an existing transaction with partial data
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param request body dtos.UpdateTransaksiRequest true "Update transaction data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update transaction success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid transaction ID or request data"
// @Failure 404 {object} helpers.ResponseError "Transaction not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi/{id} [put]
func (ctl *controller) UpdateTransaksi(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid transaction ID"))
		return
	}

	// Check if transaction exists
	existingTransaksi, err := ctl.service.FindByID(transaksiID)
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
	id := transaksiID
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

// DeleteTransaksi godoc
// @Summary Delete transaction
// @Description Delete a specific transaction by ID
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} helpers.ResponseCUDSuccess "Delete transaction success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid transaction ID"
// @Failure 404 {object} helpers.ResponseError "Transaction not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi/{id} [delete]
func (ctl *controller) DeleteTransaksi(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	transaksi, err := ctl.service.FindByID(transaksiID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if transaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi Not Found!"))
		return
	}

	err = ctl.service.Remove(transaksiID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Transaksi Success",
		Status:  true,
	})
}

// UpdateTransaksiStatus godoc
// @Summary Update transaction status
// @Description Update the status of a specific transaction
// @Tags Transaksi
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param request body object{status_id=int} true "Status update data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update transaction status success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid transaction ID or missing status_id"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi/{id}/status [patch]
func (ctl *controller) UpdateTransaksiStatus(c *gin.Context) {
	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid transaction ID"))
		return
	}

	var input struct {
		StatusID uint64 `json:"status_id" validate:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request data"))
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Status ID is required"))
		return
	}

	err = ctl.service.ModifyStatus(transaksiID, int(input.StatusID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Update Transaction Status Success",
		Status:  true,
	})
}
