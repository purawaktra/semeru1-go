package modules

import (
	"encoding/json"
	"github.com/purawaktra/semeru1-go/dto"
	"github.com/purawaktra/semeru1-go/utils"
)

type Semeru1Controller struct {
	uc Semeru1Usecase
}

func CreateSemeru1Controller(uc Semeru1Usecase) Semeru1Controller {
	return Semeru1Controller{
		uc: uc,
	}
}

func (ctrl *Semeru1Controller) SelectCityById(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCity
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectCityById", requestData)
		return nil, err, "FM"
	}

	// convert request to dto
	query := City{CityId: requestData.CityId}

	// call usecase for the city and check err
	city, err, code := ctrl.uc.SelectCityById(query)
	if err != nil {
		utils.Error(err, "SelectCityById", requestData)
		return nil, err, code
	}

	// create return
	return city, nil, code
}

func (ctrl *Semeru1Controller) SelectCityByName(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCity
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectCityByName", requestData)
		return nil, err, "FM"
	}

	// convert request to dto
	query := City{Name: requestData.CityName}

	// call usecase for the city and check err
	city, err, code := ctrl.uc.SelectCityByName(query)
	if err != nil {
		utils.Error(err, "SelectCityByName", requestData)
		return nil, err, code
	}

	// create return
	return city, nil, code
}

func (ctrl *Semeru1Controller) SelectAllCity(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyCity
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectAllCity", requestData)
		return nil, err, "FM"
	}

	// call usecase for the city and check err
	city, err, code := ctrl.uc.SelectAllCity(requestData.Limit, requestData.Offset)
	if err != nil {
		utils.Error(err, "SelectAllCity", requestData)
		return nil, err, code
	}

	// create return
	return city, nil, code
}

func (ctrl *Semeru1Controller) SelectProvinceById(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProvince
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectProvinceById", requestData)
		return nil, err, "FM"
	}

	// convert request to dto
	query := Province{ProvinceId: requestData.ProvinceId}

	// call usecase for the province and check err
	city, err, code := ctrl.uc.SelectProvinceById(query)
	if err != nil {
		utils.Error(err, "SelectProvinceById", requestData)
		return nil, err, code
	}

	// create return
	return city, nil, code
}

func (ctrl *Semeru1Controller) SelectAllProvince(req []byte) (any, error, string) {
	// unmarshal to struct and check err
	var requestData dto.BodyProvince
	err := json.Unmarshal(req, &requestData)
	if err != nil {
		utils.Error(err, "SelectAllProvince", requestData)
		return nil, err, "FM"
	}

	// call usecase for the province and check err
	city, err, code := ctrl.uc.SelectAllProvince(requestData.Limit, requestData.Offset)
	if err != nil {
		utils.Error(err, "SelectAllProvince", requestData)
		return nil, err, code
	}

	// create return
	return city, nil, code
}
