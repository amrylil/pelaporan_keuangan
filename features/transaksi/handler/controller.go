package handler

import (
	"log"
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

var validate = validator.New()

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
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi [get]
func (ctl *controller) GetTransaksi(c *gin.Context) {
	var pagination dtos.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Mohon sediakan data paginasi yang valid!"))
		return
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.Size < 1 {
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
		Message:    "Sukses Mendapatkan Semua Transaksi",
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
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("ID transaksi tidak valid!"))
		return
	}

	transaksi, err := ctl.service.FindByID(transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if transaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi Tidak Ditemukan!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    transaksi,
		Status:  true,
		Message: "Sukses Mendapatkan Detail Transaksi",
	})
}

// CreateTransaksi godoc
// @Summary Create new transaction with optional proof
// @Description Create a new financial transaction. The request should be multipart/form-data.
// @Tags Transaksi
// @Accept multipart/form-data
// @Produce json
// @Param judul formData string true "Judul Transaksi"
// @Param deskripsi formData string false "Deskripsi Transaksi"
// @Param jumlah formData number true "Jumlah Transaksi"
// @Param tipe formData string true "Tipe Transaksi ('pemasukan' atau 'pengeluaran')"
// @Param bukti_transaksi formData file false "Bukti Transaksi (Gambar/PDF)"
// @Success 201 {object} helpers.ResponseCUDSuccess "Create transaction success"
// @Failure 400 {object} helpers.ResponseError "Bad request - Invalid input data or validation error"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /transaksi [post]
func (ctl *controller) CreateTransaksi(c *gin.Context) {
	var input dtos.InputTransaksi

	if err := c.ShouldBind(&input); err != nil {
		log.Printf("Error binding form data: %v", err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Data formulir tidak valid!"))
		return
	}

	userIDFromToken, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, helpers.BuildErrorResponse("Sesi pengguna tidak valid atau tidak ditemukan."))
		return
	}

	var idUserUint64 uint64
	switch v := userIDFromToken.(type) {
	case float64:
		idUserUint64 = uint64(v)
	case string:
		parsedID, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Format userID di token tidak valid."))
			return
		}
		idUserUint64 = parsedID
	case uint:
		idUserUint64 = uint64(v)
	case uint64:
		idUserUint64 = v
	default:
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Tipe userID di context tidak dikenal."))
		return
	}

	input.IDUser = idUserUint64

	file, err := c.FormFile("bukti_transaksi")
	if err == nil {
		uploadResult, uploadErr := helpers.UploadFile(file, "transaksi")
		if uploadErr != nil {
			log.Printf("Cloudinary upload error: %v", uploadErr)
			c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Gagal mengunggah file bukti."))
			return
		}
		input.BuktiTransaksi = uploadResult.SecureURL
	} else if err != http.ErrMissingFile {
		log.Printf("Error getting form file: %v", err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Error saat memproses file."))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Data tidak valid!", gin.H{
			"errors": errMap,
		}))
		return
	}

	// 6. Panggil service untuk membuat transaksi
	err = ctl.service.Create(input)
	if err != nil {
		log.Printf("Service create error: %v", err)
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helpers.ResponseCUDSuccess{
		Message: "Sukses Membuat Transaksi",
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
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("ID transaksi tidak valid"))
		return
	}

	// Cek apakah transaksi ada
	existingTransaksi, err := ctl.service.FindByID(transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}
	if existingTransaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi tidak ditemukan"))
		return
	}

	// Gunakan UpdateTransaksiRequest untuk pembaruan parsial
	var input dtos.UpdateTransaksiRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Data permintaan tidak valid"))
		return
	}

	// Set ID untuk pembaruan
	id := transaksiID
	input.ID = &id

	// Validasi input
	if err := validate.Struct(input); err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Validasi gagal", gin.H{
			"errors": errMap,
		}))
		return
	}

	err = ctl.service.ModifyPartial(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Sukses Memperbarui Transaksi",
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
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("ID transaksi tidak valid!"))
		return
	}

	// Cek apakah transaksi ada sebelum menghapus
	transaksi, err := ctl.service.FindByID(transaksiID)
	if err != nil {
		// Sebaiknya tidak menampilkan error internal ke user
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Terjadi kesalahan saat mencari transaksi"))
		return
	}

	if transaksi == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Transaksi Tidak Ditemukan!"))
		return
	}

	err = ctl.service.Remove(transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Sukses Menghapus Transaksi",
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
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("ID transaksi tidak valid"))
		return
	}

	var input struct {
		StatusID uint `json:"status_id" validate:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Data permintaan tidak valid"))
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Status ID wajib diisi"))
		return
	}

	err = ctl.service.ModifyStatus(transaksiID, int(input.StatusID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Sukses Memperbarui Status Transaksi",
		Status:  true,
	})
}
