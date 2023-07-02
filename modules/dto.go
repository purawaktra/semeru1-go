package modules

type City struct {
	CityId   uint   `json:"city_id"`
	Name     string `json:"name"`
	Province uint   `json:"province"`
}

type Province struct {
	ProvinceId uint   `json:"province_id"`
	Name       string `json:"name"`
}
