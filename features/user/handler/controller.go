package handler

import (
	"blueprint_golang/features/user"
	"blueprint_golang/features/user/dtos"
	"blueprint_golang/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service user.Usecase
}

func New(service user.Usecase) user.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetUsers(c *gin.Context) {
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

	users, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if users == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Users!"))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildErrorResponse("Success!", gin.H{
		"data": users,
	}))
}

func (ctl *controller) UserDetails(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	user, err := ctl.service.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildErrorResponse("Success!", gin.H{
		"data": user,
	}))
}

func (ctl *controller) CreateUser(c *gin.Context) {
	var input dtos.InputUser

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

	c.JSON(http.StatusOK, helpers.BuildErrorResponse("Success!", gin.H{
		"data": "succes",
	}))
}

func (ctl *controller) UpdateUser(c *gin.Context) {
	var input dtos.InputUser
	userID, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(errParam.Error()))
		return
	}

	user, err := ctl.service.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
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

	err = ctl.service.Modify(input, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildErrorResponse("User Successfully Updated!"))
}

func (ctl *controller) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	user, err := ctl.service.FindByID(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
		return
	}

	err = ctl.service.Remove(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildErrorResponse("User Successfully Deleted!"))
}
