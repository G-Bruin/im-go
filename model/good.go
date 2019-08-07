package model

type Good struct {
	Id    int `gorm:"primary_key" json:"id"`
	Total int `gorm:"index" json:"total"`
	BaseModel
}
