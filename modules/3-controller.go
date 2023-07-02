package modules

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/purawaktra/semeru1-go/dto"
	"github.com/purawaktra/semeru1-go/utils"
	"strconv"
	"time"
)

type Semeru1Controller struct {
	uc Semeru1Usecase
}

func CreateSemeru1Controller(uc Semeru1Usecase) Semeru1Controller {
	return Semeru1Controller{
		uc: uc,
	}
}

type Semeru1ControllerInterface interface {
	SelectCityById(req dto.Request) (any, error)
	SelectCityByName(req dto.Request) (any, error)
	SelectCityByProvince(req dto.Request) (any, error)
	SelectAllCity(req dto.Request) (any, error)
	SelectProvinceById(req dto.Request) (any, error)
	SelectAllProvince(req dto.Request) (any, error)
}

func (ctrl Semeru1Controller) SelectCityById(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectCityById", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestCity
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectCityById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the city
	cities, err := ctrl.uc.SelectCityById(
		requestData.CityId,
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectCityById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectCityById",
		Data:    cities,
	}

	// create return
	return result, nil
}

func (ctrl Semeru1Controller) SelectCityByName(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectCityByName", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestCity
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectCityByName", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the city
	cities, err := ctrl.uc.SelectCityByName(
		requestData.CityName,
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectCityByName", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectCityByName",
		Data:    cities,
	}

	// create return
	return result, nil
}

func (ctrl Semeru1Controller) SelectCityByProvince(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectCityByProvince", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestCity
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectCityByProvince", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the city
	cities, err := ctrl.uc.SelectCityByProvince(
		requestData.CityProvince,
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectCityByProvince", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectCityByProvince",
		Data:    cities,
	}

	// create return
	return result, nil
}

func (ctrl Semeru1Controller) SelectAllCity(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectAllCity", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestCity
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectAllCity", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the city
	cities, err := ctrl.uc.SelectAllCity(
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectAllCity", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectAllCity",
		Data:    cities,
	}

	// create return
	return result, nil
}

func (ctrl Semeru1Controller) SelectProvinceById(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectProvinceById", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestProvince
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectProvinceById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the province
	provinces, err := ctrl.uc.SelectProvinceById(
		requestData.ProvinceId,
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectProvinceById", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectProvinceById",
		Data:    provinces,
	}

	// create return
	return result, nil
}

func (ctrl Semeru1Controller) SelectAllProvince(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to json data
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "SelectAllProvince", marshaledData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestProvince
	err = json.Unmarshal(marshaledData, &requestData)

	if err != nil {
		utils.Error(err, "SelectAllProvince", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the province
	provinces, err := ctrl.uc.SelectAllProvince(
		requestData.Limit,
		requestData.Offset)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "SelectAllProvince", requestData)
		// stop timer
		end := time.Now()

		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success SelectAllProvince",
		Data:    provinces,
	}

	// create return
	return result, nil
}
