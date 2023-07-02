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

type Semeru1RepoInterface interface {
	SelectCityById(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error)
	SelectCityByName(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error)
	SelectCityByProvince(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error)
	SelectAllCity(limit uint, offset uint) ([]entities.Cities, error)
	SelectProvinceById(query entities.Provinces, limit uint, offset uint) ([]entities.Provinces, error)
	SelectAllProvince(limit uint, offset uint) ([]entities.Provinces, error)
}

func (sr Semeru1Repo) SelectCityById(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error) {
	utils.Debug("SelectCityById", "=== New Query ===")
	utils.Debug("SelectCityById", query)

	var cities = make([]entities.Cities, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE city_id = ", query.CityId, " LIMIT ", limit, " OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCityById", query)
		return nil, err
	}
	return cities, nil
}

func (sr Semeru1Repo) SelectCityByName(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error) {
	utils.Debug("SelectCityByName", "=== New Query ===")
	utils.Debug("SelectCityByName", query)

	var cities = make([]entities.Cities, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE name LIKE \"", query.Name, "\" LIMIT ", limit, " OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCityByName", query)
		return nil, err
	}
	return cities, nil
}

func (sr Semeru1Repo) SelectCityByProvince(query entities.Cities, limit uint, offset uint) ([]entities.Cities, error) {
	utils.Debug("SelectCityByProvince", "=== New Query ===")
	utils.Debug("SelectCityByProvince", query)

	var cities = make([]entities.Cities, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE province = ", query.Province, " LIMIT ", limit, " OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectCityByProvince", query)
		return nil, err
	}
	return cities, nil
}

func (sr Semeru1Repo) SelectAllCity(limit uint, offset uint) ([]entities.Cities, error) {
	utils.Debug("SelectAllCity", "=== New Query ===")
	utils.Debug("SelectAllCity", struct {
		offset uint
	}{offset: offset})

	var cities = make([]entities.Cities, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM cities LIMIT ", limit, " OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectAllCity", struct {
			offset uint
		}{offset: offset})
		return nil, err
	}
	return cities, nil
}

func (sr Semeru1Repo) SelectProvinceById(query entities.Provinces, limit uint, offset uint) ([]entities.Provinces, error) {
	utils.Debug("SelectProvinceById", "=== New Query ===")
	utils.Debug("SelectProvinceById", struct {
		provinces entities.Provinces
		offset    uint
	}{
		provinces: query,
		offset:    offset,
	})

	var provinces = make([]entities.Provinces, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM provinces WHERE province_id = ", query.ProvinceId, " LIMIT ", limit, " OFFSET ", offset)).Scan(
		&provinces)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectProvinceById", struct {
			provinces entities.Provinces
			offset    uint
		}{
			provinces: query,
			offset:    offset,
		})
		return nil, err
	}
	return provinces, nil
}

func (sr Semeru1Repo) SelectAllProvince(limit uint, offset uint) ([]entities.Provinces, error) {
	utils.Debug("SelectAllProvince", "=== New Query ===")
	utils.Debug("SelectAllProvince", struct {
		offset uint
	}{offset: offset})

	var provinces = make([]entities.Provinces, 0)
	tx := sr.db.Raw(
		fmt.Sprint("SELECT * FROM provinces LIMIT ", limit, " OFFSET ", offset)).Scan(
		&provinces)
	err := tx.Error
	if err != nil {
		utils.Error(err, "SelectAllProvince", struct {
			offset uint
		}{offset: offset})
		return nil, err
	}
	return provinces, nil
}
