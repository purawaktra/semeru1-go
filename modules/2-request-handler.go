package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru1-go/dto"
	"github.com/purawaktra/semeru1-go/functions"
	"github.com/purawaktra/semeru1-go/utils"
	"net/http"
	"strconv"
	"time"
)

type Semeru1RequestHandler struct {
	ctrl Semeru1Controller
}

func CreateSemeru1RequestHandler(ctrl Semeru1Controller) Semeru1RequestHandler {
	return Semeru1RequestHandler{
		ctrl: ctrl,
	}
}

func (rh *Semeru1RequestHandler) SelectCityById(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectCityById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectCityById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectCityById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city and check err
	utils.Debug("SelectCityById", header)
	data, err, code := rh.ctrl.SelectCityById(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectCityById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh *Semeru1RequestHandler) SelectCityByName(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectCityByName", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectCityByName", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectCityByName", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city and check err
	utils.Debug("SelectCityByName", header)
	data, err, code := rh.ctrl.SelectCityByName(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectCityByName", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh *Semeru1RequestHandler) SelectAllCity(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectAllCity", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectAllCity", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectAllCity", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the city and check err
	utils.Debug("SelectAllCity", header)
	data, err, code := rh.ctrl.SelectAllCity(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectAllCity", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh *Semeru1RequestHandler) SelectProvinceById(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectProvinceById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectProvinceById", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectProvinceById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the province and check err
	utils.Debug("SelectProvinceById", header)
	data, err, code := rh.ctrl.SelectProvinceById(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectProvinceById", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}

func (rh *Semeru1RequestHandler) SelectAllProvince(c *gin.Context) {
	utils.Info("=== New Request ===")
	// start timer
	start := time.Now()

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "SelectAllProvince", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FH",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if !functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "SelectAllProvince", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse request body to struct and check err
	bodyBytes, err := c.GetRawData()
	if err != nil {
		utils.Error(err, "SelectAllProvince", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			"FB",
			err.Error())
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the province and check err
	utils.Debug("SelectAllProvince", header)
	data, err, code := rh.ctrl.SelectAllProvince(bodyBytes)
	if err != nil {
		utils.Error(err, "SelectAllProvince", "")
		end := time.Now()
		functions.SetBaseHeader(c,
			header.RequestId,
			strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			code,
			err.Error())
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	end := time.Now()
	functions.SetBaseHeader(c,
		header.RequestId,
		strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		code,
		"Success")
	c.JSON(http.StatusOK, data)
	return
}
