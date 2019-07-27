package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"datetime" json:"updated_at"`
}

type DeletedAt struct {
	DeletedAt time.Time `gorm:"datetime;index" json:"deleted_at"`
}

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(64);unique_index" json:"name" `
	Password string `gorm:"type:varchar(255)" json:"-" `
	BaseModel
	DeletedAt
	//CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Friends []Friend `json:"-" ` // One-To-Many (拥有多个 - Friend表的UserID作外键)
	//Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}
