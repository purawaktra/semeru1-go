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

func (uc *Semeru1Usecase) SelectCityById(query City) (City, error, string) {
	// create check for city id
	if query.CityId < 1 {
		utils.Error(errors.New("city id can not be zero or negative"), "SelectCityById", query)
		return City{}, errors.New("city id can not be zero or negative"), "FC"
	}

	// convert input to entity
	entity := entities.City{CityId: uint(query.CityId)}

	// call repo for the city and check err
	city, err, code := uc.repo.SelectCityById(entity)
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectCityById", city)
		return City{}, err, code
	}

	// convert entity to dto
	result := City{
		CityId:   int(city.CityId),
		Name:     city.Name,
		Province: int(city.Province),
	}

	// create return
	return result, nil, code
}

func (uc *Semeru1Usecase) SelectCityByName(query City) (City, error, string) {
	// create check input on name
	if query.Name == "" {
		utils.Error(errors.New("city name can not be empty"), "SelectCityByName", query)
		return City{}, errors.New("city name can not be empty"), "FC"
	}

	// convert input to entity
	entity := entities.City{Name: query.Name}

	// call repo for the city and check err
	city, err, code := uc.repo.SelectCityByName(entity)
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectCityByName", city)
		return City{}, err, code
	}

	// convert entity to dto
	result := City{
		CityId:   int(city.CityId),
		Name:     city.Name,
		Province: int(city.Province),
	}

	// create return
	return result, nil, code
}

func (uc *Semeru1Usecase) SelectAllCity(limit int, offset int) ([]City, error, string) {
	// create check input on limit and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectAllCity", offset)
		return nil, errors.New("offset can not be negative"), "FC"
	}
	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectAllCity", limit)
		return nil, errors.New("limit can not be zero or negative"), "FC"
	}

	// call repo for the city and check err
	cities, err, code := uc.repo.SelectAllCity(uint(limit), uint(offset))
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectAllCity", "")
		return nil, err, code
	}

	// convert entity to dto
	results := make([]City, 0)
	for _, city := range cities {
		results = append(results, City{
			CityId:   int(city.CityId),
			Name:     city.Name,
			Province: int(city.Province),
		})
	}

	// create return
	return results, nil, code
}

func (uc *Semeru1Usecase) SelectProvinceById(query Province) (Province, error, string) {
	// create check input on id
	if query.ProvinceId < 1 {
		utils.Error(errors.New("city id can not be zero or negative"), "SelectProvinceById", query)
		return Province{}, errors.New("province id can not be zero or negative"), "FC"
	}

	// convert input to entity
	entity := entities.Province{ProvinceId: uint(query.ProvinceId)}

	// call repo for the province and check err
	province, err, code := uc.repo.SelectProvinceById(entity)
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectProvinceById", entity)
		return Province{}, err, code
	}

	// convert entity to dto
	result := Province{
		ProvinceId: int(province.ProvinceId),
		Name:       province.Name,
	}

	// create return
	return result, nil, code
}

func (uc *Semeru1Usecase) SelectAllProvince(limit int, offset int) ([]Province, error, string) {
	// create check input on limit and offset
	if offset < 0 {
		utils.Error(errors.New("offset can not be negative"), "SelectAllProvince", offset)
		return nil, errors.New("offset can not be negative"), "FC"
	}
	if limit < 1 {
		utils.Error(errors.New("limit can not be zero or negative"), "SelectAllProvince", limit)
		return nil, errors.New("limit can not be zero or negative"), "FC"
	}

	// call repo for the province and check err
	provinces, err, code := uc.repo.SelectAllProvince(uint(limit), uint(offset))
	if err != nil {
		utils.Error(errors.New("error call usecase"), "SelectAllProvince", "")
		return nil, err, code
	}

	// convert entity to dto
	results := make([]Province, 0)
	for _, province := range provinces {
		results = append(results, Province{
			ProvinceId: int(province.ProvinceId),
			Name:       province.Name,
		})
	}

	// create return
	return results, nil, code
}
