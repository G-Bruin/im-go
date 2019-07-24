package model

type Language struct {
	ID   int    `form:"id" json:"id"`                                //格式化返回值
	Name string `gorm:"index:idx_name_code" form:"name" json:"name"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code" form:"code" json:"code"` // `unique_index` also works
}
