package handler

import (
	"net/http"
	"pelaporan_keuangan/features/log_audit"
	"pelaporan_keuangan/features/log_audit/dtos"
	"pelaporan_keuangan/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service log_audit.Usecase
}

func New(service log_audit.Usecase) log_audit.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetLog_audit(c *gin.Context) {
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

	log_audits, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if log_audits == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Log_audits!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Log_audits Success",
		Data:       log_audits,
		Pagination: paginationData,
	})
}

func (ctl *controller) Log_auditDetails(c *gin.Context) {
	log_auditID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	log_audit, err := ctl.service.FindByID(log_auditID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if log_audit == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Log_audit Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    log_audit,
		Status:  true,
		Message: " Get Log_audit Detail Success",
	})
}

func (ctl *controller) CreateLog_audit(c *gin.Context) {
	var input dtos.InputLog_audit

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
		Message: " Create Log_audit Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateLog_audit(c *gin.Context) {
	var input dtos.InputLog_audit
	log_auditID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	log_audit, err := ctl.service.FindByID(log_auditID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if log_audit == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Log_audit Not Found!"))
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

	err = ctl.service.Modify(input, log_auditID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update Log_audit Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteLog_audit(c *gin.Context) {
	log_auditID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	log_audit, err := ctl.service.FindByID(log_auditID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if log_audit == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Log_audit Not Found!"))
		return
	}

	err = ctl.service.Remove(log_auditID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Log_audit Success",
		Status:  true,
	})
}
