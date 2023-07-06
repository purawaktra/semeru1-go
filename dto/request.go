package dto

type Request struct {
	RequestId string `json:"request_id"`
	Data      any    `json:"data"`
}

type RequestCity struct {
	CityId       int    `json:"city_id"`
	CityName     string `json:"city_name"`
	CityProvince int    `json:"city_province"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type RequestProvince struct {
	ProvinceId   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}
