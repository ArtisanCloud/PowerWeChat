package models

type Department struct {
	ID       int    `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	NameEN   string `gorm:"column:name_en" json:"name_en"`
	ParentID int    `gorm:"column:parentid" json:"parentid"`
	Order    int    `gorm:"column:order" json:"order"`
}
