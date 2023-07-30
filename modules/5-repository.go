package modules

import (
	"fmt"
	"github.com/purawaktra/semeru1-go/entities"
	"github.com/purawaktra/semeru1-go/utils"
	"gorm.io/gorm"
)

type Semeru1Repo struct {
	db *gorm.DB
}

func CreateSemeru1Repo(gorm *gorm.DB) Semeru1Repo {
	return Semeru1Repo{
		db: gorm,
	}
}

func (sr *Semeru1Repo) SelectCityById(query entities.City) (entities.City, error, string) {
	utils.Debug("SelectCityById", query)

	var city entities.City
	tx := sr.db.Raw(
		fmt.Sprint("SELECT city_id, name, province FROM cities WHERE city_id = ", query.CityId, " LIMIT 1 ")).Scan(
		&city)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCityById", query)
		return entities.City{}, err, "DB"
	}
	return city, nil, "00"
}

func (sr *Semeru1Repo) SelectCityByName(query entities.City) (entities.City, error, string) {
	utils.Debug("SelectCityByName", query)

	var city entities.City
	tx := sr.db.Raw(
		fmt.Sprint("SELECT city_id, name, province FROM cities WHERE name LIKE \"", query.Name, "\" LIMIT 1 ")).Scan(
		&city)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCityByName", query)
		return entities.City{}, err, "DB"
	}
	return city, nil, "00"
}

func (sr *Semeru1Repo) SelectAllCity(limit uint, offset uint) ([]entities.City, error, string) {
	utils.Debug("SelectAllCity", struct {
		limit  uint
		offset uint
	}{
		limit:  limit,
		offset: offset,
	})

	var cities = make([]entities.City, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT city_id, name, province FROM cities LIMIT ", limit, " OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectAllCity", struct {
			limit  uint
			offset uint
		}{
			limit:  limit,
			offset: offset,
		})
		return nil, err, "DB"
	}
	return cities, nil, "00"
}

func (sr *Semeru1Repo) SelectProvinceById(query entities.Province) (entities.Province, error, string) {
	utils.Debug("SelectProvinceById", query)

	var province entities.Province
	tx := sr.db.Raw(
		fmt.Sprint("SELECT province_id, name FROM provinces WHERE province_id = ", query.ProvinceId, " LIMIT 1")).Scan(
		&province)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectProvinceById", query)
		return entities.Province{}, err, "DB"
	}
	return province, nil, "00"
}

func (sr *Semeru1Repo) SelectAllProvince(limit uint, offset uint) ([]entities.Province, error, string) {
	utils.Debug("SelectAllProvince", struct {
		limit  uint
		offset uint
	}{
		limit:  limit,
		offset: offset,
	})

	var provinces = make([]entities.Province, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT province_id, name FROM provinces LIMIT ", limit, " OFFSET ", offset)).Scan(
		&provinces)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectAllProvince", struct {
			limit  uint
			offset uint
		}{
			limit:  limit,
			offset: offset,
		})
		return nil, err, "DB"
	}
	return provinces, nil, "00"
}
