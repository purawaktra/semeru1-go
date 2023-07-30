package entities

type City struct {
	CityId   uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Province uint   `gorm:"column:province"`
}

type Province struct {
	ProvinceId uint   `gorm:"primary_key"`
	Name       string `gorm:"column:name"`
}
