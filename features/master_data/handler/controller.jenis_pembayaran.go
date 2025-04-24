package handler

import (
	"net/http"
	"pelaporan_keuangan/features/master_data/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ctl *controller) GetJenisPembayaran(c *gin.Context) {
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

func (ctl *controller) JenisPembayaranDetails(c *gin.Context) {
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(uint(JenisPembayaranID))
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

func (ctl *controller) UpdateJenisPembayaran(c *gin.Context) {
	var input dtos.InputJenisPembayaran
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(uint(JenisPembayaranID))
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

	err = ctl.service.ModifyJenisPembayaran(input, uint(JenisPembayaranID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update JenisPembayaran Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteJenisPembayaran(c *gin.Context) {
	JenisPembayaranID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	JenisPembayaran, err := ctl.service.FindJenisPembayaranByID(uint(JenisPembayaranID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if JenisPembayaran == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("JenisPembayaran Not Found!"))
		return
	}

	err = ctl.service.RemoveJenisPembayaran(uint(JenisPembayaranID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete JenisPembayaran Success",
		Status:  true,
	})
}
