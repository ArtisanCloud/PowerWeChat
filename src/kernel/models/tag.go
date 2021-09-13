package models

type Tag struct {
	id       int    `gorm:"column:id json:id"`
	name     string `gorm:"column:name json:name"`
	nameEN   string `gorm:"column:name_en json:name_en"`
	parentID string `gorm:"column:parentid json:parentid"`
	order    int    `gorm:"column:order json:order"`
}
