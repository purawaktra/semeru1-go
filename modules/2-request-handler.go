package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru1-go/dto"
	"github.com/purawaktra/semeru1-go/functions"
	"github.com/purawaktra/semeru1-go/utils"
	"net/http"
)

type Semeru1RequestHandler struct {
	ctrl Semeru1Controller
}

func CreateSemeru1RequestHandler(ctrl Semeru1Controller) Semeru1RequestHandler {
	return Semeru1RequestHandler{
		ctrl: ctrl,
	}
}

type Semeru1RequestHandlerInterface interface {
	SelectCityById(c *gin.Context)
	SelectCityByName(c *gin.Context)
	SelectCityByProvince(c *gin.Context)
	SelectAllCity(c *gin.Context)
	SelectProvinceById(c *gin.Context)
	SelectAllProvince(c *gin.Context)
}

func (rh Semeru1RequestHandler) SelectCityById(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectCityById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectCityById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city
	utils.Debug("SelectCityById", requestBody)
	response, err := rh.ctrl.SelectCityById(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectCityById", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru1RequestHandler) SelectCityByName(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectCityByName", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectCityByName", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city
	utils.Debug("SelectCityByName", requestBody)
	response, err := rh.ctrl.SelectCityByName(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectCityByName", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru1RequestHandler) SelectCityByProvince(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectCityByProvince", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectCityByProvince", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city
	utils.Debug("SelectCityByProvince", requestBody)
	response, err := rh.ctrl.SelectCityByProvince(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectCityByProvince", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru1RequestHandler) SelectAllCity(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectAllCity", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectAllCity", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city
	utils.Debug("SelectAllCity", requestBody)
	response, err := rh.ctrl.SelectAllCity(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectAllCity", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru1RequestHandler) SelectProvinceById(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectProvinceById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectProvinceById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the province
	utils.Debug("SelectProvinceById", requestBody)
	response, err := rh.ctrl.SelectProvinceById(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectProvinceById", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}

func (rh Semeru1RequestHandler) SelectAllProvince(c *gin.Context) {
	utils.Info("=== New Request ===")

	// parse request body to struct
	var requestBody dto.Request
	err := c.Bind(&requestBody)

	// check for error on parse request body
	if err != nil {
		utils.Error(err, "SelectAllProvince", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(requestBody.RequestId) {
		utils.Error(err, "SelectAllProvince", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the province
	utils.Debug("SelectAllProvince", requestBody)
	response, err := rh.ctrl.SelectAllProvince(requestBody)

	// check for error on call controller
	if err != nil {
		utils.Error(err, "SelectAllProvince", "")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}
