package entities

type Cities struct {
	CityId   uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Province uint   `gorm:"column:province"`
}

type Provinces struct {
	ProvinceId uint   `gorm:"primary_key"`
	Name       string `gorm:"column:name"`
}
