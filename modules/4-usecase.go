package modules

import (
	"errors"
	"github.com/purawaktra/semeru1-go/entities"
	"github.com/purawaktra/semeru1-go/utils"
)

type Semeru1Usecase struct {
	repo Semeru1Repo
}

func CreateSemeru1Usecase(repo Semeru1Repo) Semeru1Usecase {
	return Semeru1Usecase{
		repo: repo,
	}
}

type Semeru1UsecaseInterface interface {
	SelectCityById(id int, limit int, offset int) ([]City, error)
	SelectCityByName(name string, limit int, offset int) ([]City, error)
	SelectCityByProvince(provinceId int, limit int, offset int) ([]City, error)
	SelectAllCity(limit int, offset int) ([]City, error)
	SelectProvinceById(id int, limit int, offset int) ([]Province, error)
	SelectAllProvince(limit int, offset int) ([]Province, error)
}

func (uc Semeru1Usecase) SelectCityById(id int, limit int, offset int) ([]City, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectCityById", offset)
		return nil, errors.New("offset can not be negative")
	}

	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectCityById", limit)
		return nil, errors.New("limit can not be zero or negative")
	}

	if id < 1 {
		utils.Error(errors.New("city id can not be zero or negative"), "SelectCityById", id)
		return nil, errors.New("city id can not be zero or negative")
	}

	// convert input to entity
	city := entities.Cities{CityId: uint(id)}

	// call repo for the city
	cities, err := uc.repo.SelectCityById(city, uint(limit), uint(offset))

	// check for error on call repo
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectCityById", city)
		return nil, err
	}

	// convert entity to dto
	results := make([]City, 0)
	for _, city := range cities {
		results = append(results, City{
			CityId:   city.CityId,
			Name:     city.Name,
			Province: city.Province,
		})
	}

	// create return
	return results, nil
}

func (uc Semeru1Usecase) SelectCityByName(name string, limit int, offset int) ([]City, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectCityByName", offset)
		return nil, errors.New("offset can not be negative")
	}

	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectCityByName", limit)
		return nil, errors.New("limit can not be zero or negative")
	}

	if name == "" {
		utils.Error(errors.New("city name can not be empty"), "SelectCityByName", name)
		return nil, errors.New("city name can not be empty")
	}

	// convert input to entity
	city := entities.Cities{Name: name}

	// call repo for the city
	cities, err := uc.repo.SelectCityByName(city, uint(limit), uint(offset))

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectCityByName", city)
		return nil, err
	}

	// convert entity to dto
	results := make([]City, 0)
	for _, city := range cities {
		results = append(results, City{
			CityId:   city.CityId,
			Name:     city.Name,
			Province: city.Province,
		})
	}

	// create return
	return results, nil
}

func (uc Semeru1Usecase) SelectCityByProvince(provinceId int, limit int, offset int) ([]City, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectCityByProvince", offset)
		return nil, errors.New("offset can not be negative")
	}

	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectCityByProvince", limit)
		return nil, errors.New("limit can not be zero or negative")
	}

	if provinceId < 1 {
		return nil, errors.New("province id can not be zero or negative")
	}

	// convert input to entity
	city := entities.Cities{Province: uint(provinceId)}

	// call repo for the city
	cities, err := uc.repo.SelectCityByProvince(city, uint(limit), uint(offset))

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectCityByProvince", city)
		return nil, err
	}

	// convert entity to dto
	results := make([]City, 0)
	for _, city := range cities {
		results = append(results, City{
			CityId:   city.CityId,
			Name:     city.Name,
			Province: city.Province,
		})
	}

	// create return
	return results, nil
}

func (uc Semeru1Usecase) SelectAllCity(limit int, offset int) ([]City, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectAllCity", offset)
		return nil, errors.New("offset can not be negative")
	}

	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectAllCity", limit)
		return nil, errors.New("limit can not be zero or negative")
	}

	// call repo for the city
	cities, err := uc.repo.SelectAllCity(uint(limit), uint(offset))

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectAllCity", "")
		return nil, err
	}

	// convert entity to dto
	results := make([]City, 0)
	for _, city := range cities {
		results = append(results, City{
			CityId:   city.CityId,
			Name:     city.Name,
			Province: city.Province,
		})
	}

	// create return
	return results, nil
}

func (uc Semeru1Usecase) SelectProvinceById(id int, limit int, offset int) ([]Province, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectProvinceById", offset)
		return nil, errors.New("offset can not be negative")
	}

	if id < 1 {
		utils.Error(errors.New("city id can not be zero or negative"), "SelectProvinceById", id)
		return nil, errors.New("province id can not be zero or negative")
	}

	// convert input to entity
	province := entities.Provinces{ProvinceId: uint(id)}

	// call repo for the province
	provinces, err := uc.repo.SelectProvinceById(province, uint(limit), uint(offset))

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectProvinceById", province)
		return nil, err
	}

	// convert entity to dto
	results := make([]Province, 0)
	for _, province := range provinces {
		results = append(results, Province{
			ProvinceId: province.ProvinceId,
			Name:       province.Name,
		})
	}

	// create return
	return results, nil
}

func (uc Semeru1Usecase) SelectAllProvince(limit int, offset int) ([]Province, error) {
	// create check input on id and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectAllProvince", offset)
		return nil, errors.New("offset can not be negative")
	}

	// call repo for the province
	provinces, err := uc.repo.SelectAllProvince(uint(limit), uint(offset))

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectAllProvince", "")
		return nil, err
	}

	// convert entity to dto
	results := make([]Province, 0)
	for _, province := range provinces {
		results = append(results, Province{
			ProvinceId: province.ProvinceId,
			Name:       province.Name,
		})
	}

	// create return
	return results, nil
}
