package handler

import (
	"blueprint_golang/features/_blueprint"
	"blueprint_golang/features/_blueprint/dtos"
	"blueprint_golang/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service _blueprint.Usecase
}

func New(service _blueprint.Usecase) _blueprint.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetPlaceholders(c *gin.Context) {
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

	placeholders, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if placeholders == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Placeholders!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     "true",
		Message:    "Get All Placeholders Success",
		Data:       placeholders,
		Pagination: pagination,
	})
}

func (ctl *controller) PlaceholderDetails(c *gin.Context) {
	placeholderID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	placeholder, err := ctl.service.FindByID(placeholderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if placeholder == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Placeholder Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    placeholder,
		Status:  "true",
		Message: " Get Placeholder Detail Success",
	})
}

func (ctl *controller) CreatePlaceholder(c *gin.Context) {
	var input dtos.InputPlaceholder

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
		Message: " Create Placeholder Success",
		Status:  "true",
	})
}

func (ctl *controller) UpdatePlaceholder(c *gin.Context) {
	var input dtos.InputPlaceholder
	placeholderID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(errParam.Error()))
		return
	}

	placeholder, err := ctl.service.FindByID(placeholderID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if placeholder == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Placeholder Not Found!"))
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

	err = ctl.service.Modify(input, placeholderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update Placeholder Success",
		Status:  "true",
	})
}

func (ctl *controller) DeletePlaceholder(c *gin.Context) {
	placeholderID, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	placeholder, err := ctl.service.FindByID(placeholderID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if placeholder == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Placeholder Not Found!"))
		return
	}

	err = ctl.service.Remove(placeholderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Placeholder Success",
		Status:  "true",
	})
}
