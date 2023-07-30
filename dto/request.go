package dto

type Request struct {
	RequestId string `json:"request_id"`
	Data      any    `json:"data"`
}

type BodyCity struct {
	CityId       int    `json:"city_id"`
	CityName     string `json:"city_name"`
	CityProvince int    `json:"city_province"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type BodyProvince struct {
	ProvinceId   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

type Header struct {
	RequestId string `header:"request-id"`
	Host      string `header:"host"`
}
