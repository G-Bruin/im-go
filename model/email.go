package model

import (
	"github.com/jinzhu/gorm"
)

type Email struct {
	//ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
	gorm.Model //默认生成 created_at updated_at deleted_at  deleted_at 默认索引
}
