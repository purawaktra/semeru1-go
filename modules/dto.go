package modules

type City struct {
	CityId   int    `json:"city_id"`
	Name     string `json:"name"`
	Province int    `json:"province"`
}

type Province struct {
	ProvinceId int    `json:"province_id"`
	Name       string `json:"name"`
}
