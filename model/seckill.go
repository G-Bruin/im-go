package model

type Seckill struct {
	Id     int `gorm:"primary_key" json:"id"`
	UserId int `gorm:"index" json:"user_id"`
	GoodId int `gorm:"index" json:"good_id"`
	BaseModel
}
